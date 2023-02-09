CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS certification (
    certification_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL UNIQUE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS specialty (
    specialty_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL UNIQUE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS jobtitle (
    title_id BIGSERIAL PRIMARY KEY,
    title VARCHAR(256) NOT NULL UNIQUE ON DELETE RESTRICT
);

CREATE TABLE IF EXISTS workgroup (
    workgroup_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL UNIQUE ON DELETE RESTRICT
);