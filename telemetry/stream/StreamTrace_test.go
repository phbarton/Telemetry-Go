package stream

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/phbarton/Telemetry-Go/telemetry"
)

const (
	checkMark = "\u2713"
	ballotX   = "\u2717"
)

type testWriter struct {
	f func(s string)
}

func (tw testWriter) Write(p []byte) (n int, err error) {
	tw.f(string(p))

	return len(p), nil
}

func newTestWriter(f func(s string)) io.Writer {
	tw := &testWriter{f: f}

	return tw
}

func TestTraceListenerDoesNotLogLowerServerities(t *testing.T) {
	testMessage := "Test message"
	expectedValue := ""
	actualMessage := ""
	tw := newTestWriter(func(s string) { actualMessage = s })

	t.Log("Given a StreamTraceListener with a minimum severity of 'Error'")
	{
		tl := NewStreamTraceListener(telemetry.Error, &tw)

		defer tl.Close()

		t.Log("\tWhen a 'Verbose' severity message is traced")
		{
			tl.TraceMessage(testMessage, telemetry.Verbose)
			time.Sleep(1 * time.Second) // Since the write is asynchronous, wait a bit for it to be handled

			if actualMessage == expectedValue {
				t.Logf("\t\t[%v] No message is written to the underlying stream.", checkMark)
			} else {
				t.Errorf("\t\t[%v] No message is written to the underlying stream. Expected: \"%v\", Actual: \"%v\"", ballotX, expectedValue, actualMessage)
			}
		}

		t.Log("\tWhen an 'Information' severity message is traced")
		{
			tl.TraceMessage(testMessage, telemetry.Information)
			time.Sleep(1 * time.Second) // Since the write is asynchronous, wait a bit for it to be handled

			if actualMessage == expectedValue {
				t.Logf("\t\t[%v] No message is written to the underlying stream.", checkMark)
			} else {
				t.Errorf("\t\t[%v] No message is written to the underlying stream. Expected: \"%v\", Actual: \"%v\"", ballotX, expectedValue, actualMessage)
			}
		}

		t.Log("\tWhen an 'Warning' severity message is traced")
		{
			tl.TraceMessage(testMessage, telemetry.Warning)
			time.Sleep(1 * time.Second) // Since the write is asynchronous, wait a bit for it to be handled

			if actualMessage == expectedValue {
				t.Logf("\t\t[%v] No message is written to the underlying stream.", checkMark)
			} else {
				t.Errorf("\t\t[%v] No message is written to the underlying stream. Expected: \"%v\", Actual: \"%v\"", ballotX, expectedValue, actualMessage)
			}
		}
	}
}

func TestTraceListenerWritesAtSameSeverityLevel(t *testing.T) {
	testMessage := "Test message"
	actualMessage := ""
	tw := newTestWriter(func(s string) { actualMessage = s })

	t.Log("Given a StreamTraceListener with a minimum severity of 'Error'")
	{
		tl := NewStreamTraceListener(telemetry.Error, &tw)

		defer tl.Close()

		t.Log("\tWhen an 'Error' severity message is traced")
		{
			tl.TraceMessage(testMessage, telemetry.Error)
			time.Sleep(1 * time.Second) // Since the write is asynchronous, wait a bit for it to be handled

			if actualMessage != "" {
				t.Logf("\t\t[%v] A message is written to the underlying stream.", checkMark)
			} else {
				t.Errorf("\t\t[%v] A message is written to the underlying stream. Expected: A non-empty string value, Actual: \"%v\"", ballotX, actualMessage)
			}
		}
	}
}

func TestTraceListenerWritesAtHigherSeverityLevel(t *testing.T) {
	testMessage := "Test message"
	actualMessage := ""
	tw := newTestWriter(func(s string) { actualMessage = s })

	t.Log("Given a StreamTraceListener with a minimum severity of 'Error'")
	{
		tl := NewStreamTraceListener(telemetry.Error, &tw)

		defer tl.Close()

		t.Log("\tWhen an 'Critical' severity message is traced")
		{
			tl.TraceMessage(testMessage, telemetry.Critical)
			time.Sleep(1 * time.Second) // Since the write is asynchronous, wait a bit for it to be handled

			if actualMessage != "" {
				t.Logf("\t\t[%v] A message is written to the underlying stream.", checkMark)
			} else {
				t.Errorf("\t\t[%v] A message is written to the underlying stream. Expected: A non-empty string value, Actual: \"%v\"", ballotX, actualMessage)
			}
		}
	}
}

func TestVerboseMessagesAreFormattedCorrectly(t *testing.T) {
	testMessage := "Test message"
	expectedMessage := fmt.Sprintf("[VRB]: %v", testMessage)
	actualMessage := ""
	tw := newTestWriter(func(s string) { actualMessage = s })

	t.Log("Given a StreamTraceListener with a minimum severity of 'Verbose'")
	{
		tl := NewStreamTraceListener(telemetry.Verbose, &tw)

		defer tl.Close()

		t.Log("\tWhen a 'Verbose' severity message is traced")
		{
			tl.TraceMessage(testMessage, telemetry.Verbose)
			time.Sleep(1 * time.Second) // Since the write is asynchronous, wait a bit for it to be handled

			actualMinusDate := strings.SplitN(actualMessage, " ", 4) // Get rid of the date/time since that can't be captured
			resultValue := strings.TrimSpace(actualMinusDate[3])

			if resultValue == expectedMessage {
				t.Logf("\t\t[%v] A properly formatted message is written to the underlying stream.", checkMark)
			} else {
				t.Errorf("\t\t[%v] A properly formatted message is written to the underlying stream. Expected: \"%v\", Actual: \"%v\"", ballotX, expectedMessage, resultValue)
			}
		}
	}
}

