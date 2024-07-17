-- +goose Up
-- +goose StatementBegin
CREATE TABLE partners (
                          id SERIAL PRIMARY KEY,
                          url VARCHAR NOT NULL,
                          description TEXT
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE partners
-- +goose StatementEnd
