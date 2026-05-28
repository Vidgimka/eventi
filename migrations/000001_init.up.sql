CREATE SCHEMA eventi;

CREATE TABLE eventi.users (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    full_name VARCHAR(100) NOT NULL CHECK (char_length(full_name) BETWEEN 3 AND 100),
    phone_number VARCHAR(100) CHECK (
        phone_number ~ '^+[0-9]+$'
        AND
        char_length(phone_number) BETWEEN 10 AND 15
    )
);

CREATE TABLE eventi.tasks (
    id SERIAL PRIMARY KEY,
    version BIGINT NOT NULL DEFAULT 1,
    title VARCHAR(100) NOT NULL CHECK (char_length(title) BETWEEN 1 AND 100),
    description VARCHAR(100) NOT NULL CHECK (char_length(description) BETWEEN 1 AND 100),
    complited BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    complited_at TIMESTAMPTZ,

    CHECK (
        (complited=FALSE AND complited_at IS NULL)
        OR
        (complited=TRUE AND complited_at IS NOT NULL AND  complited_at >= created_at)
    ),

    author_user_id INT NOT NULL,
    FOREIGN KEY (author_user_id) REFERENCES eventi.users(id)
);
