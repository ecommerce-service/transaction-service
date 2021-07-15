-- +migrate Up

CREATE TABLE IF NOT EXISTS "car_types"
(
    "id"         char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
    "brand_id"   char(36)    NOT NULL,
    "name"       varchar(50) NOT NULL,
    "created_at" timestamp   NOT NULL,
    "updated_at" timestamp   NOT NULL,
    "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "car_types";