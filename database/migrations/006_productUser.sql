-- +goose Up
CREATE TABLE productUser (
    id UUID PRIMARY KEY,
    productName VARCHAR(225) NOT NULL,
    price BIGINT NOT NULL,
    rating SMALLINT NOT NULL,
    image VARCHAR(225) NOT NULL
);

-- +goose Down
DROP TABLE productUser;