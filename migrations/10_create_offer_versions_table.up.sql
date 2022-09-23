CREATE TABLE IF NOT EXISTS "offer_versions"(
    id uuid NOT NULL PRIMARY KEY DEFAULT uuid_generate_v4(),
    offer_code char(5) NOT NULL,
    activated_at timestamptz default current_timestamp,
    deactivated_at timestamptz default current_timestamp,
    created_at timestamptz default current_timestamp NOT NULL,
    FOREIGN KEY (offer_code) REFERENCES offers (code) ON DELETE CASCADE
)