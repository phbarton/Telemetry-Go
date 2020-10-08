package telemetry

// DurationTrace provides an interface for those traces which require a duration
type DurationTrace interface {
	// Complete indicates a successful completion of the measured duration activity
	Complete()

	// Fail indicates an unsuccessful completion of the measured duration activity
	Fail(statusCode string)

	// Done indicates that the trace is complete and should be committed to the telemetry source
	Done()
}
