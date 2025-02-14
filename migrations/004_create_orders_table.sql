-- Tabel orders: Menyimpan informasi utama pesanan
CREATE TABLE IF NOT EXISTS orders (
    id              SERIAL          PRIMARY KEY,
    user_id         INT             NOT NULL,
    restaurant_id   INT             NOT NULL,
    driver_id       INT             DEFAULT NULL,
    total_price     DECIMAL(10,2)   NOT NULL,
    payment_status  TEXT            NOT NULL DEFAULT 'pending' CHECK (payment_status IN ('pending', 'paid', 'failed')),
    order_status    TEXT            NOT NULL DEFAULT 'new' CHECK (order_status IN ('new', 'processed', 'delivered', 'cancelled')),
    created_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id) ON DELETE CASCADE,
    FOREIGN KEY (driver_id) REFERENCES users(id) ON DELETE SET NULL
);
