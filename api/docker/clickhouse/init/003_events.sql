USE analytics;

CREATE TABLE IF NOT EXISTS events
(
    event_id UUID,

    event_type LowCardinality(String),

    event_name String,

    user_id String,

    session_id String,

    captured_at DateTime64(3),

    page_url String,

    page_path String,

    page_title String,

    properties String,

    ip_address String,

    country String,

    city String,

    user_agent String,

    referrer String,

    language String,

    sdk_version String,

    request_id String
)
ENGINE = MergeTree
ORDER BY
(
    captured_at,
    event_type,
    user_id
);