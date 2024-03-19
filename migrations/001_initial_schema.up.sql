
CREATE SCHEMA IF NOT EXISTS extensions;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA extensions;

CREATE TYPE person_t as ENUM ('patient', 'care team');

CREATE TABLE IF NOT EXISTS persons (
    id UUID PRIMARY KEY DEFAULT extensions.uuid_generate_v4(),
    person_type person_t NOT NULL,
    fname VARCHAR NOT NULL,
    lname VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS surveys_eav (
    id UUID PRIMARY KEY DEFAULT extensions.uuid_generate_v4(),
    display_nm VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS survey_attributes (
    id UUID PRIMARY KEY DEFAULT extensions.uuid_generate_v4(),
    survey_id UUID NOT NULL,
    display_nm VARCHAR NOT NULL,
    description VARCHAR,
    dtype VARCHAR(50) NOT NULL,
    order_execution INTEGER NOT NULL,
    FOREIGN KEY (survey_id) REFERENCES surveys_eav(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS survey_values (
    id UUID PRIMARY KEY DEFAULT extensions.uuid_generate_v4(),
    submitted_by UUID NOT NULL,
    survey_id UUID NOT NULL,
    survey_attribute_id UUID NOT NULL, 
    input JSONB NOT NULL,
    FOREIGN KEY (survey_id) REFERENCES surveys_eav(id),
    FOREIGN KEY (survey_attribute_id) REFERENCES survey_attributes(id),
    FOREIGN KEY (submitted_by) REFERENCES persons(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
