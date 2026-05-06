CREATE TABLE IF NOT EXISTS master_client_list (
    id SERIAL PRIMARY KEY,
    client_name VARCHAR(255) NOT NULL UNIQUE,
    client_email VARCHAR(255),
    client_phone VARCHAR(20) NOT NULL,
    client_address TEXT,
    comment TEXT,
    flag_is_deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);


CREATE TRIGGER update_master_clients_updated_at
BEFORE UPDATE ON master_client_list
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();