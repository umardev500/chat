CREATE TABLE IF NOT EXISTS user_rooms (
    user_id UUID NOT NULL,
    room_id UUID NOT NULL,
    PRIMARY KEY (user_id, room_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (room_id) REFERENCES chat_rooms(id) ON DELETE CASCADE,
    deleted_at TIMESTAMPTZ DEFAULT NULL
);
