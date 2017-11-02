package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// https://prometheus.io/docs/practices/naming/ is what we are basing naming conventions off of.
var (
	// CreateFailures is a metric for the number of failures to create a release
	CreateFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of failed create events",
		Name:      "create_error_total",
		Namespace: "releases",
	})

	// CreatedReleases is a metric for the number of releases created by this operator
	CreatedReleases = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of successful create events",
		Name:      "create_total",
		Namespace: "releases",
	})

	// DeleteFailures is a metric for the number of times a delete by this operator
	DeleteFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of failed delete events",
		Name:      "delete_error_total",
		Namespace: "releases",
	})

	// DeletedReleases is a metric for the number of releases deleted by this operator
	DeletedReleases = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of successful delete events",
		Name:      "delete_total",
		Namespace: "releases",
	})

	// ManagedReleases is a metric of the number of current releases managed by the operator
	ManagedReleases = prometheus.NewGauge(prometheus.GaugeOpts{
		Help:      "The number of releases managed by this operator",
		Name:      "total",
		Namespace: "releases",
	})

	// UpdateFailures is a metric for the number of times an update operation failed by this operator
	UpdateFailures = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of failed update events",
		Name:      "update_error_total",
		Namespace: "releases",
	})

	// UpdatedReleases is a metric for the number of releases updated successfully
	UpdatedReleases = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of successful update events",
		Name:      "update_total",
		Namespace: "releases",
	})

	// TotalEvents is a metric for the number of events that have been handled by this operator
	TotalEvents = prometheus.NewCounter(prometheus.CounterOpts{
		Help:      "The number of events (create/delete/updates) processed by this operator",
		Name:      "events_total",
		Namespace: "releases",
	})
)

func init() {
	prometheus.MustRegister(CreatedReleases)
	prometheus.MustRegister(CreateFailures)
	prometheus.MustRegister(DeletedReleases)
	prometheus.MustRegister(DeleteFailures)
	prometheus.MustRegister(ManagedReleases)
	prometheus.MustRegister(UpdatedReleases)
	prometheus.MustRegister(UpdateFailures)
	prometheus.MustRegister(TotalEvents)
}
