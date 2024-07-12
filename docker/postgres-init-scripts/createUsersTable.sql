CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    status INT NOT NULL
);



CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    name_event VARCHAR(255) NOT NULL,
    shape VARCHAR(255) NOT NULL,
    place VARCHAR(255) NOT NULL,
    begin_time TIMESTAMP NOT NULL,
    duration VARCHAR(255) NOT NULL
);


CREATE TABLE contacts (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE
);


INSERT INTO contacts (id, username, email)
SELECT id, username, email
FROM users;


-- Триггер для вставки в таблицу users
CREATE OR REPLACE FUNCTION sync_contacts_insert() RETURNS TRIGGER AS $$
BEGIN
    INSERT INTO contacts (id, username, email)
    VALUES (NEW.id, NEW.username, NEW.email);
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_users_insert
AFTER INSERT ON users
FOR EACH ROW
EXECUTE FUNCTION sync_contacts_insert();

-- Триггер для обновления таблицы users
CREATE OR REPLACE FUNCTION sync_contacts_update() RETURNS TRIGGER AS $$
BEGIN
    UPDATE contacts
    SET username = NEW.username, email = NEW.email
    WHERE id = NEW.id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_users_update
AFTER UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION sync_contacts_update();

-- Триггер для удаления из таблицы users
CREATE OR REPLACE FUNCTION sync_contacts_delete() RETURNS TRIGGER AS $$
BEGIN
    DELETE FROM contacts
    WHERE id = OLD.id;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_users_delete
AFTER DELETE ON users
FOR EACH ROW
EXECUTE FUNCTION sync_contacts_delete();