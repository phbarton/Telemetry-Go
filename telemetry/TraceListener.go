package telemetry

// TraceListener is the interface to be implemented by all implementations for a common tracing capability
type TraceListener interface {
	TraceMessage(message string, severity Severity)

	TraceInformation(message string)

	TraceVerbose(message string)

	TraceWarning(message string)

	TraceError(message string)

	TraceCritical(message string)

	TraceException(err interface{})

	Flush()
}
