-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name  VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    password VARCHAR(225) NOT NULL,
    email VARCHAR(225) UNIQUE NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL,
    token VARCHAR(225),
    refresh_token VARCHAR(225),
    user_id VARCHAR(225) NOT NULL,
    user_cart JSONB,
    address_details JSONB,
    order_status JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE users;