CREATE TABLE IF NOT EXISTS master_vendor_list (
    id SERIAL PRIMARY KEY,
    vendor_name TEXT NOT NULL,
    vendor_primary_contact TEXT,
    vendor_phone TEXT,
    vendor_email TEXT,
    vendor_address_street TEXT,
    vendor_address_unit TEXT,
    vendor_address_city TEXT,
    vendor_address_state TEXT,
    vendor_address_zip TEXT,
    comment TEXT,
    flag_is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS master_vendor_list_vendor_name_unique_idx
ON master_vendor_list (LOWER(vendor_name));

INSERT INTO master_vendor_list (vendor_name)
VALUES
('84 Lumber'),
('Allegheny Mineral Corp'),
('Fireplace and Patio'),
('Glacial Sand & Gravel'),
('Grove City Agway'),
('Interstate Pipe Supply'),
('Slippery Rock Borough'),
('Slippery Rock Law'),
('Sontagg Excavating'),
('SRMA'),
('Susi Builders Supply'),
('WestPenn Power'),
('Northern Surveyers'),
('Jorgensen Bros'),
('Ace Fix-it'),
('Armina Stone'),
('Central Heating and Plumbing'),
('Doors-N-More'),
('Lowes Hardware'),
('Home Depot'),
('Montgomery Block Works'),
('Nicklas Supply'),
('Pella Windows and Doors'),
('Sherwin-Williams'),
('Sheetz'),
('Slippery Rock Hardware'),
('Tru-Guard Waterproofing'),
('WPM Masonry')
ON CONFLICT DO NOTHING;