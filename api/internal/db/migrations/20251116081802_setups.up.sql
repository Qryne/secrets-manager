CREATE TABLE IF NOT EXISTS setups(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    is_setup_complete BOOLEAN NOT NULL DEFAULT FALSE,
    destroy_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS setup_owners(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id  UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    setup_id UUID NOT NULL REFERENCES setups(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ,

    CONSTRAINT uq_setup_owners UNIQUE (user_id, setup_id)
)
