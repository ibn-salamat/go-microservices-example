CREATE TABLE
    users (
        id serial PRIMARY KEY,
        email varchar NOT NULL UNIQUE,
        username varchar NOT NULL
    );