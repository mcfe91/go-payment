CREATE USER 'auth_user'@'localhost' IDENTIFIED BY "Auth123";

CREATE DATABASE auth;

GRANT ALL PRIVLEGES ON auth.* TO 'auth_user'@'localhost';

USER auth;

CREATE TABLE user (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
)

INSERT INTO user (user_id, password) VALUES ('mmcferren91@gmail.com', 'Admin123')