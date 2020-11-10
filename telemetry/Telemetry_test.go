package telemetry

import (
	"testing"
	"time"
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
				t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
			} else {
				t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
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
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
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
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
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
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
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
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
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
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
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
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
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
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedCount, actualCount)
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
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedCount, actualCount)
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
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
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

func TestEnsureDurationTraceIsCompletedAndDone(t *testing.T) {
	expectedStatusCode := "OK"
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		trace := newTrackingTraceInformation()
		dtl := newDurationTraceListener(&trace)

		AddListener(&dtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a duration-related message is sent")
		{
			dt := TrackRequest("GET", "/api/test")

			t.Log("\t and the duration tracking is completed and marked as done")
			{
				(*dt).Complete()
				(*dt).Done()

				if trace.success {
					t.Logf("\t\t[%v] The duration trace is marked as a success.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace is marked as a success.", ballotX)
				}

				if trace.statusCode == expectedStatusCode {
					t.Logf("\t\t[%v] The duration trace status code is set correctly.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace status code is set correctly. Expected: '%v', Actual: '%v'", ballotX, expectedStatusCode, trace.statusCode)
				}

				if trace.completed {
					t.Logf("\t\t[%v] The duration trace is marked as done.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace is marked as done.", ballotX)
				}
			}
		}
	}
}

func TestEnsureDurationTraceIsFailedAndDone(t *testing.T) {
	expectedStatusCode := "Bad Request"
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		trace := newTrackingTraceInformation()
		dtl := newDurationTraceListener(&trace)

		AddListener(&dtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a duration-related message is sent")
		{
			dt := TrackRequest("GET", "/api/test")

			t.Log("\t and the duration tracking is failed and marked as done")
			{
				(*dt).Fail(expectedStatusCode)
				(*dt).Done()

				if !trace.success {
					t.Logf("\t\t[%v] The duration trace is marked as a failure.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace is marked as a failure.", ballotX)
				}

				if trace.statusCode == expectedStatusCode {
					t.Logf("\t\t[%v] The duration trace status code is set correctly.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace status code is set correctly. Expected: '%v', Actual: '%v'", ballotX, expectedStatusCode, trace.statusCode)
				}

				if trace.completed {
					t.Logf("\t\t[%v] The duration trace is marked as done.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace is marked as done.", ballotX)
				}
			}
		}
	}
}

func TestEnsureDurationTraceIsCompletedAndNotDone(t *testing.T) {
	expectedStatusCode := "OK"
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		trace := newTrackingTraceInformation()
		dtl := newDurationTraceListener(&trace)

		AddListener(&dtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a duration-related message is sent")
		{
			dt := TrackRequest("GET", "/api/test")

			t.Log("\t and the duration tracking is completed and not marked as done")
			{
				(*dt).Complete()

				if trace.success {
					t.Logf("\t\t[%v] The duration trace is marked as a success.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace is marked as a success.", ballotX)
				}

				if trace.statusCode == expectedStatusCode {
					t.Logf("\t\t[%v] The duration trace status code is set correctly.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace status code is set correctly. Expected: '%v', Actual: '%v'", ballotX, expectedStatusCode, trace.statusCode)
				}

				if !trace.completed {
					t.Logf("\t\t[%v] The duration trace is not marked as done.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace is not marked as done.", ballotX)
				}
			}
		}
	}
}

func TestEnsureDurationTraceIsFailedAndNotDone(t *testing.T) {
	expectedStatusCode := "Bad Request"
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		trace := newTrackingTraceInformation()
		dtl := newDurationTraceListener(&trace)

		AddListener(&dtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a duration-related message is sent")
		{
			dt := TrackRequest("GET", "/api/test")

			t.Log("\t and the duration tracking is failed and not marked as done")
			{
				(*dt).Fail(expectedStatusCode)

				if !trace.success {
					t.Logf("\t\t[%v] The duration trace is marked as a failure.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace is marked as a failure.", ballotX)
				}

				if trace.statusCode == expectedStatusCode {
					t.Logf("\t\t[%v] The duration trace status code is set correctly.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace status code is set correctly. Expected: '%v', Actual: '%v'", ballotX, expectedStatusCode, trace.statusCode)
				}

				if !trace.completed {
					t.Logf("\t\t[%v] The duration trace is not marked as done.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace is not marked as done.", ballotX)
				}
			}
		}
	}
}

func TestEnsureDurationTraceIsNotCompletedAndNotDone(t *testing.T) {
	expectedStatusCode := "Incomplete"
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		trace := newTrackingTraceInformation()
		dtl := newDurationTraceListener(&trace)

		AddListener(&dtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a duration-related message is sent")
		{
			TrackRequest("GET", "/api/test")

			t.Log("\t and the duration tracking is not completed and not marked as done")
			{
				if !trace.success {
					t.Logf("\t\t[%v] The duration trace is not marked as a success.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace is not marked as a success.", ballotX)
				}

				if trace.statusCode == expectedStatusCode {
					t.Logf("\t\t[%v] The duration trace status code is set correctly.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace status code is set correctly. Expected: '%v', Actual: '%v'", ballotX, expectedStatusCode, trace.statusCode)
				}

				if !trace.completed {
					t.Logf("\t\t[%v] The duration trace is not marked as done.", checkMark)
				} else {
					t.Errorf("\t\t[%v] The duration trace is not marked as done.", ballotX)
				}
			}
		}
	}
}

