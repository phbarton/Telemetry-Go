package telemetry

import (
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

// TestAddTraceListener checks whether a trace listener can be added
func TestAddTraceListener(t *testing.T) {
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener inteface")
	{
		t.Log("\tWhen the trace listener is added to the Telemetry")
		{
			tl := newEmptyTraceListener()
			AddListener(&tl)

			actualValue := len(traceListener)

			if actualValue == expectedValue {
				t.Logf("\t\t[%v] There should only be one trace listener in the global list of listeners.", checkMark)
			} else {
				t.Errorf("\t\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
			}
		}
	}
}

// TestClose checks whether all trace listeners are removed when Close() is called
func TestClose(t *testing.T) {
	expectedValue := 1
	expectedAfterCloseValue := 0

	t.Log("Given an implementation of the TraceListener inteface")
	{
		t.Log("\tWhen the trace listener is added to the Telemetry")
		{
			tl := newEmptyTraceListener()
			AddListener(&tl)

			actualValue := len(traceListener)

			if actualValue == expectedValue {
				t.Logf("\t\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
			} else {
				t.Fatalf("\t\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
			}

			Close()

			actualAfterCloseValue := len(traceListener)

			if actualAfterCloseValue == expectedAfterCloseValue {
				t.Logf("\t\t[%v] There should be no trace listeners in the global list of listeners after closing.", checkMark)
			} else {
				t.Errorf("\t\t[%v] There should be no trace listeners in the global list of listeners after closing. Expected: %v, Actual: %v", ballotX, expectedAfterCloseValue, actualAfterCloseValue)
			}
		}
	}
}

func newEmptyTraceListener() TraceListener {
	return &emptyTraceListener{}
}

type emptyTraceListener struct{}

func (etl *emptyTraceListener) TraceMessage(message string, severity Severity) {}

func (etl *emptyTraceListener) TraceException(err error) {}

func (etl *emptyTraceListener) TracePanic(rethrow bool) {}

func (etl *emptyTraceListener) TrackAvailability(name string) *DurationTrace { return nil }

func (etl *emptyTraceListener) TrackRequest(method string, uri string) *DurationTrace {
	return nil
}

func (etl *emptyTraceListener) TrackDependency(name string, dependencyType string, target string) *DurationTrace {
	return nil
}

func (etl *emptyTraceListener) TraceMetric(name string, value float64) {}

func (etl *emptyTraceListener) TraceEvent(name string) {}

func (etl *emptyTraceListener) Flush() {}

func (etl *emptyTraceListener) Close() {}
