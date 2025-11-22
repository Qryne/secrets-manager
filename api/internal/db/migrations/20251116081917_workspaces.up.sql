CREATE TABLE IF NOT EXISTS workspaces (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS workspace_roles(
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
name VARCHAR(255) NOT NULL,
slug VARCHAR(255) UNIQUE NOT NULL,
permissions JSONB NOT NULL DEFAULT '{}'::jsonb,
workspace_id  UUID NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
updated_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS workspace_members(
id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
user_id  UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
workspace_id  UUID NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
role_id  UUID NOT NULL REFERENCES workspace_roles(id) ON DELETE CASCADE,
created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
updated_at TIMESTAMPTZ,

CONSTRAINT uq_wks_member UNIQUE (user_id, workspace_id)
);
