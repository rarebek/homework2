CREATE TABLE IF NOT EXISTS companies (
    id SERIAL PRIMARY KEY,
    company_name VARCHAR(255) NOT NULL,
    description TEXT,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    address TEXT,
    profile_picture TEXT,
    website TEXT,
    industry TEXT,
    employee_count INTEGER,
    phone_number VARCHAR(20),
    refresh_token TEXT
);
