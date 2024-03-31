-- +migrate Up
-- Создание таблицы пользователей (profile)
CREATE TABLE IF NOT EXISTS profile (
    id SERIAL PRIMARY KEY,
    login TEXT UNIQUE CHECK (LENGTH(login) <= 50),
    password TEXT CHECK (LENGTH(password) <= 200),
    firstname TEXT CHECK (LENGTH(firstname) <= 50),
    surname TEXT CHECK (LENGTH(surname) <= 50),
    patronymic TEXT CHECK (LENGTH(patronymic) <= 50),
    gender TEXT CHECK (LENGTH(gender) <= 10),
    birthday DATE,
    registration_date DATE,
    avatar_id TEXT CHECK (LENGTH(avatar_id) <= 200),
    phone_number TEXT CHECK (LENGTH(phone_number) <= 20),
    description TEXT CHECK (LENGTH(description) <= 300)
);

-- Создание таблицы сессий (session)
CREATE TABLE IF NOT EXISTS session (
    id TEXT PRIMARY KEY CHECK (char_length(id) <= 50),
    profile_id INTEGER REFERENCES profile(id) ON DELETE CASCADE,
    creation_date TIMESTAMP,
    device TEXT CHECK (LENGTH(device) <= 100),
    life_time INTEGER,
    csrf_token TEXT CHECK (char_length(csrf_token) <= 50)
);

-- Вставка начальных данных в таблицу users
INSERT INTO profile
    (id, login, password, firstname, surname, patronymic, gender, birthday, registration_date, avatar_id, phone_number, description)
VALUES
    (1, 'sergey@mailhub.su', '$2a$10$4PcooWbEMRjvdk2cMFumO.ajWaAclawIljtlfu2.2f5/fV8LkgEZe', 'Sergey', 'Fedasov', 'Aleksandrovich', 'Male', '2003-08-20', NOW(), '', '+77777777777', 'Description'),
    (2, 'ivan@mailhub.su', '$2a$10$4PcooWbEMRjvdk2cMFumO.ajWaAclawIljtlfu2.2f5/fV8LkgEZe', 'Ivan', 'Karpov', 'Aleksandrovich', 'Male', '2003-10-17', NOW(), '', '+79697045539', 'Description'),
    (3, 'max@mailhub.su', '$2a$10$4PcooWbEMRjvdk2cMFumO.ajWaAclawIljtlfu2.2f5/fV8LkgEZe', 'Maxim', 'Frelich', 'Aleksandrovich', 'Male', '2003-08-20', NOW(), '', '+79099099090', 'Description')
ON CONFLICT (login) DO NOTHING;

-- Создание таблицы писем (email)
CREATE TABLE IF NOT EXISTS email (
    id SERIAL PRIMARY KEY,
    topic TEXT,
    text TEXT,
    date_of_dispatch DATE,      /*TIMESTAMPTZ NOT NULL DEFAULT '2022-08-10 10:10:00',*/
    photoid TEXT CHECK (LENGTH(photoid) <= 200),
    sender_email TEXT CHECK (LENGTH(sender_email) <= 50),
    recipient_email TEXT CHECK (LENGTH(recipient_email) <= 50),
    read_status BOOLEAN NOT NULL,
    deleted_status BOOLEAN NOT NULL,
    draft_status BOOLEAN NOT NULL,
    reply_to_email_id INTEGER REFERENCES email(id) ON DELETE NO ACTION DEFAULT NULL,
    flag BOOLEAN NOT NULL
);

-- Вставка начальных данных в таблицу email
INSERT INTO email
    (topic, text, date_of_dispatch, photoid, sender_email, recipient_email, read_status, deleted_status, draft_status, reply_to_email_id, flag)
VALUES
    ('Topic1 Enough pretended estimating.', 'Laughing say assurance indulgence mean unlocked stairs denote above prudent get use latter margaret. Unreserved another abode blushes old steepest lady disposing enjoyment immediate prevailed charm. Looked ladies civil sigh. Because cold offended quiet bred the. Hastened outlived supported.', '2022-08-10 10:10:00', '', 'sergey@mailhub.su', 'ivan@mailhub.su', False, False, False, Null, False),
    ('Topic2 Enough pretended estimating.', 'Laughing say assurance indulgence mean unlocked stairs denote above prudent get use latter margaret. Unreserved another abode blushes old steepest lady disposing enjoyment immediate prevailed charm. Looked ladies civil sigh. Because cold offended quiet bred the. Hastened outlived supported.', '2022-08-10 10:10:00', '', 'sergey@mailhub.su', 'max@mailhub.su', False, False, False, Null, False),
    ('Topic3 Enough pretended estimating.', 'Laughing say assurance indulgence mean unlocked stairs denote above prudent get use latter margaret. Unreserved another abode blushes old steepest lady disposing enjoyment immediate prevailed charm. Looked ladies civil sigh. Because cold offended quiet bred the. Hastened outlived supported.', '2022-08-10 10:10:00', '', 'ivan@mailhub.su', 'max@mailhub.su', False, False, False, Null, False),
    ('Topic4 Enough pretended estimating.', 'Laughing say assurance indulgence mean unlocked stairs denote above prudent get use latter margaret. Unreserved another abode blushes old steepest lady disposing enjoyment immediate prevailed charm. Looked ladies civil sigh. Because cold offended quiet bred the. Hastened outlived supported.', '2022-08-10 10:10:00', '', 'max@mailhub.su', 'sergey@mailhub.su', False, False, False, Null, False),
    ('Topic5 Enough pretended estimating.', 'Laughing say assurance indulgence mean unlocked stairs denote above prudent get use latter margaret. Unreserved another abode blushes old steepest lady disposing enjoyment immediate prevailed charm. Looked ladies civil sigh. Because cold offended quiet bred the. Hastened outlived supported.', '2022-08-10 10:10:00', '', 'max@mailhub.su', 'ivan@mailhub.su', False, False, False, Null, False),
    ('Topic6 Enough pretended estimating.', 'Laughing say assurance indulgence mean unlocked stairs denote above prudent get use latter margaret. Unreserved another abode blushes old steepest lady disposing enjoyment immediate prevailed charm. Looked ladies civil sigh. Because cold offended quiet bred the. Hastened outlived supported.', '2022-08-10 10:10:00', '', 'sergey@mailhub.su', 'ivan@mailhub.su', False, False, False, Null, False)
;