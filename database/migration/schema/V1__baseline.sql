CREATE DATABASE IF NOT EXISTS sayaka;
use sayaka;

CREATE TABLE users
(
    id           INT AUTO_INCREMENT,
    line_user_id VARCHAR(255) NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    photo_url    VARCHAR(255),
    created_at   DATETIME     NOT NULL,
    updated_at   DATETIME     NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE flash_cards
(
    id         INT AUTO_INCREMENT,
    user_id    INT          NOT NULL,
    front      VARCHAR(255) NOT NULL,
    back       VARCHAR(255) NOT NULL,
    created_at DATETIME     NOT NULL,
    updated_at DATETIME     NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY fk_user_id (user_id) REFERENCES users (id) ON DELETE CASCADE
);
