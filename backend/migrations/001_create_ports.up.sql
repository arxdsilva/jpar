CREATE TABLE IF NOT EXISTS "ports" (
    "id" TEXT,
    "name" TEXT,
    "city" TEXT,
    "country" TEXT,
    "alias" jsonb,
    "regions" jsonb,
    "coordinates" jsonb,
    "province" TEXT,
    "timezone" TEXT,
    "unlocs" jsonb,
    "code" INT,
	created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	last_update timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	CONSTRAINT ports_pkey PRIMARY KEY (id)
);

CREATE OR REPLACE FUNCTION update_datetime()	
RETURNS TRIGGER AS $$
BEGIN
    NEW.last_update = now();
    RETURN NEW;	
END;
$$ language 'plpgsql';

CREATE TRIGGER update_ports
BEFORE UPDATE ON ports 
FOR EACH ROW EXECUTE PROCEDURE update_datetime();