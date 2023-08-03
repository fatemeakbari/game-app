
-- +migrate Up
CREATE TABLE access_controls(
                      id INT PRIMARY KEY AUTO_INCREMENT,
                      actor_id INT NOT NULL,
                      actor_type VARCHAR(255) NOT NULL,
                      permission_id INT NOT NULL,
                      create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                      FOREIGN KEY (permission_id) REFERENCES permissions(id)
);


-- +migrate Down

DROP TABLE access_controls;