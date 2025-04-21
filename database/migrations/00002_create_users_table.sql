-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  username TEXT,
  password TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE baz;
-- +goose StatementEnd
