CREATE USER 'ledger_user'@'%' IDENTIFIED BY 'Auth123';

CREATE DATABASE ledger;

GRANT ALL PRIVILEGES ON ledge.* TO 'ledger_user'@'%';

USE ledger;

CREATE TABLE `ledger` (
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    order_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    amount INT NOT NULL,
    operation VARCHAR(255) NOT NULL,
    transaction_date VARCHAR(255) NOT NULL
);