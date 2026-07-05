USE analytics;

CREATE TABLE IF NOT EXISTS pipeline_stage_metrics
(
    execution_id String,

    stage String,

    started_at DateTime64(3),
    finished_at DateTime64(3),

    latency_ms Float64,

    input UInt32,
    output UInt32,

    dropped UInt32,
    errors UInt32,

    throughput Float64
)
ENGINE = MergeTree
ORDER BY (started_at, stage);