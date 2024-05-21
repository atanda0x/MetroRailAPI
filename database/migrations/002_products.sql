-- +goose Up
CREATE TABLE products (
    id UUID PRIMARY KEY,
    productName VARCHAR(225),
    price BIGINT,
    rating SMALLINT,
    image VARCHAR(225)
);

-- +goose Down
DROP TABLE products;