package stream

import "sync"

type streamTraceListenerChannel struct {
	emitChannel    chan string
	controlChannel chan *channelControl
	waitGroup      sync.WaitGroup
}

func (ch *streamTraceListenerChannel) Send(message string) {

}

type channelControl struct {
	stop      bool
	completed chan struct{}
}
