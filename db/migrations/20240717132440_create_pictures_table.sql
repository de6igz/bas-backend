-- +goose Up
-- +goose StatementBegin
-- Создание таблицы pictures
CREATE TABLE pictures (
                          id SERIAL PRIMARY KEY,
                          project_id INTEGER NOT NULL,
                          url VARCHAR NOT NULL,
                          FOREIGN KEY (project_id) REFERENCES projects (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pictures
-- +goose StatementEnd
