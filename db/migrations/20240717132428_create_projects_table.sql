-- +goose Up
-- +goose StatementBegin
-- Создание таблицы projects
CREATE TABLE projects (
                          id SERIAL PRIMARY KEY,
                          full_name VARCHAR NOT NULL,
                          url VARCHAR,
                          status VARCHAR,
                          short_name VARCHAR NOT NULL,
                          builder_name VARCHAR NOT NULL,
                          body TEXT,
                          coordinates point
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE projects

-- +goose StatementEnd
