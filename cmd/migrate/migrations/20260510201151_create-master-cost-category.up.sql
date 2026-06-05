CREATE TABLE IF NOT EXISTS MASTER_COST_CATEGORY (
    ID SERIAL PRIMARY KEY,
    COST_CATEGORY_PARENT TEXT NOT NULL,
    cost_category_child TEXT NOT NULL DEFAULT '',
    FLAG_IS_DELETED BOOLEAN DEFAULT FALSE,
    CREATED_AT TIMESTAMPTZ DEFAULT NOW(),
    UPDATED_AT TIMESTAMPTZ DEFAULT NOW(),

    UNIQUE (COST_CATEGORY_PARENT, COST_CATEGORY_CHILD)
);

INSERT INTO MASTER_COST_CATEGORY (
    COST_CATEGORY_PARENT,
    COST_CATEGORY_CHILD
)
VALUES
('Pre-Construction', 'Land Acquisition'),
('Pre-Construction', 'Permits'),
('Pre-Construction', 'Architectural Plans'),
('Pre-Construction', 'Engineering'),
('Pre-Construction', 'Surveying'),

('Site Work', 'Clearing'),
('Site Work', 'Grading'),
('Site Work', 'Excavation'),
('Site Work', 'Temporary Utilities'),

('Foundation', 'Footings'),
('Foundation', 'Slab'),
('Foundation', 'Basement'),
('Foundation', 'Waterproofing'),

('Framing', 'Lumber'),
('Framing', 'Labor'),
('Framing', 'Trusses'),
('Framing', 'Sheathing'),

('Exterior', 'Roofing'),
('Exterior', 'Siding'),
('Exterior', 'Windows'),
('Exterior', 'Exterior Doors'),
('Exterior', 'Gutters'),

('Mechanical', 'Plumbing'),
('Mechanical', 'Electrical'),
('Mechanical', 'HVAC'),

('Interior', 'Insulation'),
('Interior', 'Drywall'),
('Interior', 'Painting'),
('Interior', 'Trim'),
('Interior', 'Interior Doors'),

('Finishes', 'Flooring'),
('Finishes', 'Cabinets'),
('Finishes', 'Countertops'),
('Finishes', 'Tile'),
('Finishes', 'Fixtures'),
('Finishes', 'Appliances'),

('Exterior Living', 'Deck'),
('Exterior Living', 'Porch'),
('Exterior Living', 'Driveway'),
('Exterior Living', 'Landscaping'),

('General', 'Dumpster'),
('General', 'Temporary Toilet'),
('General', 'Cleaning'),
('General', 'Insurance'),
('General', 'Project Management'),

('Contingency', 'Unexpected Costs')

ON CONFLICT (
    COST_CATEGORY_PARENT,
    COST_CATEGORY_CHILD
) DO NOTHING;