func TestEnsureDurationTraceIsTransferringAvailabilityData(t *testing.T) {
	expectedName := "Expected Name"
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		trace := newTrackingTraceInformation()
		dtl := newDurationTraceListener(&trace)

		AddListener(&dtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a TrackAvailability message is sent")
		{
			TrackAvailability(expectedName)

			if trace.name == expectedName {
				t.Logf("\t\t[%v] The duration trace name is set correctly.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The duration trace name is set correctly. Expected: '%v', Actual: '%v'", ballotX, expectedName, trace.statusCode)
			}
		}
	}
}

func TestEnsureDurationTraceIsTransferringRequestData(t *testing.T) {
	expectedMethod := "GET"
	expectedUri := "/api/example"
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		trace := newTrackingTraceInformation()
		dtl := newDurationTraceListener(&trace)

		AddListener(&dtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a TrackRequest message is sent")
		{
			TrackRequest(expectedMethod, expectedUri)

			if trace.method == expectedMethod {
				t.Logf("\t\t[%v] The duration trace method is set correctly.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The duration trace method is set correctly. Expected: '%v', Actual: '%v'", ballotX, expectedMethod, trace.method)
			}

			if trace.uri == expectedUri {
				t.Logf("\t\t[%v] The duration trace URI is set correctly.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The duration trace URI is set correctly. Expected: '%v', Actual: '%v'", ballotX, expectedUri, trace.uri)
			}
		}
	}
}

func TestEnsureDurationTraceIsTransferringDependencyData(t *testing.T) {
	expectedName := "Expected Name"
	expectedDependencyType := "Dependency Type"
	expectedTarget := "Target"
	expectedValue := 1

	defer Close()

	t.Log("Given an implementation of the TraceListener interface")
	{
		trace := newTrackingTraceInformation()
		dtl := newDurationTraceListener(&trace)

		AddListener(&dtl)
		actualValue := len(traceListeners)

		if actualValue == expectedValue {
			t.Logf("\t[%v] There should only be one trace listener in the global list of listeners", checkMark)
		} else {
			t.Fatalf("\t[%v] There should only be one trace listener in the global list of listeners. Expected: %v, Actual: %v", ballotX, expectedValue, actualValue)
		}

		t.Log("\tWhen a TrackDependency message is sent")
		{
			TrackDependency(expectedName, expectedDependencyType, expectedTarget)

			if trace.name == expectedName {
				t.Logf("\t\t[%v] The duration trace name is set correctly.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The duration trace name is set correctly. Expected: '%v', Actual: '%v'", ballotX, expectedName, trace.statusCode)
			}

			if trace.dependencyType == expectedDependencyType {
				t.Logf("\t\t[%v] The duration trace dependency type is set correctly.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The duration trace dependency type is set correctly. Expected: '%v', Actual: '%v'", ballotX, expectedDependencyType, trace.dependencyType)
			}

			if trace.target == expectedTarget {
				t.Logf("\t\t[%v] The duration trace target is set correctly.", checkMark)
			} else {
				t.Errorf("\t\t[%v] The duration trace target is set correctly. Expected: '%v', Actual: '%v'", ballotX, expectedTarget, trace.target)
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

func (etl *emptyTraceListener) TrackAvailability(name string) *DurationTrace {
	return nil
}

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

func (rtl *recordingTraceListener) TrackAvailability(name string) *DurationTrace {
	return nil
}

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

type trackingTraceInformation struct {
	statusCode     string
	success        bool
	startTime      time.Time
	name           string
	dependencyType string
	target         string
	method         string
	uri            string
	completed      bool
	duration       time.Duration
}

func newTrackingTraceInformation() trackingTraceInformation {
	return trackingTraceInformation{
		statusCode: "Incomplete",
		success:    false,
		startTime:  time.Now(),
		completed:  false,
	}
}

// Complete indicates a successful completion of the measured duration activity
func (tti *trackingTraceInformation) Complete() {
	tti.success = true
	tti.statusCode = "OK"
}

// Fail indicates an unsuccessful completion of the measured duration activity
func (tti *trackingTraceInformation) Fail(statusCode string) {
	tti.success = false
	tti.statusCode = statusCode
}

// Done indicates that the trace is complete and should be committed to the telemetry source
func (tti *trackingTraceInformation) Done() {
	tti.duration = time.Now().Sub(tti.startTime)
	tti.completed = true
}

func newDurationTraceListener(tti *trackingTraceInformation) TraceListener {
	return &durationTraceListener{trace: tti}
}

type durationTraceListener struct {
	trace *trackingTraceInformation
}

func (dtl *durationTraceListener) TraceMessage(message string, severity Severity) {}

func (dtl *durationTraceListener) TraceException(err error) {}

func (dtl *durationTraceListener) TracePanic(rethrow bool) {}

func (dtl *durationTraceListener) TrackAvailability(name string) *DurationTrace {
	dtl.trace.name = name

	var trace DurationTrace = dtl.trace

	return &trace
}

func (dtl *durationTraceListener) TrackRequest(method string, uri string) *DurationTrace {
	dtl.trace.method = method
	dtl.trace.uri = uri

	var trace DurationTrace = dtl.trace

	return &trace
}

func (dtl *durationTraceListener) TrackDependency(name string, dependencyType string, target string) *DurationTrace {
	dtl.trace.name = name
	dtl.trace.dependencyType = dependencyType
	dtl.trace.target = target

	var trace DurationTrace = dtl.trace

	return &trace
}

func (dtl *durationTraceListener) TraceMetric(name string, value float64) {}

func (dtl *durationTraceListener) TraceEvent(name string) {}

func (dtl *durationTraceListener) Flush() {}

func (dtl *durationTraceListener) Close() {}
