-- +goose Up

CREATE TABLE orders (
    id UUID PRIMARY KEY,
    price BIGINT NOT NULL,
    discount BIGINT NOT NULL,
    paymentMethod JSONB NOT NULL,
    order_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE orders;