CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE cart_sessions (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    total int DEFAULT 0,
    status VARCHAR(16) NOT NULL, 
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
ALTER TABLE ONLY cart_sessions
    ADD CONSTRAINT cart_sessions_pkey PRIMARY KEY (id);