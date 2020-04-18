-- +goose Up
CREATE TABLE users (
  id bigserial NOT NULL PRIMARY KEY,
  email character varying NOT NULL UNIQUE,
  encrypted_password character varying(128) NOT NULL,
  confirmation_token character varying(128),
  remember_token character varying(128) NOT NULL,
  created_at timestamp(6) without time zone NOT NULL,
  updated_at timestamp(6) without time zone NOT NULL
);

CREATE UNIQUE INDEX index_users_on_email ON public.users USING btree (email);
CREATE INDEX index_users_on_remember_token ON public.users USING btree (remember_token);

-- +goose Down
DROP TABLE users;

