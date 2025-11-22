CREATE TABLE IF NOT EXISTS user_auth_sessions(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    access_version INT NOT NULL DEFAULT 1,
    refresh_version INT NOT NULL DEFAULT 1,
    user_agent TEXT,
    id_address VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ
);
