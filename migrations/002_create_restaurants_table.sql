-- Migration: Create restaurants table
CREATE TABLE IF NOT EXISTS restaurants (
    id              SERIAL          PRIMARY KEY,
    name            VARCHAR(255)    NOT NULL,
    address         VARCHAR(255)    NOT NULL,
    cuisine_type    VARCHAR(100)    NOT NULL,
    rating          FLOAT DEFAULT   0
);
