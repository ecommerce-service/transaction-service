-- +migrate Up
CREATE TABLE IF NOT EXISTS "transaction_details"
(
    "id"              char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "transaction_id"  char(36)    NOT NULL,
    "car_id"          char(36)    NOT NULL,
    "car_brand"       varchar(50) NOT NULL,
    "car_type"        varchar(50) NOT NULL,
    "car_color"      varchar(50) NOT NULL,
    "production_year" char(4)        NOT NULL,
    "price"           numeric     NOT NULL,
    "quantity"        numeric     NOT NULL,
    "sub_total"       numeric     NOT NULL,
    "created_at"      timestamp   NOT NULL,
    "updated_at"      timestamp   NOT NULL,
    "deleted_at"      timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "transaction_details";