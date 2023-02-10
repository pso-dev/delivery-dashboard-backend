CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS clearance (
    clearance_id BIGSERIAL PRIMARY KEY,
    description VARCHAR(256) NOT NULL UNIQUE ON DELETE RESTRICT
);

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

CREATE TABLE IF NOT EXISTS workgroup (
    workgroup_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL UNIQUE ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS location (
    location_id BIGSERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL UNIQUE ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS resource (
    employee_id INTEGER PRIMARY KEY,
    name VARCHAR(256) NOT NULL ON DELETE RESTRICT,
    email VARCHAR(256) NOT NULL,
    title_id INTEGER NOT NULL,
    manager_id INTEGER NOT NULL,
    location_id INTEGER NOT NULL,
    workgroup_id INTEGER NOT NULL,
    clearance_id INTEGER,
    active BOOLEAN NOT NULL DEFAULT 't',
    CONSTRAINT fk_title FOREIGN KEY(title_id) REFERENCES jobtitle(title_id),
    CONSTRAINT fk_manager FOREIGN KEY(manager_id) REFERENCES resource(employee_id),
    CONSTRAINT fk_location FOREIGN KEY(location_id) REFERENCES location(location_id),
    CONSTRAINT fk_workgroup FOREIGN KEY(workgroup_id) REFERENCES workgroup(workgroup_id),
    CONSTRAINT fk_clearance FOREIGN KEY(clearance_id) REFERENCES clearance(clearance_id),
);

CREATE TABLE IF NOT EXISTS resource_specialty (
    employee_id INTEGER NOT NULL,
    specialty_id INTEGER NOT NULL,
    PRIMARY KEY(resource_id, specialty_id)
);

CREATE TABLE IF NOT EXISTS project (
    project_id BIGSERIAL PRIMARY KEY,
    opportunity_id VARCHAR(256),
    changepoint_id VARCHAR(256),
    revenue_type INTEGER NOT NULL,
    name VARCHAR(256) NOT NULL,
    customer VARCHAR(256) NOT NULL,
    end_customer VARCHAR(256),
    project_manager_id INTEGER,
    status_id INTEGER,
);