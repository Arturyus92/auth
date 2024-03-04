-- +goose Up
CREATE TABLE IF NOT EXISTS log_transaction(
    id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    info TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT pk_log_transaction_id PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE IF EXISTS log_transaction;
