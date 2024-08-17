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





