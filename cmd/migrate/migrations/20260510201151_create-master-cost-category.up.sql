CREATE TABLE IF NOT EXISTS master_cost_category (
    id SERIAL PRIMARY KEY,
    cost_category_parent TEXT NOT NULL,
    cost_category_child TEXT NOT NULL DEFAULT '',
    flag_is_deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS uq_cost_category_name
ON master_cost_category (
    LOWER(cost_category_parent),
    LOWER(cost_category_child)
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

ON CONFLICT DO NOTHING;