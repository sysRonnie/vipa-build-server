CREATE TABLE IF NOT EXISTS user_project_note (
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES user_auth(uuid),
    project_id INT REFERENCES master_project_list(id),
    note_body TEXT NOT NULL,
    note_photo_url TEXT,
    flag_is_deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);