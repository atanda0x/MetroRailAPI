-- +goose Up

CREATE TABLE payments (
    order_id UUID PRIMARY KEY,
    digital BOOLEAN NOT NULL,
    cod BOOLEAN NOT NULL
);

-- +goose Down
DROP TABLE payments;