CREATE TABLE IF NOT EXISTS freelancers(
    id uuid NOT NULL PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    username text unique NOT NULL,
    biography text,
    profile_picture text,
    email text NOT NULL,
    password text NOT NULL,
    phone_number text,
    address text,
    resume text,
    refresh_token text,
    created_at timestamp default current_timestamp NOT NULL,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS skills(
    skill_id uuid NOT NULL PRIMARY KEY,
    freelancer_id uuid NOT NULL,
    skill_name text NOT NULL,
    FOREIGN KEY(freelancer_id) REFERENCES freelancers(id)
);