-- Tabel order_items: Menyimpan detail pesanan per item menu
CREATE TABLE IF NOT EXISTS order_items (
    id              SERIAL          PRIMARY KEY,
    order_id        INT             NOT NULL,
    menu_id         INT             NOT NULL,
    quantity        INT             NOT NULL CHECK (quantity > 0),
    price           DECIMAL(10,2)   NOT NULL,
    subtotal        DECIMAL(10,2)   GENERATED ALWAYS AS (quantity * price) STORED,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (menu_id) REFERENCES menu_items(id) ON DELETE CASCADE
);