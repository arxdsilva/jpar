CREATE TABLE IF NOT EXISTS "ports" (
    "id" text,
    "name" text,
    "city" text,
    "country" text,
    "alias" text[],
    "regions" text[],
    "coordinates" numeric[],
    "province" text,
    "timezone" text,
    "unlocs" text[],
    "code" text,
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