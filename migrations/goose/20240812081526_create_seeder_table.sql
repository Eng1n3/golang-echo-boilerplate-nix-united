-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE seeder (
    version_id INTEGER NOT NULL,
    is_applied BOOLEAN NOT NULL DEFAULT(TRUE),
    tstamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (version_id) CONSTRAINT PK_seeder
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS seeder;
-- +goose StatementEnd
