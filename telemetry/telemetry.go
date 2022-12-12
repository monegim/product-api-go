package telemetry

import (
	api "go.opentelemetry.io/otel/api/metric"
	"go.opentelemetry.io/otel/sdk/metric/controller/push"
)

type Telemetry struct {
	pusher   *push.Controller
	meter    api.Meter
	measures map[string]*api.Float64Measure
	counters map[string]*api.Float64Counter
}
