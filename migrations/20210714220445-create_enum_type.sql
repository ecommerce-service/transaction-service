
-- +migrate Up
CREATE TYPE "money_flow" AS ENUM (
    'in',
    'out'
    );

CREATE TYPE "transaction_type" AS ENUM (
    'cart',
    'on_going',
    'success',
    'canceled'
    );

-- +migrate Down
