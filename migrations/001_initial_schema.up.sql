
CREATE SCHEMA IF NOT EXISTS extensions;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA extensions;

CREATE TABLE IF NOT EXISTS patients (
    id UUID PRIMARY KEY DEFAULT extensions.uuid_generate_v4(),
    fname VARCHAR NOT NULL,
    lname VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TYPE assessment AS ENUM('vital signs');

CREATE TABLE IF NOT EXISTS assessments_eav (
    id UUID PRIMARY KEY DEFAULT extensions.uuid_generate_v4(),
    assessment assessment NOT NULL,
    display_nm VARCHAR NOT NULL,
    FOREIGN KEY (patient_id) REFERENCES patients(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TYPE attribute AS ENUM('temperature');

CREATE TABLE IF NOT EXISTS assessment_attributes (
    id UUID PRIMARY KEY DEFAULT extensions.uuid_generate_v4(),
    assessment_id UUID NOT NULL,
    display_nm VARCHAR NOT NULL,
    attribute attribute NOT NULL,
    dtype VARCHAR(50) NOT NULL,
    FOREIGN KEY (assessment_id) REFERENCES assessments(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS assessment_values (
    id UUID PRIMARY KEY DEFAULT extensions.uuid_generate_v4(),
    patient_id UUID NOT NULL,
    assessment_id UUID NOT NULL,
    assessment_attribute_id UUID NOT NULL, 
    input JSONB NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
