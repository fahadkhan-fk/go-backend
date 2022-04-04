-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id bytea PRIMARY KEY NOT NULL,
  user_name VARCHAR (30),
  password TEXT,
  email VARCHAR(50),
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ,
  deleted_at TIMESTAMPTZ
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
