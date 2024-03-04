CREATE TABLE IF NOT EXISTS user_rooms (
    id UUID DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    room_id UUID NOT NULL,
    PRIMARY KEY (id, user_id, room_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (room_id) REFERENCES chat_rooms(id) ON DELETE CASCADE,
    deleted_at TIMESTAMPTZ DEFAULT NULL
);
