-- Create default tenant
INSERT INTO public.tenants (
    id,
    name,
    created_at,
    updated_at
) VALUES (
    '00000000-0000-0000-0000-000000000000',
    'Default Tenant',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (id) DO NOTHING;

-- Create default user
INSERT INTO public.users (
    id,
    email,
    tenant_id,
    created_at,
    updated_at,
    created_by,
    updated_by
) VALUES (
    '00000000-0000-0000-0000-000000000000',
    'admin@flexprice.dev',
    '00000000-0000-0000-0000-000000000000',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    '00000000-0000-0000-0000-000000000000',
    '00000000-0000-0000-0000-000000000000'
) ON CONFLICT (id) DO NOTHING;

-- Create two default ENVironments
INSERT INTO public.environments (
    id,
    name,
    type,
    tenant_id,
    status,
    created_by,
    updated_by,
    created_at,
    updated_at
) VALUES (
    '00000000-0000-0000-0000-000000000000',
    'Sandbox',
    'development',
    '00000000-0000-0000-0000-000000000000',
    'published',
    '00000000-0000-0000-0000-000000000000',
    '00000000-0000-0000-0000-000000000000',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (id) DO NOTHING;

INSERT INTO public.environments (
    id,
    name,
    type,
    tenant_id,
    status,
    created_by,
    updated_by,
    created_at,
    updated_at
) VALUES (
    '00000000-0000-0000-0000-000000000001',
    'Production',
    'production',
    '00000000-0000-0000-0000-000000000000',
    'published',
    '00000000-0000-0000-0000-000000000000',
    '00000000-0000-0000-0000-000000000000',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (id) DO NOTHING;
