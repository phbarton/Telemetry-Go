package appinsights

import (
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type applicationInsightsDependencyDurationTrace struct {
	client         *appinsights.TelemetryClient
	statusCode     string
	success        bool
	startTime      time.Time
	name           string
	dependencyType string
	target         string
}

func (aiddt *applicationInsightsDependencyDurationTrace) Complete() {
	aiddt.success = true
	aiddt.statusCode = "OK"
}

func (aiddt *applicationInsightsDependencyDurationTrace) Fail(statusCode string) {
	aiddt.success = false
	aiddt.statusCode = statusCode
}

func (aiddt *applicationInsightsDependencyDurationTrace) Done() {
	endTime := time.Now()
	track := appinsights.NewRemoteDependencyTelemetry(aiddt.name, aiddt.dependencyType, aiddt.target, aiddt.success)
	track.ResultCode = aiddt.statusCode
	track.MarkTime(aiddt.startTime, endTime)

	(*aiddt.client).Track(track)
}
