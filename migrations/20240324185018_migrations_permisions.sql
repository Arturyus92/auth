-- +goose Up
CREATE TABLE IF NOT EXISTS permissions(
    id INT GENERATED ALWAYS AS IDENTITY NOT NULL,
    role INT,
    path TEXT,

    CONSTRAINT pk_permissions_id PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE IF EXISTS permissions;
