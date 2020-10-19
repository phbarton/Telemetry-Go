package stream

import (
	"io"
	"log"
	"sync"
)

// streamTraceListenerChannel is a complex channel that allows for async send of messages, and control of the channel.
type streamTraceListenerChannel struct {
	emitChannel    chan string
	controlChannel chan *channelStateControl
	waitGroup      sync.WaitGroup
	writer         *io.Writer
}

// newStreamTraceListenerChannel creates a new instance of the streamTraceListenerChannel and begins listening
func newStreamTraceListenerChannel(writer *io.Writer) *streamTraceListenerChannel {
	channel := &streamTraceListenerChannel{
		emitChannel:    make(chan string, 10), // Buffered so that we can get some performance boost
		controlChannel: make(chan *channelStateControl),
		writer:         writer,
	}

	// Start the listening loop
	go channel.receiveLoop()

	return channel
}

// Close flushes the content and closes the channel
func (ch *streamTraceListenerChannel) Close() chan struct{} {
	if ch.controlChannel != nil {
		// Make a callback channel to indicate that everything is done
		c := make(chan struct{})
		ctl := &channelStateControl{
			completed: c,
			stop:      true,
		}

		// Push the message to the control channel
		ch.controlChannel <- ctl
		return c
	}

	return nil
}

// Stop force stops the channel. It will not wait for any messages to finish completing before closing.
func (ch *streamTraceListenerChannel) Stop() {
	if ch.controlChannel != nil {
		ch.controlChannel <- &channelStateControl{stop: true}
	}
}

// Send puts the message in the channel to be picked up and written to the target
func (ch *streamTraceListenerChannel) Send(message string) {
	if ch.emitChannel != nil {
		ch.waitGroup.Add(1)
		ch.emitChannel <- message
	}
}

// receiveLoop is the listener for message in the emit and control channels.
func (ch *streamTraceListenerChannel) receiveLoop() {
	control := newChannelState(ch)

	// Process until we are told to stop
	for !control.stopping {
		control.process()
	}

	control.stop()
}

// notifyComplete is used to signal when all messages are sent to the target (flushed) if a callback is supplied.
func (ch *streamTraceListenerChannel) notifyComplete(callback chan struct{}) {
	if callback != nil {
		go func() {
			ch.waitGroup.Wait()
			close(callback)
		}()
	}
}

// channelStateControl is a structure used to control the state of the complex channel
type channelStateControl struct {
	completed chan struct{}
	stop      bool
}

// channelState represents the state of the channel and performs the processing of the messages.
type channelState struct {
	stopping bool
	channel  *streamTraceListenerChannel
}

// newChannelState creates a new instance of the channelState class
func newChannelState(channel *streamTraceListenerChannel) *channelState {
	return &channelState{channel: channel, stopping: false}
}

// process handles the processing of messages from the two channels in the complex channel
func (ctl *channelState) process() {
	select {
	case msg := <-ctl.channel.emitChannel:
		// If we get something from the emit channel, we are writing to the target
		ctl.send(msg)

	case control := <-ctl.channel.controlChannel:
		// If we get something from the control channel, we are changing the state of the complex channel
		if control.stop {
			ctl.channel.notifyComplete(control.completed)
			ctl.stopping = true
		}
	}
}

// send writes the supplied message to the target
func (ctl *channelState) send(message string) {
	// Tell the channel that we're done when we exit the function
	defer ctl.channel.waitGroup.Done()

	// Write the message to the target
	if _, err := (*ctl.channel.writer).Write([]byte(message)); err != nil {
		log.Fatalf("Unexpected error when writing message to target: %v", err.Error())
	}
}

// stop closes all channels and returns when all messages have been written
func (ctl *channelState) stop() {
	close(ctl.channel.controlChannel)
	close(ctl.channel.emitChannel)

	ctl.channel.controlChannel = nil
	ctl.channel.emitChannel = nil

	ctl.channel.waitGroup.Wait()
}
