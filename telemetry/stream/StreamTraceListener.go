package stream

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/phbarton/Telemetry-Go/telemetry"
)

type streamTraceListener struct {
	loggingLevel telemetry.Severity
	writer       *io.Writer
}

// NewStreamTraceListener creates a trace listener which outputs to the console. It limits output based on the logging level supplied
func NewStreamTraceListener(loggingLevel telemetry.Severity, writer *io.Writer) telemetry.TraceListener {
	traceListener := streamTraceListener{loggingLevel: loggingLevel, writer: writer}

	return &traceListener
}

func (stl *streamTraceListener) TraceMessage(message string, severity telemetry.Severity) {
	if severity >= stl.loggingLevel {
		entry := fmt.Sprintf("%v [%v]: %v\n", time.Now().Format(time.StampMilli), getSeverityTag(severity), message)

		if _, err := (*stl.writer).Write([]byte(entry)); err != nil {
			log.Printf("streamTraceListener.TraceMessage failed: %v", err.Error())
		}
	}
}

func (stl *streamTraceListener) TraceException(err error) {
	stl.TraceMessage(err.Error(), telemetry.Error)
}

func (stl *streamTraceListener) TrackAvailability(name string) *telemetry.DurationTrace {
	durationTrace := stl.newDurationTrace(fmt.Sprintf("AVAILABILITY: %v", name))

	return &durationTrace
}

func (stl *streamTraceListener) TrackRequest(method string, uri string) *telemetry.DurationTrace {
	durationTrace := stl.newDurationTrace(fmt.Sprintf("REQUEST: %v %v", method, uri))

	return &durationTrace
}

func (stl *streamTraceListener) TrackDependency(name string, dependencyType string, target string) *telemetry.DurationTrace {
	durationTrace := stl.newDurationTrace(fmt.Sprintf("DEPENDENCY: %v (%v) %v", name, dependencyType, target))

	return &durationTrace
}

func (stl *streamTraceListener) TraceMetric(name string, value float64) {
	stl.TraceMessage(fmt.Sprintf("METRIC: '%v': %v", name, value), telemetry.Information)
}

func (stl *streamTraceListener) TraceEvent(name string) {
	stl.TraceMessage(fmt.Sprintf("EVENT: %v", name), telemetry.Verbose)
}

func (stl *streamTraceListener) Flush() {
	// Unused
}

func (stl *streamTraceListener) Close() {
	// Unsused
}

func (stl *streamTraceListener) newDurationTrace(output string) telemetry.DurationTrace {
	return &streamDurationTrace{
		traceListener: stl,
		output:        output,
		startTime:     time.Now(),
	}
}

func getSeverityTag(severity telemetry.Severity) string {
	switch severity {
	case telemetry.Verbose:
		return "VRB"
	case telemetry.Information:
		return "INF"
	case telemetry.Warning:
		return "WRN"
	case telemetry.Error:
		return "ERR"
	case telemetry.Critical:
		return "CRT"
	default:
		return "UNK"
	}
}
