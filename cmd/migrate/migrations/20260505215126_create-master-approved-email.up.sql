CREATE TABLE IF NOT EXISTS user_auth_approved_email (
    email TEXT PRIMARY KEY,
    is_admin BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

INSERT INTO user_auth_approved_email (email, is_admin) VALUES
('SYSRONNIE@GMAIL.COM', true);