package telemetry

import (
	"testing"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

func TestCloseWithNoListeners(t *testing.T) {
	t.Log("Give the Telemetry static methods")
	{
		t.Log("\tWhen no listeners are present")
		{
			Close()

			if r := recover(); r == nil {
				t.Logf("\t\t[%v] There should be no error/panic created when Close() is called.", checkMark)
			} else {
				t.Errorf("\t\t[%v] There should be no error/panic created when Close() is called. Error: %v", ballotX, r)
			}
		}
	}
}

func TestFlushWithNoListeners(t *testing.T) {
	t.Log("Give the Telemetry static methods")
	{
		t.Log("\tWhen no listeners are present")
		{
			Flush()

			if r := recover(); r == nil {
				t.Logf("\t\t[%v] There should be no error/panic created when Flush() is called.", checkMark)
			} else {
				t.Errorf("\t\t[%v] There should be no error/panic created when Flush() is called. Error: %v", ballotX, r)
			}
		}
	}
}

// TestAddTraceListener checks whether a trace listener can be added
func TestAddTraceListener(t *testing.T) {
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener inteface")
	{
		tl := newEmptyTraceListener()

		t.Log("\tWhen the trace listener is added to the Telemetry")
		{
			AddListener(&tl)
			actualValue := len(traceListeners)

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
		tl := newEmptyTraceListener()

		t.Log("\tWhen the trace listener is added to the Telemetry")
		{
			AddListener(&tl)
			actualValue := len(traceListeners)

			if actualValue == expectedValue {
				t.Logf("\t\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
			} else {
				t.Fatalf("\t\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
			}

			Close()

			actualAfterCloseValue := len(traceListeners)

			if actualAfterCloseValue == expectedAfterCloseValue {
				t.Logf("\t\t[%v] There should be no trace listeners in the global list of listeners after closing.", checkMark)
			} else {
				t.Errorf("\t\t[%v] There should be no trace listeners in the global list of listeners after closing. Expected: %v, Actual: %v", ballotX, expectedAfterCloseValue, actualAfterCloseValue)
			}
		}
	}
}

func TestEnsureTraceVerboseDataIsPassedToListener(t *testing.T) {
	expectedMessage := "Message"
	expectedSeverity := Verbose
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		info := &trackingInformation{}
		rtl := newRecordingTraceListener(info)

		AddListener(&rtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a verbose trace message is sent")
		{
			TraceVerbose(expectedMessage)
			actualMessage := info.message
			actualSeverity := info.severity

			if actualMessage == expectedMessage && actualSeverity == expectedSeverity {
				t.Logf("\t\t[%v] The message and severity are correctly passed to the underlying trace listener.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The message and severity are correctly passed to the underlying trace listener. Expected message: \"%v\", Actual message: \"%v\"; Expected severity: %v, Actual severity %v", ballotX, expectedMessage, actualMessage, expectedSeverity.ToString(), actualSeverity.ToString())
			}
		}
	}
}

func TestEnsureTraceInformationDataIsPassedToListener(t *testing.T) {
	expectedMessage := "Message"
	expectedSeverity := Information
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		info := &trackingInformation{}
		rtl := newRecordingTraceListener(info)

		AddListener(&rtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a information trace message is sent")
		{
			TraceInformation(expectedMessage)
			actualMessage := info.message
			actualSeverity := info.severity

			if actualMessage == expectedMessage && actualSeverity == expectedSeverity {
				t.Logf("\t\t[%v] The message and severity are correctly passed to the underlying trace listener.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The message and severity are correctly passed to the underlying trace listener. Expected message: \"%v\", Actual message: \"%v\"; Expected severity: %v, Actual severity %v", ballotX, expectedMessage, actualMessage, expectedSeverity.ToString(), actualSeverity.ToString())
			}
		}
	}
}

func TestEnsureTraceWarningDataIsPassedToListener(t *testing.T) {
	expectedMessage := "Message"
	expectedSeverity := Warning
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		info := &trackingInformation{}
		rtl := newRecordingTraceListener(info)

		AddListener(&rtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a warning trace message is sent")
		{
			TraceWarning(expectedMessage)
			actualMessage := info.message
			actualSeverity := info.severity

			if actualMessage == expectedMessage && actualSeverity == expectedSeverity {
				t.Logf("\t\t[%v] The message and severity are correctly passed to the underlying trace listener.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The message and severity are correctly passed to the underlying trace listener. Expected message: \"%v\", Actual message: \"%v\"; Expected severity: %v, Actual severity %v", ballotX, expectedMessage, actualMessage, expectedSeverity.ToString(), actualSeverity.ToString())
			}
		}
	}
}

func TestEnsureTraceErrorDataIsPassedToListener(t *testing.T) {
	expectedMessage := "Message"
	expectedSeverity := Error
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		info := &trackingInformation{}
		rtl := newRecordingTraceListener(info)

		AddListener(&rtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a error trace message is sent")
		{
			TraceError(expectedMessage)
			actualMessage := info.message
			actualSeverity := info.severity

			if actualMessage == expectedMessage && actualSeverity == expectedSeverity {
				t.Logf("\t\t[%v] The message and severity are correctly passed to the underlying trace listener.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The message and severity are correctly passed to the underlying trace listener. Expected message: \"%v\", Actual message: \"%v\"; Expected severity: %v, Actual severity %v", ballotX, expectedMessage, actualMessage, expectedSeverity.ToString(), actualSeverity.ToString())
			}
		}
	}
}

func TestEnsureTraceCriticalDataIsPassedToListener(t *testing.T) {
	expectedMessage := "Message"
	expectedSeverity := Critical
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		info := &trackingInformation{}
		rtl := newRecordingTraceListener(info)

		AddListener(&rtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a critical trace message is sent")
		{
			TraceCritical(expectedMessage)
			actualMessage := info.message
			actualSeverity := info.severity

			if actualMessage == expectedMessage && actualSeverity == expectedSeverity {
				t.Logf("\t\t[%v] The message and severity are correctly passed to the underlying trace listener.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The message and severity are correctly passed to the underlying trace listener. Expected message: \"%v\", Actual message: \"%v\"; Expected severity: %v, Actual severity %v", ballotX, expectedMessage, actualMessage, expectedSeverity.ToString(), actualSeverity.ToString())
			}
		}
	}
}

func TestEnsureTraceExceptionDataIsPassedToListener(t *testing.T) {
	expectedMessage := "Message"
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		info := &trackingInformation{}
		rtl := newRecordingTraceListener(info)
		err := &testError{err: expectedMessage}

		AddListener(&rtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a exception trace message is sent")
		{
			TraceException(err)
			actualMessage := info.err.Error()

			if actualMessage == expectedMessage {
				t.Logf("\t\t[%v] The error is correctly passed to the underlying trace listener.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The error is correctly passed to the underlying trace listener. Expected message: \"%v\", Actual message: \"%v\"", ballotX, expectedMessage, actualMessage)
			}
		}
	}
}

func TestEnsureTracePanicDataIsPassedToListener(t *testing.T) {
	expectedValue := true
	expectedCount := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		info := &trackingInformation{}
		rtl := newRecordingTraceListener(info)

		AddListener(&rtl)
		actualCount := len(traceListeners)

		if actualCount == expectedCount {
			t.Logf("\t\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedCount, actualCount)
		}

		t.Log("\tWhen a exception trace message is sent")
		{
			TracePanic(expectedValue)
			actualValue := info.rethrow

			if actualValue == expectedValue {
				t.Logf("\t\t[%v] The rethrow value is correctly passed to the underlying trace listener.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The rethrow value is correctly passed to the underlying trace listener. Expected value: \"%v\", Actual value: \"%v\"", ballotX, expectedValue, actualValue)
			}
		}
	}
}

func TestEnsureTraceMetricDataIsPassedToListener(t *testing.T) {
	expectedName := "Name"
	expectedValue := 3.1415
	expectedCount := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		info := &trackingInformation{}
		rtl := newRecordingTraceListener(info)

		AddListener(&rtl)
		actualCount := len(traceListeners)

		if actualCount == expectedCount {
			t.Logf("\t\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedCount, actualCount)
		}

		t.Log("\tWhen a metric trace message is sent")
		{
			TraceMetric(expectedName, expectedValue)
			actualName := info.name
			actualValue := info.value

			if actualName == expectedName && actualValue == expectedValue {
				t.Logf("\t\t[%v] The name and value are correctly passed to the underlying trace listener.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The name and value are correctly passed to the underlying trace listener. Expected name: \"%v\", Actual name: \"%v\"; Expected value: \"%v\", Actual value: \"%v\"", ballotX, expectedName, actualName, expectedValue, actualValue)
			}
		}
	}
}

func TestEnsureTraceEventDataIsPassedToListener(t *testing.T) {
	expectedName := "Name"
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		info := &trackingInformation{}
		rtl := newRecordingTraceListener(info)

		AddListener(&rtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a event trace message is sent")
		{
			TraceEvent(expectedName)
			actualName := info.name

			if actualName == expectedName {
				t.Logf("\t\t[%v] The name is correctly passed to the underlying trace listener.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The name is correctly passed to the underlying trace listener. Expected name: \"%v\", Actual name: \"%v\"", ballotX, expectedName, actualName)
			}
		}
	}
}

type testError struct {
	err string
}

func (te *testError) Error() string {
	return te.err
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

type trackingInformation struct {
	message  string
	severity Severity
	err      error
	rethrow  bool
	name     string
	value    float64
}

type recordingTraceListener struct {
	info *trackingInformation
}

func newRecordingTraceListener(info *trackingInformation) TraceListener {
	return &recordingTraceListener{info: info}
}

func (rtl *recordingTraceListener) TraceMessage(message string, severity Severity) {
	rtl.info.message = message
	rtl.info.severity = severity
}

func (rtl *recordingTraceListener) TraceException(err error) {
	rtl.info.err = err
}

func (rtl *recordingTraceListener) TracePanic(rethrow bool) {
	rtl.info.rethrow = rethrow
}

func (rtl *recordingTraceListener) TrackAvailability(name string) *DurationTrace { return nil }

func (rtl *recordingTraceListener) TrackRequest(method string, uri string) *DurationTrace {
	return nil
}

func (rtl *recordingTraceListener) TrackDependency(name string, dependencyType string, target string) *DurationTrace {
	return nil
}

func (rtl *recordingTraceListener) TraceMetric(name string, value float64) {
	rtl.info.name = name
	rtl.info.value = value
}

func (rtl *recordingTraceListener) TraceEvent(name string) {
	rtl.info.name = name
}

func (rtl *recordingTraceListener) Flush() {}

func (rtl *recordingTraceListener) Close() {}
