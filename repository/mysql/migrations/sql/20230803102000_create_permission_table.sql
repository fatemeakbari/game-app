
-- +migrate Up
CREATE TABLE permissions(
                                id INT PRIMARY KEY AUTO_INCREMENT,
                                title VARCHAR(255) NOT NULL UNIQUE
);


-- +migrate Down

DROP TABLE permissions;