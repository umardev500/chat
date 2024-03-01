CREATE OR REPLACE FUNCTION update_function()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    NEW.version = OLD.version + 1;
    return NEW;
END;
$$ LANGUAGE plpgsql;