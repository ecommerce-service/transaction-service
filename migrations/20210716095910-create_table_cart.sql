-- +migrate Up
CREATE TABLE IF NOT EXISTS "carts" (
                         "id" char(36) PRIMARY KEY DEFAULT (uuid_generate_v4()),
                         "user_id" char(36) NOT NULL,
                         "car_id" char(36) NOT NULL,
                         "car_brand" varchar(50) NOT NULL,
                         "car_type" varchar(50) NOT NULL,
                         "car_color" varchar(50) NOT NULL,
                         "production_year" char(4) NOT NULL,
                         "price" numeric,
                         "quantity" numeric,
                         "sub_total" numeric,
                         "created_at" timestamp NOT NULL,
                         "updated_at" timestamp NOT NULL,
                         "deleted_at" timestamp
);

-- +migrate Down
DROP TABLE IF EXISTS "carts";