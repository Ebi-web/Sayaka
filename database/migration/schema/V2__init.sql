CREATE TABLE users
(
    id           int generated always as identity primary key,
    line_user_id VARCHAR(255)            NOT NULL,
    display_name VARCHAR(255)            NOT NULL,
    photo_url    VARCHAR(255),
    created_at   TIMESTAMP DEFAULT now() NOT NULL,
    updated_at   TIMESTAMP DEFAULT now() NOT NULL,
    UNIQUE (line_user_id)
);

CREATE TABLE flash_cards
(
    id         int generated always as identity primary key,
    user_id    INT                     NOT NULL,
    front      VARCHAR(255)            NOT NULL,
    back       VARCHAR(255)            NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    updated_at TIMESTAMP DEFAULT now() NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
