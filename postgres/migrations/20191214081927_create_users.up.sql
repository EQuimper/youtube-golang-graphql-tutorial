CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(55) UNIQUE                     NOT NULL,
    email      VARCHAR(255) UNIQUE                    NOT NULL,
    first_name VARCHAR(255)                           NOT NULL,
    last_name  VARCHAR(255)                           NOT NULL,
    password   TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);