DROP DATABASE IF EXISTS sayaka;
CREATE DATABASE IF NOT EXISTS sayaka DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
use sayaka;

CREATE TABLE users
(
    id           INT AUTO_INCREMENT,
    line_user_id VARCHAR(255)                                                   NOT NULL,
    display_name VARCHAR(255)                                                   NOT NULL,
    photo_url    VARCHAR(255),
    created_at   DATETIME DEFAULT CURRENT_TIMESTAMP                             NOT NULL,
    updated_at   DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

CREATE TABLE flash_cards
(
    id         INT AUTO_INCREMENT,
    user_id    INT          NOT NULL,
    front      VARCHAR(255) NOT NULL,
    back       VARCHAR(255) NOT NULL,
    created_at   DATETIME DEFAULT CURRENT_TIMESTAMP                             NOT NULL,
    updated_at   DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY fk_user_id (user_id) REFERENCES users (id) ON DELETE CASCADE
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
