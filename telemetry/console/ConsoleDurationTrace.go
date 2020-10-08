package console

import (
	"fmt"
	"time"

	"github.com/phbarton/Telemetry-Go/telemetry"
)

type consoleDurationTrace struct {
	traceListener *consoleTraceListener
	statusCode    string
	success       bool
	startTime     time.Time
	output        string
}

// Complete indicates a successful completion of the measured duration activity
func (cdt *consoleDurationTrace) Complete() {
	cdt.success = true
	cdt.statusCode = "OK"
}

// Fail indicates an unsuccessful completion of the measured duration activity
func (cdt *consoleDurationTrace) Fail(statusCode string) {
	cdt.success = false
	cdt.statusCode = statusCode
}

// Done indicates that the trace is complete and should be committed to the telemetry source
func (cdt *consoleDurationTrace) Done() {
	duration := time.Now().Sub(cdt.startTime)

	if cdt.success {
		cdt.traceListener.TraceMessage(fmt.Sprintf("%v, Duration: %vms, Success", cdt.output, duration.Milliseconds()), telemetry.Information)
	} else {
		cdt.traceListener.TraceMessage(fmt.Sprintf("%v, Duration: %vms, Failed: %v", cdt.output, duration.Milliseconds(), cdt.statusCode), telemetry.Error)
	}
}
