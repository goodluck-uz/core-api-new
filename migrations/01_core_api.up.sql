-- Step 1: Create ENUM types
CREATE TYPE user_role AS ENUM ('admin', 'user', 'supervisor');


-- Countries table
CREATE TABLE countries (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);

-- Categories table
CREATE TABLE categories (
    id BIGSERIAL PRIMARY KEY,
    status BOOLEAN NOT NULL
);

-- Country translation table
CREATE TABLE country_translation (
    id BIGSERIAL PRIMARY KEY,
    country_id BIGINT NOT NULL REFERENCES countries(id),
    name VARCHAR(255) NOT NULL,
    locale VARCHAR(10) NOT NULL,
    UNIQUE(country_id, locale),  -- Ensure combination of country_id and locale is unique
    UNIQUE(locale)  -- Ensure locale is unique
);

-- Users table
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    about TEXT NOT NULL,
    avatar VARCHAR(255) NOT NULL,
    banner VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    country_id BIGINT NOT NULL REFERENCES countries(id),
    role user_role NOT NULL,
    category_id BIGINT NOT NULL REFERENCES categories(id)
);

-- Languages table
CREATE TABLE languages (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    locale VARCHAR(10) NOT NULL UNIQUE,
    CONSTRAINT languages_locale_fk FOREIGN KEY (locale) REFERENCES country_translation(locale)
);

-- Payments table
CREATE TABLE payments (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL
);

-- Categories translation table
CREATE TABLE categories_translation (
    id BIGSERIAL PRIMARY KEY,
    category_id BIGINT NOT NULL REFERENCES categories(id),
    name VARCHAR(255) NOT NULL,
    locale VARCHAR(10) NOT NULL,
    UNIQUE(category_id, locale) -- Ensure uniqueness for each category and locale combination
);

-- Socials variants table
CREATE TABLE socials_variants (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Socials table
CREATE TABLE socials (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    link VARCHAR(255) NOT NULL,
    user_id BIGINT NOT NULL REFERENCES users(id)
);