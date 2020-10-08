package telemetry

// TraceListener is the interface to be implemented by all implementations for a common tracing capability
type TraceListener interface {
	TraceMessage(message string, severity Severity)

	TraceInformation(message string)

	TraceVerbose(message string)

	TraceWarning(message string)

	TraceError(message string)

	TraceCritical(message string)

	TraceException(err error)

	TrackAvailability(name string) *DurationTrace

	TrackRequest(method string, uri string) *DurationTrace

	TrackDependency(name string, dependencyType string, target string) *DurationTrace

	TraceMetric(name string, value float64)

	TraceEvent(name string)

	Flush()
}
