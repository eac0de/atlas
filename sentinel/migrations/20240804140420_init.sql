-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    users (
        id UUID PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        created_at TIMESTAMP NOT NULL,
        is_super BOOLEAN NOT NULL
    );

CREATE INDEX email_users_idx ON users USING hash (email);

CREATE TABLE
    email_codes (
        id UUID PRIMARY KEY,
        email VARCHAR(255) NOT NULL,
        code SMALLINT NOT NULL,
        expires_at TIMESTAMP NOT NULL,
        number_of_attempts SMALLINT NOT NULL
    );

CREATE TABLE
    sessions (
        id UUID PRIMARY KEY,
        token VARCHAR(255) UNIQUE NOT NULL,
        user_id UUID NOT NULL,
        ip INET NOT NULL,
        location VARCHAR(255) NOT NULL,
        client_info VARCHAR(255) NOT NULL,
        last_login TIMESTAMP NOT NULL,
        CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE refresh_tokens;

DROP TABLE email_codes;

DROP TABLE users;

-- +goose StatementEnd