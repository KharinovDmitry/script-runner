-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS commands (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    text TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS launches (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    command_id BIGINT,
    output text,
    FOREIGN KEY (command_id) REFERENCES commands(id)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE commands;
DROP TABLE launches;
