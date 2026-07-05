package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	EventsReceived = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "events_received_total",
			Help: "Events accepted",
		},
	)

	EventsProcessed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "events_processed_total",
			Help: "Events processed",
		},
	)

	EventsFailed = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "events_failed_total",
			Help: "Failed events",
		},
	)

	BatchInsertFailures = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "batch_insert_failures_total",
			Help: "Failed batch inserts",
		},
	)

	QueueDepth = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "event_queue_depth",
			Help: "Queue depth",
		},
	)

	BatchSize = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "batch_size",
			Buckets: prometheus.LinearBuckets(
				10,
				50,
				20,
			),
		},
	)

	InsertLatency = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "clickhouse_insert_seconds",
			Buckets: prometheus.DefBuckets,
		},
	)

	PipelineLatency = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "pipeline_seconds",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func Register() {
	prometheus.MustRegister(
		EventsReceived,
		EventsProcessed,
		EventsFailed,
		BatchInsertFailures,
		QueueDepth,
		BatchSize,
		InsertLatency,
		PipelineLatency,
	)
}
