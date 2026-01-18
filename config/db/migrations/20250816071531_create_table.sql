-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    role VARCHAR(100) DEFAULT 'user'::VARCHAR NOT NULL,
    email_verfied_at timestamp without time zone
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
