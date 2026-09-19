CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    sender_id UUID NOT NULL REFERENCES users(id),
    
    chat_id UUID REFERENCES chats(id),
    channel_id UUID REFERENCES channels(id),

    ciphertext BYTEA NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_messages_sender_id ON messages(sender_id, created_at DESC);
CREATE INDEX idx_messages_chat_id ON messages(chat_id, created_at DESC);
CREATE INDEX idx_messages_channel_id ON messages(channel_id, created_at DESC);