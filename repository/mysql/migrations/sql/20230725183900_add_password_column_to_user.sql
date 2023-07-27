
-- +migrate Up
ALTER TABLE users ADD COLUMN password VARCHAR(255) NOT NULL;


-- +migrate Down
DROP TABLE users DROP COLUMN password;

