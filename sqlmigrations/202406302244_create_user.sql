CREATE TABLE users (
    id UUID NOT NULL,
    email VARCHAR(254) NOT NULL,
    password VARCHAR(72) NOT NULL,
    is_admin BOOL NOT_NULL DEFAULT FALSE,
    details JSONB NOT_NULL,
    created_at INTERGER NOT_NULL DEFAULT EXTRACT(EPOCH FROM now())::int,
    updated_at INTERGER,
    CONSTRAINT users_id_pk PRIMARY_KEY (id),
    CONSTRAINT users_email_uk UNIQUE (email)

),

COMMENT ON TABLE users IS 'Storage the admins and customers for the e-commerce',