CREATE TABLE
    users (
        id integer PRIMARY KEY,
        email varchar NOT NULL UNIQUE,
        username varchar NOT NULL UNIQUE
    );