CREATE TABLE IF NOT EXISTS master_event_category (
    id SERIAL PRIMARY KEY,
    event_category_parent TEXT NOT NULL,
    event_category_child TEXT NOT NULL DEFAULT '',
    flag_is_deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS uq_event_category_name
ON master_event_category (
    LOWER(event_category_parent),
    LOWER(event_category_child)
);

INSERT INTO master_event_category (
    event_category_parent
)
VALUES
('Pre-Construction'),
('Site Work'),
('Foundation'),
('Framing'),
('Roofing'),
('Exterior'),
('Plumbing'),
('Electrical'),
('HVAC'),
('Insulation'),
('Drywall'),
('Painting'),
('Flooring'),
('Cabinetry'),
('Finish Work'),
('Landscaping'),
('Inspection'),
('Punch List'),
('Final Walkthrough'),
('General')
ON CONFLICT DO NOTHING;