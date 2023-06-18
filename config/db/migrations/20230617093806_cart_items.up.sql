CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE cart_items (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    item_id uuid NOT NULL,
    session_id uuid NOT NULL,
    quantity int DEFAULT 0 NOT NULL, 
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
ALTER TABLE ONLY cart_items
    ADD CONSTRAINT cart_items_pkey PRIMARY KEY (id);