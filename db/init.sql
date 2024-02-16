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
);

-- Groups

CREATE TABLE IF NOT EXISTS groups (
    group_id SERIAL PRIMARY KEY,
    creator_id SERIAL,
    name VARCHAR(50) NOT NULL,
    time_created TIMESTAMP,

    FOREIGN KEY (creator_id) REFERENCES user_credentials(user_id)
);

CREATE TABLE IF NOT EXISTS group_decks (
    group_id SERIAL,
    deck_id SERIAL, -- колода, которую видят все члены группы
    time_shared TIMESTAMP, -- время шаринга колодой
    UNIQUE(group_id, deck_id),

    FOREIGN KEY (group_id) REFERENCES groups(group_id),
    FOREIGN KEY (deck_id) REFERENCES deck_config(deck_id)
);

CREATE TABLE IF NOT EXISTS group_members (
    group_id SERIAL,
    user_id SERIAL, -- юзер является членом группы group_id
    time_joined TIMESTAMP,
    UNIQUE(group_id, user_id),

    FOREIGN KEY (group_id) REFERENCES groups(group_id),
    FOREIGN KEY (user_id) REFERENCES user_credentials(user_id)
);

CREATE TABLE IF NOT EXISTS group_invites (
    group_id SERIAL,
    user_id SERIAL, -- юзер получил инвайт группы, но ещё его не принял
    time_sent TIMESTAMP,
    UNIQUE(group_id, user_id),

    FOREIGN KEY (group_id) REFERENCES groups(group_id),
    FOREIGN KEY (user_id) REFERENCES user_credentials(user_id)
);

--
-- CREATE TABLE IF NOT EXISTS links (
--     deck_id SERIAL PRIMARY KEY, -- дека откопирована от copied_from
--     copied_from SERIAL,
--     updatable BOOLEAN, -- можно ли обновлять колоду
--
--     FOREIGN KEY (deck_id) REFERENCES deck_config(deck_id),
--     FOREIGN KEY (copied_from) REFERENCES deck_config(deck_id),
--     CHECK ( deck_id != links.copied_from )
-- );

CREATE TABLE IF NOT EXISTS copied_by (
    copier_id INT, -- кто скопировал деку
    deck_id INT, -- какая дека скопирована
    time_copied TIMESTAMP,

    FOREIGN KEY (copier_id) REFERENCES user_credentials(user_id),
    FOREIGN KEY (deck_id) REFERENCES deck_config(deck_id)
);

CREATE TABLE IF NOT EXISTS public_shared (
    deck_id SERIAL PRIMARY KEY,
    time_shared TIMESTAMP,

    FOREIGN KEY (deck_id) REFERENCES deck_config(deck_id)
);

CREATE TABLE IF NOT EXISTS followers (
    user_id SERIAL,
    follower_id SERIAL, -- фолловер отслеживает действия юзера в фиде
    time_followed TIMESTAMP,
    UNIQUE(user_id, follower_id),

    FOREIGN KEY (user_id) REFERENCES user_credentials(user_id),
    FOREIGN KEY (follower_id) REFERENCES user_credentials(user_id)
);

-- Сценарии колод
--
-- 1. Создание
--      При создании колоды человеком user_id в транзакции создаётся запись в deck_config
--      (автоинкремент deck_id, user_id),
--      а карточки помещаются во flashcards, для каждой карточки новая строка, причём поле
--      deck_id у всех одинаковое - id новой колоды.
--
-- 2. Шаринг подписавшимся [на меня] или участникам в группе
--      Если юзер делится своей колодой, она добавляется в одну из таблиц.
--          group_shared    |deck_id, timestamp|    (только пользователи некоторой группы, в которой
--          ВЫ ЯВЛЯЕТЕСЬ СОЗДАТЕЛЕМ могут видеть ваши колоды)
--          public_shared   |deck_id, timestamp|    (все пользователи видят в фид вашу колоду и могут добавить её к себе)
--
-- 3. Копирование колоды из фида
--  При нажатии на кнопку "копировать" (например, колоду с id=2), сначала проверяется,
--  нет ли у юзера колоды, которая уже откопирована от колоды с id=2. Используется таблица links.
--  Для юзера в deck_config создается новая колода, получаем deck_id.
--  Затем в эту таблицу заносится строка (deck_id, copied_from=2, updatable=true). Таблица
--  flashcards заполняется новыми картами, эквивалентные из старой колоды, но deck_id будет новый.
--  Таким образом, мы сделали deep copy колоды.
--
-- 4. Копирование колоды из группы
--  Когда создатель шарит колоду группе, люди автоматически её заимствуют себе, но без deep copy
--  всей колоды.
--  По аналогии выше создаётся запись в deck_config, но в deck_config заносится строка
--  (deck_id, copied_from=2, updatable=false), а таблица flashcards не трогается.
--
-- 5. Изменение колоды
--  Изменять можно только свои колоды или скопированные из фида.
--
-- 6. Удаление колоды.
--  Удалить можно только свои колоды или скопированные из фида.
--  Колоды групповые удалять нельзя (если не создатель).
--  Если человек, удаляющий колоду, создатель... ЗАВЕРШИТЬ.

