

-- +migrate Up
ALTER TABLE `users` ADD COLUMN `role` INT DEFAULT 1;

ALTER TABLE `users` MODIFY `role` INT NOT NULL;


-- +migrate Down
ALTER TABLE `users` DROP COLUMN `role`;