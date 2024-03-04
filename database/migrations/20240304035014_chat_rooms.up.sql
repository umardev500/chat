CREATE TABLE IF NOT EXISTS chat_rooms (
    id UUID PRIMARY KEY,
    name VARCHAR(255) DEFAULT NULL,
    is_group BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    "version" INT NOT NULL DEFAULT 0
);

CREATE TRIGGER update_chat_rooms_trigger BEFORE
UPDATE ON chat_rooms FOR EACH ROW
    WHEN (
        NEW.name IS DISTINCT
        FROM OLD.name
            OR NEW.is_group IS DISTINCT
        FROM OLD.is_group
    ) EXECUTE FUNCTION update_function();