package stream

import (
	"fmt"
	"io"
	"log"

	"github.com/phbarton/Telemetry-Go/telemetry"
)

type closeableStreamTraceListener struct {
	inner  *telemetry.TraceListener
	writer *io.WriteCloser
}

// NewCloseableStreamTraceListener creates a trace listener which outputs to the provided implementation of io.WriteCloser interface. It limits output based on the logging level supplied
func NewCloseableStreamTraceListener(loggingLevel telemetry.Severity, writer *io.WriteCloser) telemetry.TraceListener {
	var w io.Writer = *writer

	innerListener := NewStreamTraceListener(loggingLevel, &w)
	traceListener := closeableStreamTraceListener{inner: &innerListener, writer: writer}

	return &traceListener
}

func (cstl *closeableStreamTraceListener) TraceMessage(message string, severity telemetry.Severity) {
	(*cstl.inner).TraceMessage(message, severity)
}

func (cstl *closeableStreamTraceListener) TraceException(err error) {
	(*cstl.inner).TraceMessage(err.Error(), telemetry.Error)
}

func (cstl *closeableStreamTraceListener) TracePanic(rethrow bool) {
	if r := recover(); r != nil {
		cstl.TraceMessage(fmt.Sprint(r), telemetry.Critical)

		if rethrow {
			panic(r)
		}
	}
}

func (cstl *closeableStreamTraceListener) TrackAvailability(name string) *telemetry.DurationTrace {
	return (*cstl.inner).TrackAvailability(name)
}

func (cstl *closeableStreamTraceListener) TrackRequest(method string, uri string) *telemetry.DurationTrace {
	return (*cstl.inner).TrackRequest(method, uri)
}

func (cstl *closeableStreamTraceListener) TrackDependency(name string, dependencyType string, target string) *telemetry.DurationTrace {
	return (*cstl.inner).TrackDependency(name, dependencyType, target)
}

func (cstl *closeableStreamTraceListener) TraceMetric(name string, value float64) {
	(*cstl.inner).TraceMetric(name, value)
}

func (cstl *closeableStreamTraceListener) TraceEvent(name string) {
	(*cstl.inner).TraceEvent(name)
}

func (cstl *closeableStreamTraceListener) Flush() {
	// Unused
}

func (cstl *closeableStreamTraceListener) Close() {
	if err := (*cstl.writer).Close(); err != nil {
		log.Printf("closeableStreamTraceListener.Close() failed: %v", err.Error())
	}
}
