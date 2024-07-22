-- +goose Up
-- +goose StatementBegin
-- Создание таблицы pictures
CREATE TABLE pictures (
                          project_id INTEGER NOT NULL,
                          url VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pictures
-- +goose StatementEnd
