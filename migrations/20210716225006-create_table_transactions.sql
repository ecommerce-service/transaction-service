-- +migrate Up
CREATE TABLE IF NOT EXISTS "transactions"
(
    "id"                 char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "user_id"            char(36)  NOT NULL,
    "transaction_type"   transaction_type,
    "transaction_number" char(14),
    "total_amount"       numeric   NOT NULL,
    "payment_received"   numeric,
    "created_at"         timestamp NOT NULL,
    "updated_at"         timestamp NOT NULL,
    "paid_at"            timestamp,
    "canceled_at"        timestamp,
    "deleted_at"         timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "transactions";