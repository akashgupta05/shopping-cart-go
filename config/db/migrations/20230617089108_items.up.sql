CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE items (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    name VARCHAR(64) NOT NULL,
    quantity int DEFAULT 0 NOT NULL,
    price int DEFAULT 0, 
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
ALTER TABLE ONLY items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);