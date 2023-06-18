CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE access_tokens (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    token VARCHAR(40) DEFAULT false NOT NULL,
    active boolean DEFAULT true NOT NULL, 
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
ALTER TABLE ONLY access_tokens
    ADD CONSTRAINT access_tokens_pkey PRIMARY KEY (id);