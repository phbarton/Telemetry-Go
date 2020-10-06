package telemetry

import "github.com/phbarton/Telemetry-Go/telemetry/contracts"

// TraceListener is the interface to be implemented by all implementations for a common tracing capability
type TraceListener interface {
	TraceMessage(message string, severity contracts.Severity)

	TraceInformation(message string)

	TraceVerbose(message string)

	TraceWarning(message string)

	TraceError(message string)

	TraceCritical(message string)

	TraceException(err interface{})
}
