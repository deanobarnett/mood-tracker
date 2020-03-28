-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
-- Add the UUID extension. Add this to its own setup migration.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE entries (
    id bigserial NOT NULL PRIMARY KEY,
    uuid uuid DEFAULT uuid_generate_v4(),
    date VARCHAR(10) NOT NULL,
    mood integer NOT NULL,
    stress integer NOT NULL,
    sleep integer NOT NULL,
    created_at timestamp(6) without time zone NOT NULL,
    updated_at timestamp(6) without time zone NOT NULL,
    deleted_at timestamp(6) without time zone
);

CREATE INDEX index_entries_on_date ON entries USING btree (date);
CREATE INDEX index_entries_on_uuid ON entries USING btree (uuid);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE entries;
