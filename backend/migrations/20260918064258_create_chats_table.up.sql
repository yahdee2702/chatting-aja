CREATE TYPE chat_type AS ENUM (
    'individual',
    'group'
);

CREATE TABLE chats (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    type chat_type NOT NULL,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    deleted_at TIMESTAMPTZ
);

CREATE TABLE chats_users (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    chat_id UUID NOT NULL REFERENCES chats(id) ON DELETE CASCADE,

    PRIMARY KEY (chat_id, user_id)
);