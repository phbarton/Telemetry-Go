package telemetry

// TraceListener is the interface to be implemented by all implementations for a common tracing capability
type TraceListener interface {
	TraceMessage(message string, severity Severity)

	TraceException(err error)

	TrackAvailability(name string) *DurationTrace

	TrackRequest(method string, uri string) *DurationTrace

	TrackDependency(name string, dependencyType string, target string) *DurationTrace

	TraceMetric(name string, value float64)

	TraceEvent(name string)

	Flush()

	Close()
}
