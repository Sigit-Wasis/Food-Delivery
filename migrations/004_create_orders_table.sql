-- Tabel orders: Menyimpan informasi utama pesanan
CREATE TABLE IF NOT EXISTS orders (
    id              SERIAL          PRIMARY KEY,
    user_id        INT              NOT NULL,
    restaurant_id   INT             NOT NULL,
    driver_id       INT             DEFAULT NULL,
    total_price     DECIMAL(10,2)   NOT NULL,
    payment_status  ENUM('pending', 'paid', 'failed') NOT NULL DEFAULT 'pending',
    order_status    ENUM('new', 'processed', 'delivered', 'cancelled') NOT NULL DEFAULT 'new',
    created_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id) ON DELETE CASCADE,
    FOREIGN KEY (driver_id) REFERENCES users(id) ON DELETE SET NULL
);
