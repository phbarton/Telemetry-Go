package appinsights

import (
	"os"
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	"github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
	"github.com/phbarton/Telemetry-Go/telemetry"
)

type appInsightsTraceListener struct {
	client appinsights.TelemetryClient
}

// NewApplicationInsightsTraceListener creates a trace listener which outputs to the Azure ApplicationInsights instance specified by the provided
// instrumentation key. It limits output based on the logging level supplied
func NewApplicationInsightsTraceListener(service, version string, instrumentationKey string) telemetry.TraceListener {
	client := appinsights.NewTelemetryClient(instrumentationKey)
	host, _ := os.Hostname()

	client.Context().Tags.Cloud().SetRole(service)
	client.Context().Tags.Cloud().SetRoleInstance(host)
	client.Context().Tags.Application().SetVer(version)

	traceListener := &appInsightsTraceListener{client: client}
	return traceListener
}

func (aitl *appInsightsTraceListener) TraceMessage(message string, severity telemetry.Severity) {
	track := appinsights.NewTraceTelemetry(message, toAppInsightsSeverity(severity))
	track.Timestamp = time.Now()

	aitl.client.Track(track)
}

func (aitl *appInsightsTraceListener) TraceException(err error) {
	track := appinsights.NewExceptionTelemetry(err)
	track.SeverityLevel = contracts.Error
	track.Frames = appinsights.GetCallstack(0)

	aitl.client.Track(track)
}

func (aitl *appInsightsTraceListener) TracePanic(rethrow bool) {
	appinsights.TrackPanic(aitl.client, rethrow)
}

func (aitl *appInsightsTraceListener) TrackAvailability(name string) *telemetry.DurationTrace {
	trace := newAvailabilityDurationTrace(aitl, name)

	return &trace
}

func (aitl *appInsightsTraceListener) TrackRequest(method string, uri string) *telemetry.DurationTrace {
	trace := newRequestDurationTrace(aitl, method, uri)

	return &trace
}

func (aitl *appInsightsTraceListener) TrackDependency(name string, dependencyType string, target string) *telemetry.DurationTrace {
	trace := newDependencyDurationTrace(aitl, name, dependencyType, target)

	return &trace
}

func (aitl *appInsightsTraceListener) TraceMetric(name string, value float64) {
	track := appinsights.NewMetricTelemetry(name, value)

	aitl.client.Track(track)
}

func (aitl *appInsightsTraceListener) TraceEvent(name string) {
	track := appinsights.NewEventTelemetry(name)

	aitl.client.Track(track)
}

func (aitl *appInsightsTraceListener) Flush() {
	aitl.client.Channel().Flush()
}

func (aitl *appInsightsTraceListener) Close() {
	select {
	case <-aitl.client.Channel().Close(10 * time.Second):

	case <-time.After(30 * time.Second):
		aitl.client.Channel().Stop()
	}

	aitl.client.SetIsEnabled(false)
}

func newAvailabilityDurationTrace(aitl *appInsightsTraceListener, name string) telemetry.DurationTrace {
	return &applicationInsightsAvailabilityDurationTrace{
		client:    &aitl.client,
		name:      name,
		startTime: time.Now(),
	}
}

func newRequestDurationTrace(aitl *appInsightsTraceListener, method, uri string) telemetry.DurationTrace {
	return &applicationInsightsRequestDurationTrace{
		client:    &aitl.client,
		method:    method,
		uri:       uri,
		startTime: time.Now(),
	}
}

func newDependencyDurationTrace(aitl *appInsightsTraceListener, name, dependencyType, target string) telemetry.DurationTrace {
	return &applicationInsightsDependencyDurationTrace{
		client:         &aitl.client,
		name:           name,
		dependencyType: dependencyType,
		target:         target,
		startTime:      time.Now(),
	}
}

func toAppInsightsSeverity(severity telemetry.Severity) contracts.SeverityLevel {
	switch severity {
	case telemetry.Verbose:
		return contracts.Verbose
	case telemetry.Information:
		return contracts.Information
	case telemetry.Warning:
		return contracts.Warning
	case telemetry.Error:
		return contracts.Error
	case telemetry.Critical:
		return contracts.Critical
	default:
		return contracts.Information
	}
}
