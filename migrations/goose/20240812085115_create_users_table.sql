-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE users (
    id INTEGER NOT NULL CONSTRAINT PK_users PRIMARY KEY AUTOINCREMENT,
    email VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT UQ_users_email UNIQUE (email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS users;
-- +goose StatementEnd
