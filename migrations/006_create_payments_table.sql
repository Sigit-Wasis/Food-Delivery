-- Tabel payments: Menyimpan informasi transaksi pembayaran
CREATE TABLE IF NOT EXISTS payments (
    id              SERIAL          PRIMARY KEY,
    order_id        INT             NOT NULL,
    payment_method  ENUM('e-wallet', 'bank transfer', 'COD') NOT NULL,
    payment_status  ENUM('pending', 'paid', 'failed', 'refunded') NOT NULL DEFAULT 'pending',
    amount          DECIMAL(10,2)   NOT NULL,
    discount_amount DECIMAL(10,2)   DEFAULT 0.00 CHECK (discount_amount >= 0),
    final_amount    DECIMAL(10,2)   GENERATED ALWAYS AS (amount - discount_amount) STORED,
    paid_at         TIMESTAMP       NULL,
    created_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
);
