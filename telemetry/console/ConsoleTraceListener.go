package console

import (
	"io"
	"os"

	"github.com/gookit/color"
	"github.com/phbarton/Telemetry-Go/telemetry"
	"github.com/phbarton/Telemetry-Go/telemetry/stream"
)

type consoleTraceListener struct {
	inner *telemetry.TraceListener
}

// NewConsoleTraceListener creates a trace listener which outputs to the console. It limits output based on the logging level supplied
func NewConsoleTraceListener(loggingLevel telemetry.Severity) telemetry.TraceListener {
	var console io.Writer = os.Stdout

	inner := stream.NewStreamTraceListener(loggingLevel, &console)
	traceListener := consoleTraceListener{inner: &inner}

	return &traceListener
}

func (ctl *consoleTraceListener) TraceMessage(message string, severity telemetry.Severity) {
	(*ctl.inner).TraceMessage(message, severity)
}

func (ctl *consoleTraceListener) TraceException(err error) {
	ctl.TraceMessage(err.Error(), telemetry.Error)
}

func (ctl *consoleTraceListener) TracePanic(rethrow bool) {
	(*ctl.inner).TracePanic(rethrow)
}

func (ctl *consoleTraceListener) TrackAvailability(name string) *telemetry.DurationTrace {
	return (*ctl.inner).TrackAvailability(name)
}

func (ctl *consoleTraceListener) TrackRequest(method string, uri string) *telemetry.DurationTrace {
	return (*ctl.inner).TrackRequest(method, uri)
}

func (ctl *consoleTraceListener) TrackDependency(name string, dependencyType string, target string) *telemetry.DurationTrace {
	return (*ctl.inner).TrackDependency(name, dependencyType, target)
}

func (ctl *consoleTraceListener) TraceMetric(name string, value float64) {
	(*ctl.inner).TraceMetric(name, value)
}

func (ctl *consoleTraceListener) TraceEvent(name string) {
	(*ctl.inner).TraceEvent(name)
}

func (ctl *consoleTraceListener) Flush() {
	// Unused
}

func (ctl *consoleTraceListener) Close() {
	// Unused
}

func getSeverityTag(severity telemetry.Severity) string {
	switch severity {
	case telemetry.Verbose:
		return color.New(color.FgBlack, color.BgGray).Render("VRB")
	case telemetry.Information:
		return color.New(color.FgWhite, color.BgBlue).Render("INF")
	case telemetry.Warning:
		return color.New(color.FgWhite, color.BgYellow).Render("WRN")
	case telemetry.Error:
		return color.New(color.FgWhite, color.BgRed).Render("ERR")
	case telemetry.Critical:
		return color.New(color.FgBlack, color.BgRed).Render("CRT")
	default:
		return color.New(color.FgGray).Render("UNK")
	}
}
