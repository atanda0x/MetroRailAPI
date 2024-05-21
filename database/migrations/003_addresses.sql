-- +goose Up
CREATE TABLE addresses (
    id UUID PRIMARY KEY,
    house VARCHAR(225) NOT NULL,
    street VARCHAR(225) NOT NULL,
    city VARCHAR(225) NOT NULL,
    pin_code VARCHAR(20) NOT NULL
);

-- +goose Down
DROP TABLE addresses;