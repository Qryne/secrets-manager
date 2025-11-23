CREATE TABLE IF NOT EXISTS api_keys(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    prefix VARCHAR(10) NOT NULL,
    public_id VARCHAR(100) UNIQUE NOT NULL,
    scope TEXT[] NOT NULL,
    encryption_iv TEXT UNIQUE NOT NULL,
    encrypted_text TEXT UNIQUE NOT NULL,
    algorithm VARCHAR(100) NOT NULL,
    rotations INT UNIQUE NOT NULL DEFAULT 0,
    setup_id UUID NOT NULL REFERENCES setups(id) ON DELETE CASCADE,
    last_rotated_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ
)
