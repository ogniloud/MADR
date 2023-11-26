CREATE TABLE IF NOT EXISTS user_credentials (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    salt VARCHAR(100) NOT NULL,
    hash VARCHAR(150) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS deck_config (
    deck_id SERIAL PRIMARY KEY,
    user_id SERIAL,
    name VARCHAR(50) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user_credentials(user_id)
);

CREATE TABLE IF NOT EXISTS flashcard (
    card_id SERIAL PRIMARY KEY,
    word VARCHAR(100) NOT NULL,
    backside TEXT NOT NULL,
    deck_id SERIAL,
    answer VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS user_info (
    user_id SERIAL PRIMARY KEY,
    max_box INT NOT NULL CHECK (max_box > 0),
    FOREIGN KEY (user_id) REFERENCES user_credentials(user_id)
);

CREATE TABLE IF NOT EXISTS user_leitner (
    leitner_id SERIAL PRIMARY KEY,
    user_id SERIAL,
    card_id SERIAL,
    box INT NOT NULL,
    cool_down TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user_credentials(user_id),
    FOREIGN KEY (card_id) REFERENCES flashcard(card_id)
)