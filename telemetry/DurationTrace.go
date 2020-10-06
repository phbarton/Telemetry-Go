package telemetry

type DurationTrace interface {
	Complete()

	Fail(statusCode string)

	Dispose()
}
