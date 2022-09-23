CREATE TABLE IF NOT EXISTS "offer_version_translations"(
    offer_version_id uuid NOT NULL,
    locale varchar NOT NULL PRIMARY KEY,
    name varchar NOT NULL,
    content text NOT NULL,
    FOREIGN KEY (offer_version_id) REFERENCES offer_versions (id) ON DELETE CASCADE
)