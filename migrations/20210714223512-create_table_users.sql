-- +migrate Up
CREATE TABLE IF NOT EXISTS "users"
(
    "id"             char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "first_name"     varchar(100) NOT NULL,
    "last_name"      varchar(100) NOT NULL,
    "email"          varchar(100) NOT NULL,
    "username"       varchar(100) NOT NULL,
    "password"       varchar(128) NOT NULL,
    "address"        text,
    "phone_number"   char(15),
    "deposit_amount" numeric,
    "role_id"        int          NOT NULL,
    "created_at"     timestamp    NOT NULL,
    "updated_at"     timestamp    NOT NULL,
    "deleted_at"     timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "users";