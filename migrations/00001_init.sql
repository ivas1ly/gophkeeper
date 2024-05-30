-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
  id uuid PRIMARY KEY,
  username TEXT UNIQUE NOT NULL CHECK (length(username) <= 255),
  password_hash TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
  deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS users_username_idx ON users (username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
