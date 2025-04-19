-- +goose Up
-- +goose StatementBegin
CREATE TABLE foo (
  username TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE foo;
-- +goose StatementEnd
