CREATE OR REPLACE FUNCTION update_function()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    NEW.version = OLD.version + 1;
    return NEW;
END;
$$ LANGUAGE plpgsql;

-- Function to handle delete operation (soft delete or hard delete) and return affected rows
CREATE OR REPLACE FUNCTION delete_record(
    _table_name TEXT,
    _id UUID,
    _soft_delete BOOLEAN DEFAULT TRUE
)
RETURNS TABLE(affected_rows INT) AS
$$
BEGIN
    IF _soft_delete THEN
        -- Soft delete
        EXECUTE format('UPDATE %I SET deleted_at = NOW() WHERE id = $1', _table_name) USING _id;
    ELSE
        -- Hard delete
        EXECUTE format('DELETE FROM %I WHERE id = $1', _table_name) USING _id;
    END IF;
    
    -- Get the number of affected rows
    GET DIAGNOSTICS affected_rows = ROW_COUNT;
    
    -- Return affected_rows
    RETURN QUERY SELECT affected_rows;
END;
$$
LANGUAGE plpgsql;
