package appinsights

import (
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type applicationInsightsAvailabilityDurationTrace struct {
	client     *appinsights.TelemetryClient
	statusCode string
	success    bool
	startTime  time.Time
	name       string
}

func (aiadt *applicationInsightsAvailabilityDurationTrace) Complete() {
	aiadt.success = true
	aiadt.statusCode = "OK"
}

func (aiadt *applicationInsightsAvailabilityDurationTrace) Fail(statusCode string) {
	aiadt.success = false
	aiadt.statusCode = statusCode
}

func (aiadt *applicationInsightsAvailabilityDurationTrace) Done() {
	endTime := time.Now()
	track := appinsights.NewAvailabilityTelemetry(aiadt.name, endTime.Sub(aiadt.startTime), aiadt.success)
	track.Message = aiadt.statusCode
	track.MarkTime(aiadt.startTime, endTime)

	(*aiadt.client).Track(track)
}
