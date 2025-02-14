-- Tabel payments: Menyimpan informasi transaksi pembayaran
CREATE TABLE IF NOT EXISTS payments (
    id              SERIAL          PRIMARY KEY,
    order_id        INT             NOT NULL,
    payment_method  TEXT            NOT NULL CHECK (payment_method IN ('e-wallet', 'bank transfer', 'COD')),
    payment_status  TEXT            NOT NULL DEFAULT 'pending' CHECK (payment_status IN ('pending', 'paid', 'failed', 'refunded')),
    amount          DECIMAL(10,2)   NOT NULL,
    discount_amount DECIMAL(10,2)   DEFAULT 0.00 CHECK (discount_amount >= 0),
    final_amount    DECIMAL(10,2)   GENERATED ALWAYS AS (amount - discount_amount) STORED,
    paid_at         TIMESTAMP       NULL,
    created_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE
);
