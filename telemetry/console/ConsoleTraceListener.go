package console

import (
	"fmt"
	"time"

	"github.com/gookit/color"
	"github.com/phbarton/Telemetry-Go/telemetry"
)

type consoleTraceListener struct {
	loggingLevel telemetry.Severity
}

// NewConsoleTraceListener creates a trace listener which outputs to the console. It limits output based on the logging level supplied
func NewConsoleTraceListener(loggingLevel telemetry.Severity) telemetry.TraceListener {
	traceListener := consoleTraceListener{loggingLevel: loggingLevel}

	return &traceListener
}

func (ctl *consoleTraceListener) TraceMessage(message string, severity telemetry.Severity) {
	if severity >= ctl.loggingLevel {
		fmt.Printf("%v %v: %v\n", time.Now().Format(time.StampMilli), getSeverityTag(severity), message)
	}
}

func (ctl *consoleTraceListener) TraceException(err error) {
	ctl.TraceMessage(err.Error(), telemetry.Error)
}

func (ctl *consoleTraceListener) TrackAvailability(name string) *telemetry.DurationTrace {
	durationTrace := ctl.newDurationTrace(fmt.Sprintf("AVAILABILITY: %v", name))

	return &durationTrace
}

func (ctl *consoleTraceListener) TrackRequest(method string, uri string) *telemetry.DurationTrace {
	durationTrace := ctl.newDurationTrace(fmt.Sprintf("REQUEST: %v %v", method, uri))

	return &durationTrace
}

func (ctl *consoleTraceListener) TrackDependency(name string, dependencyType string, target string) *telemetry.DurationTrace {
	durationTrace := ctl.newDurationTrace(fmt.Sprintf("DEPENDENCY: %v (%v) %v", name, dependencyType, target))

	return &durationTrace
}

func (ctl *consoleTraceListener) TraceMetric(name string, value float64) {
	ctl.TraceMessage(fmt.Sprintf("METRIC: '%v': %v", name, value), telemetry.Information)
}

func (ctl *consoleTraceListener) TraceEvent(name string) {
	ctl.TraceMessage(fmt.Sprintf("EVENT: %v", name), telemetry.Verbose)
}

func (ctl *consoleTraceListener) Flush() {
	// Unused
}

func (ctl *consoleTraceListener) Close() {
	// Unused
}

func (ctl *consoleTraceListener) newDurationTrace(output string) telemetry.DurationTrace {
	return &consoleDurationTrace{
		traceListener: ctl,
		output:        output,
		startTime:     time.Now(),
	}
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
