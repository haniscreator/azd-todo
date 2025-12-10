-- +goose Up
CREATE TABLE IF NOT EXISTS todos (
  id UInt64,
  username String,
  title String,
  is_completed UInt8,
  created_at DateTime,
  completed_at Nullable(DateTime)
)
ENGINE = MergeTree()
ORDER BY (username, id);

-- +goose Down
DROP TABLE IF EXISTS todos;
