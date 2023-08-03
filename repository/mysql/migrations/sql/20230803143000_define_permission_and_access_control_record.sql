
-- +migrate Up

INSERT INTO permissions(id, title) VALUES (1, 'user_list');


INSERT INTO access_controls(`actor_type`, `actor_id`, `permission_id`) VALUES('role', 2, 1);
