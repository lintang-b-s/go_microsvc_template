CREATE TYPE gender as enum (
    'Male',
    'Female'
);

CREATE TABLE users(
    id  UUID DEFAULT gen_random_uuid() PRIMARY KEY ,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    dob DATE NOT NULL,
    gender gender NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_time timestamp with time zone  DEFAULT NOW() NOT NULL,
    updated_time timestamp with time zone 
);

CREATE TABLE sessions (
    id uuid PRIMARY KEY NOT NULL,
    username varchar NOT NULL,
    refresh_token varchar NOT NULL,
    expires_at timestamptz NOT NULL,
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz NOT NULL DEFAULT (now()),
    deleted_at timestamptz
);