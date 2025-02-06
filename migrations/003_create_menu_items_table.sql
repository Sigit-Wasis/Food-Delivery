-- Migration: Create menu_items table
CREATE TABLE IF NOT EXISTS menu_items (
    id              SERIAL          PRIMARY KEY,
    restaurant_id   INT             NOT NULL,
    name            VARCHAR(255)    NOT NULL,
    description     TEXT            NOT NULL,
    price           DECIMAL(10,2)   NOT NULL,
    image           VARCHAR(255)    DEFAULT NULL,
    created_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id) ON DELETE CASCADE
);
