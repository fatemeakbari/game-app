

CREATE TABLE users(
                        id INT PRIMARY KEY AUTO_INCREMENT,
                        name VARCHAR(255) NOT NULL,
                        phone_number VARCHAR(255) NOT NULL UNIQUE,
                        password VARCHAR(255) NOT NULL,
                        create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)