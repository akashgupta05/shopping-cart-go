CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    username VARCHAR(32) NOT NULL,
    password_digest VARCHAR(64) DEFAULT false NOT NULL,
    role VARCHAR(16) NOT NULL, 
    active boolean DEFAULT true NOT NULL, 
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);