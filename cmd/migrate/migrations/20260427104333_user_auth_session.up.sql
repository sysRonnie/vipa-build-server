

-- Create user_auth_session table
CREATE TABLE IF NOT EXISTS user_auth_session (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES user_auth(uuid) ON DELETE CASCADE,
    session_id UUID NOT NULL DEFAULT gen_random_uuid() UNIQUE,
    refresh_token_hash TEXT NOT NULL UNIQUE,
    user_agent TEXT,
    ip_address INET,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    last_used_at TIMESTAMPTZ DEFAULT now(),
    expires_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ,
    revoked_reason TEXT
);