func TestInformationMessagesAreFormattedCorrectly(t *testing.T) {
	testMessage := "Test message"
	expectedMessage := fmt.Sprintf("[INF]: %v", testMessage)
	actualMessage := ""
	tw := newTestWriter(func(s string) { actualMessage = s })

	t.Log("Given a StreamTraceListener with a minimum severity of 'Verbose'")
	{
		tl := NewStreamTraceListener(telemetry.Verbose, &tw)

		defer tl.Close()

		t.Log("\tWhen a 'Information' severity message is traced")
		{
			tl.TraceMessage(testMessage, telemetry.Information)
			time.Sleep(1 * time.Second) // Since the write is asynchronous, wait a bit for it to be handled

			actualMinusDate := strings.SplitN(actualMessage, " ", 4) // Get rid of the date/time since that can't be captured
			resultValue := strings.TrimSpace(actualMinusDate[3])

			if resultValue == expectedMessage {
				t.Logf("\t\t[%v] A properly formatted message is written to the underlying stream.", checkMark)
			} else {
				t.Errorf("\t\t[%v] A properly formatted message is written to the underlying stream. Expected: \"%v\", Actual: \"%v\"", ballotX, expectedMessage, resultValue)
			}
		}
	}
}

func TestWarningMessagesAreFormattedCorrectly(t *testing.T) {
	testMessage := "Test message"
	expectedMessage := fmt.Sprintf("[WRN]: %v", testMessage)
	actualMessage := ""
	tw := newTestWriter(func(s string) { actualMessage = s })

	t.Log("Given a StreamTraceListener with a minimum severity of 'Verbose'")
	{
		tl := NewStreamTraceListener(telemetry.Verbose, &tw)

		defer tl.Close()

		t.Log("\tWhen a 'Warning' severity message is traced")
		{
			tl.TraceMessage(testMessage, telemetry.Warning)
			time.Sleep(1 * time.Second) // Since the write is asynchronous, wait a bit for it to be handled

			actualMinusDate := strings.SplitN(actualMessage, " ", 4) // Get rid of the date/time since that can't be captured
			resultValue := strings.TrimSpace(actualMinusDate[3])

			if resultValue == expectedMessage {
				t.Logf("\t\t[%v] A properly formatted message is written to the underlying stream.", checkMark)
			} else {
				t.Errorf("\t\t[%v] A properly formatted message is written to the underlying stream. Expected: \"%v\", Actual: \"%v\"", ballotX, expectedMessage, resultValue)
			}
		}
	}
}

func TestErrorMessagesAreFormattedCorrectly(t *testing.T) {
	testMessage := "Test message"
	expectedMessage := fmt.Sprintf("[ERR]: %v", testMessage)
	actualMessage := ""
	tw := newTestWriter(func(s string) { actualMessage = s })

	t.Log("Given a StreamTraceListener with a minimum severity of 'Verbose'")
	{
		tl := NewStreamTraceListener(telemetry.Verbose, &tw)

		defer tl.Close()

		t.Log("\tWhen a 'Error' severity message is traced")
		{
			tl.TraceMessage(testMessage, telemetry.Error)
			time.Sleep(1 * time.Second) // Since the write is asynchronous, wait a bit for it to be handled

			actualMinusDate := strings.SplitN(actualMessage, " ", 4) // Get rid of the date/time since that can't be captured
			resultValue := strings.TrimSpace(actualMinusDate[3])

			if resultValue == expectedMessage {
				t.Logf("\t\t[%v] A properly formatted message is written to the underlying stream.", checkMark)
			} else {
				t.Errorf("\t\t[%v] A properly formatted message is written to the underlying stream. Expected: \"%v\", Actual: \"%v\"", ballotX, expectedMessage, resultValue)
			}
		}
	}
}

func TestCriticalMessagesAreFormattedCorrectly(t *testing.T) {
	testMessage := "Test message"
	expectedMessage := fmt.Sprintf("[CRT]: %v", testMessage)
	actualMessage := ""
	tw := newTestWriter(func(s string) { actualMessage = s })

	t.Log("Given a StreamTraceListener with a minimum severity of 'Verbose'")
	{
		tl := NewStreamTraceListener(telemetry.Verbose, &tw)

		defer tl.Close()

		t.Log("\tWhen a 'Critical' severity message is traced")
		{
			tl.TraceMessage(testMessage, telemetry.Critical)
			time.Sleep(1 * time.Second) // Since the write is asynchronous, wait a bit for it to be handled

			actualMinusDate := strings.SplitN(actualMessage, " ", 4) // Get rid of the date/time since that can't be captured
			resultValue := strings.TrimSpace(actualMinusDate[3])

			if resultValue == expectedMessage {
				t.Logf("\t\t[%v] A properly formatted message is written to the underlying stream.", checkMark)
			} else {
				t.Errorf("\t\t[%v] A properly formatted message is written to the underlying stream. Expected: \"%v\", Actual: \"%v\"", ballotX, expectedMessage, resultValue)
			}
		}
	}
}
