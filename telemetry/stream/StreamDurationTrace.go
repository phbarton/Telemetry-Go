package stream

import (
	"fmt"
	"time"

	"github.com/phbarton/Telemetry-Go/telemetry"
)

type streamDurationTrace struct {
	traceListener *streamTraceListener
	statusCode    string
	success       bool
	startTime     time.Time
	output        string
}

// Complete indicates a successful completion of the measured duration activity
func (sdt *streamDurationTrace) Complete() {
	sdt.success = true
	sdt.statusCode = "OK"
}

// Fail indicates an unsuccessful completion of the measured duration activity
func (sdt *streamDurationTrace) Fail(statusCode string) {
	sdt.success = false
	sdt.statusCode = statusCode
}

// Done indicates that the trace is complete and should be committed to the telemetry source
func (sdt *streamDurationTrace) Done() {
	duration := time.Now().Sub(sdt.startTime)

	if sdt.success {
		sdt.traceListener.TraceMessage(fmt.Sprintf("%v, Duration: %vms, Success", sdt.output, duration.Milliseconds()), telemetry.Information)
	} else {
		sdt.traceListener.TraceMessage(fmt.Sprintf("%v, Duration: %vms, Failed: %v", sdt.output, duration.Milliseconds(), sdt.statusCode), telemetry.Error)
	}
}
