CREATE TABLE IF NOT EXISTS user_details (
    id UUID PRIMARY KEY,
    fist_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) DEFAULT NULL,
    date_of_birth DATE DEFAULT NULL,
    bio TEXT DEFAULT NULL,
    photo TEXT DEFAULT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    "version" INT NOT NULL DEFAULT 1
);

CREATE TRIGGER update_user_details_trigger BEFORE
UPDATE ON user_details FOR EACH ROW
    WHEN (
        NEW.fist_name IS DISTINCT
        FROM OLD.fist_name
            OR NEW.last_name IS DISTINCT
        FROM OLD.last_name
            OR NEW.date_of_birth IS DISTINCT
        FROM OLD.date_of_birth
            OR NEW.photo IS DISTINCT
        FROM OLD.photo
    ) EXECUTE FUNCTION update_function();