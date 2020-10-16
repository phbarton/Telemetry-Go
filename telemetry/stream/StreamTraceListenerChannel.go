package stream

import (
	"io"
	"log"
	"sync"
)

type streamTraceListenerChannel struct {
	emitChannel    chan string
	controlChannel chan *channelControl
	waitGroup      sync.WaitGroup
	writer         *io.Writer
}

func newStreamTraceListenerChannel(writer *io.Writer) *streamTraceListenerChannel {
	channel := &streamTraceListenerChannel{
		emitChannel:    make(chan string),
		controlChannel: make(chan *channelControl),
		writer:         writer,
	}

	go channel.receiveLoop()

	return channel
}

func (ch *streamTraceListenerChannel) Close() chan struct{} {
	if ch.controlChannel != nil {
		c := make(chan struct{})
		ctl := &channelControl{
			completed: c,
			stop:      true,
		}

		ch.controlChannel <- ctl
		return c
	}

	return nil
}

func (ch *streamTraceListenerChannel) Stop() {
	if ch.controlChannel != nil {
		ch.controlChannel <- &channelControl{stop: true}
	}
}

func (ch *streamTraceListenerChannel) Send(message string) {
	if ch.emitChannel != nil {
		ch.emitChannel <- message
	}
}

func (ch *streamTraceListenerChannel) receiveLoop() {
	control := newChannelState(ch)

	for !control.stopping {
		control.process()
	}

	control.stop()
}

func (ch *streamTraceListenerChannel) notifyComplete(callback chan struct{}) {
	if callback != nil {
		go func() {
			ch.waitGroup.Wait()
			close(callback)
		}()
	}
}

type channelControl struct {
	completed chan struct{}
	stop      bool
}

type channelState struct {
	stopping bool
	channel  *streamTraceListenerChannel
}

func newChannelState(channel *streamTraceListenerChannel) *channelState {
	return &channelState{channel: channel, stopping: false}
}

func (ctl *channelState) process() {
	select {
	case msg := <-ctl.channel.emitChannel:
		ctl.send(msg)

	case control := <-ctl.channel.controlChannel:
		if control.stop {
			ctl.channel.notifyComplete(control.completed)
			ctl.stopping = true
		}
	}
}

func (ctl *channelState) send(message string) {
	ctl.channel.waitGroup.Add(1)

	go func(message string, writer *io.Writer) {
		defer ctl.channel.waitGroup.Done()

		if _, err := (*writer).Write([]byte(message)); err != nil {
			log.Fatalf("Unexpected error when writing message to target: %v", err.Error())
		}
	}(message, ctl.channel.writer)
}

func (ctl *channelState) stop() {
	close(ctl.channel.controlChannel)
	close(ctl.channel.emitChannel)

	ctl.channel.controlChannel = nil
	ctl.channel.emitChannel = nil

	ctl.channel.waitGroup.Wait()
}
