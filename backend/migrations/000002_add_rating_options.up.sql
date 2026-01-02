CREATE TABLE IF NOT EXISTS rating_options (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    label STRING(255) NOT NULL,
    value STRING(255) NOT NULL,
    is_active BOOL DEFAULT true,
    created_at TIMESTAMPTZ DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT now(),
    CONSTRAINT rating_options_label_unique UNIQUE (label),
    CONSTRAINT rating_options_value_unique UNIQUE (value)
);

CREATE INDEX IF NOT EXISTS idx_rating_options_label ON rating_options(label);
CREATE INDEX IF NOT EXISTS idx_rating_options_active ON rating_options(is_active);