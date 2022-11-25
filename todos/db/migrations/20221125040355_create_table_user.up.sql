CREATE TABLE users (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    username VARCHAR(100) NOT NULL,
    password TEXT NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    active TINYINT NOT NULL DEFAULT(1),
    created_at datetime,
    updated_at datetime
)ENGINE = InnoDB;