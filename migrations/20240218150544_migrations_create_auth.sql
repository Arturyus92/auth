-- +goose Up
CREATE TABLE IF NOT EXISTS auth(
    user_id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL,
    role INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP,

    CONSTRAINT pk_auth_user_id PRIMARY KEY(user_id)
);

-- +goose Down
DROP TABLE IF EXISTS auth;
