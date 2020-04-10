-- +goose Up
ALTER TABLE entries ADD COLUMN notes text;

-- +goose Down
ALTER TABLE entries DROP COLUMN notes;

