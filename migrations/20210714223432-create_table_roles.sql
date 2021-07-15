
-- +migrate Up
CREATE TABLE IF NOT EXISTS "roles" (
                         "id" SERIAL PRIMARY KEY NOT NULL,
                         "name" varchar(100) NOT NULL,
                         "created_at" timestamp NOT NULL,
                         "updated_at" timestamp NOT NULL,
                         "deleted_at" timestamp
);
-- +migrate Down
DROP TABLE IF EXISTS "roles";