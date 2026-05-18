CREATE TABLE IF NOT EXISTS MASTER_EVENT_CATEGORY (
    ID SERIAL PRIMARY KEY,
    EVENT_CATEGORY_PARENT TEXT NOT NULL,
    EVENT_CATEGORY_CHILD TEXT,
    FLAG_IS_DELETED BOOLEAN DEFAULT FALSE,
    CREATED_AT TIMESTAMPTZ DEFAULT NOW(),
    UPDATED_AT TIMESTAMPTZ DEFAULT NOW(),

    UNIQUE (EVENT_CATEGORY_PARENT, EVENT_CATEGORY_CHILD)
);

INSERT INTO MASTER_EVENT_CATEGORY (EVENT_CATEGORY_PARENT)
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