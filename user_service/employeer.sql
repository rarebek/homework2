CREATE TABLE IF NOT EXISTS companies(
    id uuid PRIMARY KEY NOT NULL,
    company_name text NOT NULL,
    description text NOT NULL,
    email text NOT NULL,
    password text NOT NULL,
    address text,
    profile_picture text,
    website text,
    industry text NOT NULL,
    employee_count int NOT NULL,
    phone_number text,
    refresh_token text,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
)