-- +goose Up
-- +goose StatementBegin
-- Создание таблицы docs
CREATE TABLE docs (
                      id SERIAL PRIMARY KEY,
                      preview_url VARCHAR NOT NULL,
                      document_url VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE docs
-- +goose StatementEnd
