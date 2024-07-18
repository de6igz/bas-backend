-- +goose Up
-- +goose StatementBegin
-- Создание таблицы projects
CREATE TABLE projects (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR NOT NULL,
                          url VARCHAR,
                          status VARCHAR,
                          project_name VARCHAR NOT NULL,
                          builder_name VARCHAR NOT NULL,
                          body TEXT,
                          coordinates VARCHAR
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE projects

-- +goose StatementEnd
