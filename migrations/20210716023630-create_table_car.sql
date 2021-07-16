-- +migrate Up
CREATE TABLE IF NOT EXISTS "cars"
(
    "id"              char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "car_type_id"     char(36)  NOT NULL,
    "car_color_id"    char(36)  NOT NULL,
    "production_year" char(4),
    "price"           numeric,
    "stock"           int,
    "created_at"      timestamp NOT NULL,
    "updated_at"      timestamp NOT NULL,
    "deleted_at"      timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "cars";