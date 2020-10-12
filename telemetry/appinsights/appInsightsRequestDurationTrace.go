package appinsights

import (
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type applicationInsightsRequestDurationTrace struct {
	client     *appinsights.TelemetryClient
	statusCode string
	success    bool
	startTime  time.Time
	method     string
	uri        string
}

func (airdt *applicationInsightsRequestDurationTrace) Complete() {
	airdt.success = true
	airdt.statusCode = "OK"
}

func (airdt *applicationInsightsRequestDurationTrace) Fail(statusCode string) {
	airdt.success = false
	airdt.statusCode = statusCode
}

func (airdt *applicationInsightsRequestDurationTrace) Done() {
	endTime := time.Now()
	track := appinsights.NewRequestTelemetry(airdt.method, airdt.uri, endTime.Sub(airdt.startTime), airdt.statusCode)
	track.Success = airdt.success
	track.MarkTime(airdt.startTime, endTime)

	(*airdt.client).Track(track)
}
