-- Example migration: aligns with ExampleItem in models.go.
-- Add new numbered pairs (.up.sql / .down.sql) for each schema change—do not rewrite old files in production databases.

CREATE TABLE IF NOT EXISTS example_items (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX IF NOT EXISTS idx_example_items_deleted_at ON example_items (deleted_at);
