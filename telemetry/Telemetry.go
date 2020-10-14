package telemetry

var (
	traceListener []*TraceListener
)

// AddListener adds an implementation of the TraceListener interface to the list of all listeners
func AddListener(listener *TraceListener) {
	if traceListener == nil {
		traceListener = []*TraceListener{listener}
	} else {
		traceListener = append(traceListener, listener)
	}
}

// TraceVerbose writes a verbose message (typically for debugging) to the underlyng trace listeners
func TraceVerbose(message string) {
	traceMessageImpl(message, Verbose)
}

// TraceInformation writes an informational message to the underlyng trace listeners
func TraceInformation(message string) {
	traceMessageImpl(message, Information)
}

// TraceWarning writes a warning message to the underlyng trace listeners
func TraceWarning(message string) {
	traceMessageImpl(message, Warning)
}

// TraceError writes an error message to the underlyng trace listeners
func TraceError(message string) {
	traceMessageImpl(message, Error)
}

// TraceCritical writes a critical error message to the underlyng trace listeners
func TraceCritical(message string) {
	traceMessageImpl(message, Critical)
}

// TraceException traces the specified error to the underlyng trace listeners
func TraceException(err error) {
	if traceListener != nil {
		for _, tl := range traceListener {
			(*tl).TraceException(err)
		}
	}
}

// TracePanic traces any panic error that is thrown. Typically used in a defer statement.
func TracePanic(rethrow bool) {
	if traceListener != nil {
		for _, tl := range traceListener {
			(*tl).TracePanic(rethrow)
		}
	}
}

// TraceMetric traces named single-valued metric to the underlyng trace listeners
func TraceMetric(name string, value float64) {
	if traceListener != nil {
		for _, tl := range traceListener {
			(*tl).TraceMetric(name, value)
		}
	}
}

// TraceEvent traces named event to the underlyng trace listeners
func TraceEvent(name string) {
	if traceListener != nil {
		for _, tl := range traceListener {
			(*tl).TraceEvent(name)
		}
	}
}

// TrackAvailability creates a tracking of the availability of the named service
func TrackAvailability(name string) *DurationTrace {
	traces := make([]*DurationTrace, 0)

	if traceListener != nil {
		for _, tl := range traceListener {
			traces = append(traces, (*tl).TrackAvailability(name))
		}
	}

	dt := newAggregateDurationTrace(traces)
	return &dt
}

// TrackRequest creates a tracking of the service request at the specified URI and method
func TrackRequest(method string, uri string) *DurationTrace {
	traces := make([]*DurationTrace, 0)

	if traceListener != nil {
		for _, tl := range traceListener {
			traces = append(traces, (*tl).TrackRequest(method, uri))
		}
	}

	dt := newAggregateDurationTrace(traces)
	return &dt
}

// TrackDependency creates a tracking of the specified external service dependency
func TrackDependency(name string, dependencyType string, target string) *DurationTrace {
	traces := make([]*DurationTrace, 0)

	if traceListener != nil {
		for _, tl := range traceListener {
			traces = append(traces, (*tl).TrackDependency(name, dependencyType, target))
		}
	}

	dt := newAggregateDurationTrace(traces)
	return &dt
}

// Flush causes all trace listeners to flush their data to their respective providers.
func Flush() {
	if traceListener != nil {
		for _, tl := range traceListener {
			(*tl).Flush()
		}
	}
}

// Close closes all trace listeners and removes the references to them.
func Close() {
	Flush()

	if traceListener != nil {
		for _, tl := range traceListener {
			(*tl).Close()
		}
	}

	traceListener = make([]*TraceListener, 0, 0)
}

func traceMessageImpl(message string, severity Severity) {
	if traceListener != nil {
		for _, tl := range traceListener {
			(*tl).TraceMessage(message, severity)
		}
	}
}

type aggregateDurationTrace struct {
	traces []*DurationTrace
}

func newAggregateDurationTrace(tracers []*DurationTrace) DurationTrace {
	return &aggregateDurationTrace{traces: tracers}
}

func (atl aggregateDurationTrace) Complete() {
	for _, trace := range atl.traces {
		(*trace).Complete()
	}
}

func (atl aggregateDurationTrace) Fail(statusCode string) {
	for _, trace := range atl.traces {
		(*trace).Fail(statusCode)
	}
}

func (atl aggregateDurationTrace) Done() {
	for _, trace := range atl.traces {
		(*trace).Done()
	}
}
