package telemetry

var (
	traceListener []*TraceListener
)

type aggregateDurationTrace struct {
	traces []*DurationTrace
}

func init() {
	traceListener = make([]*TraceListener, 0)
}

func AddListener(listener *TraceListener) {
	traceListener = append(traceListener, listener)
}

func TraceMessage(message string, severity Severity) {
	for _, tl := range traceListener {
		(*tl).TraceMessage(message, severity)
	}
}

func TraceVerbose(message string) {
	TraceMessage(message, Verbose)
}

func TraceInformation(message string) {
	TraceMessage(message, Information)
}

func TraceWarning(message string) {
	TraceMessage(message, Warning)
}

func TraceError(message string) {
	TraceMessage(message, Error)
}

func TraceCritical(message string) {
	TraceMessage(message, Critical)
}

func TraceException(err error) {
	for _, tl := range traceListener {
		(*tl).TraceException(err)
	}
}

func TraceMetric(name string, value float64) {
	for _, tl := range traceListener {
		(*tl).TraceMetric(name, value)
	}
}

func TraceEvent(name string) {
	for _, tl := range traceListener {
		(*tl).TraceEvent(name)
	}
}

func TrackAvailability(name string) *DurationTrace {
	traces := make([]*DurationTrace, 0)

	for _, tl := range traceListener {
		traces = append(traces, (*tl).TrackAvailability(name))
	}

	dt := newAggregateDurationTrace(traces)
	return &dt
}

func TrackRequest(method string, uri string) *DurationTrace {
	traces := make([]*DurationTrace, 0)

	for _, tl := range traceListener {
		traces = append(traces, (*tl).TrackRequest(method, uri))
	}

	dt := newAggregateDurationTrace(traces)
	return &dt
}

func TrackDependency(name string, dependencyType string, target string) *DurationTrace {
	traces := make([]*DurationTrace, 0)

	for _, tl := range traceListener {
		traces = append(traces, (*tl).TrackDependency(name, dependencyType, target))
	}

	dt := newAggregateDurationTrace(traces)
	return &dt
}

func Flush() {
	for _, tl := range traceListener {
		(*tl).Flush()
	}
}

func newAggregateDurationTrace(tracers []*DurationTrace) DurationTrace {
	return &(aggregateDurationTrace{traces: tracers})
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
