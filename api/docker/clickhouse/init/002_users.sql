USE analytics;

CREATE TABLE IF NOT EXISTS users
(
    id UUID,
    name String,
    email String,
    password_hash String,
    created_at DateTime64(3)
)
ENGINE = MergeTree
ORDER BY id;