-- +goose Up
CREATE TABLE IF NOT EXISTS key_tokens(
    id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    key TEXT NOT NULL,
    value TEXT NOT NULL,

    CONSTRAINT pk_key_tokens_id PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE IF EXISTS key_tokens;
