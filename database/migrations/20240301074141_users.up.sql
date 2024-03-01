CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    "version" INT NOT NULL DEFAULT 0
);
CREATE TRIGGER update_users_trigger BEFORE
UPDATE ON users FOR EACH ROW
    WHEN (
        NEW.username IS DISTINCT
        FROM OLD.username
            OR NEW.password IS DISTINCT
        FROM OLD.password
            OR NEW.email IS DISTINCT
        FROM OLD.email
    ) EXECUTE FUNCTION update_function();