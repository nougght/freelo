-- +goose Up

CREATE TABLE users
(
    id         UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    email      VARCHAR(256) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ  NOT NULL DEFAULT now()
);

CREATE TABLE user_profiles
(
    user_id             UUID PRIMARY KEY REFERENCES users (id) ON DELETE CASCADE,
    username VARCHAR(32) UNIQUE,
    display_name        VARCHAR(256) NOT NULL,
    client_rating_avg   NUMERIC,
    client_rating_count NUMERIC,
);

CREATE TABLE freelance_profiles
(
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    bio_short VARCHAR(100),
    bio TEXT,
    rating_avg NUMERIC NOT NULL DEFAULT 0,
    rating_count NUMERIC NOT NULL DEFAULT 0,
    orders_completed NUMERIC NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
)

-- +goose Down

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS user_profiles;
DROP TABLE IF EXISTS freelanse_profiles;