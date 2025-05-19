-- Create ENUM types
CREATE TYPE mood_type AS ENUM (
    'Happy',
    'Fussy',
    'Tired',
    'Content',
    'Gassy',
    'Crying'
);

CREATE TYPE feed_type AS ENUM ('breast_milk', 'formula', 'solids');

-- Changed to include breast_milk
CREATE TYPE diaper_smell_type AS ENUM ('None', 'Mild', 'Strong', 'Foul');

CREATE TYPE diaper_softness_type AS ENUM ('Soft', 'Medium', 'Firm');

CREATE TYPE sleep_quality_type AS ENUM ('Good', 'Restless', 'Interrupted', 'Awake');

-- Create the events table
CREATE TABLE events (
    event_id SERIAL PRIMARY KEY,
    event_type VARCHAR(50) NOT NULL,
    event_time TIMESTAMP
    WITH
        TIME ZONE NOT NULL,
        notes TEXT,
        mood mood_type,
        created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT NOW (),
        updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT NOW ()
);

CREATE INDEX idx_events_event_type ON events (event_type);

CREATE INDEX idx_events_event_time ON events (event_time);

-- Create the feeding_events table
CREATE TABLE feeding_events (
    feeding_event_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES events (event_id) ON DELETE CASCADE,
    amount DECIMAL,
    feed_type feed_type, --Using enum type
    spitup BOOLEAN,
    start_time TIMESTAMP
    WITH
        TIME ZONE,
        end_time TIMESTAMP
    WITH
        TIME ZONE,
        notes TEXT
);

-- Create an index on event_id to improve query performance.
CREATE INDEX idx_feeding_events_event_id ON feeding_events (event_id);

-- Create the diaper_events table
CREATE TABLE diaper_events (
    diaper_event_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES events (event_id) ON DELETE CASCADE,
    poop BOOLEAN,
    smell diaper_smell_type, -- Using enum Type
    size VARCHAR(50),
    softness diaper_softness_type, -- Uisng enum Type
    notes TEXT
);

-- Create an index on event_id to improve query performance.
CREATE INDEX idx_diaper_events_event_id ON diaper_events (event_id);

-- Create the sleep_events table
CREATE TABLE sleep_events (
    sleep_event_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES events (event_id) ON DELETE CASCADE,
    start_attempt_time TIMESTAMP
    WITH
        TIME ZONE,
        actual_sleep_time TIMESTAMP
    WITH
        TIME ZONE,
        end_time TIMESTAMP
    WITH
        TIME ZONE,
        quality sleep_quality_type, -- Using enum type
        environment VARCHAR(100),
        notes TEXT
);

-- Create an index on event_id to improve query performance.
CREATE INDEX idx_sleep_events_event_id ON sleep_events (event_id);

-- Create the medicines table
CREATE TABLE medicines (
    medicine_id SERIAL PRIMARY KEY,
    medicine_name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT
);

-- Create the medicine_events table
CREATE TABLE medicine_events (
    medicine_event_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES events (event_id) ON DELETE CASCADE,
    medicine_id INTEGER REFERENCES medicines (medicine_id), --Foreign key to medicine table
    dosage VARCHAR(50),
    unit VARCHAR(50),
    route VARCHAR(50),
    notes TEXT
);

-- Create an index on event_id to improve query performance.
CREATE INDEX idx_medicine_events_event_id ON medicine_events (event_id);

-- Create the vaccination_events table
CREATE TABLE vaccination_events (
    vaccination_event_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES events (event_id) ON DELETE CASCADE,
    vaccine_name VARCHAR(100) NOT NULL,
    administered_by VARCHAR(100),
    location VARCHAR(100),
    notes TEXT
);

-- Create an index on event_id to improve query performance.
CREATE INDEX idx_vaccination_events_event_id ON vaccination_events (event_id);

-- Create the doctor_appointment_events table
CREATE TABLE doctor_appointment_events (
    doctor_appointment_event_id SERIAL PRIMARY KEY,
    event_id INTEGER NOT NULL REFERENCES events (event_id) ON DELETE CASCADE,
    doctor_name VARCHAR(100),
    reason VARCHAR(200),
    notes TEXT
);

-- Create an index on event_id to improve query performance.
CREATE INDEX idx_doctor_appointment_events_event_id ON doctor_appointment_events (event_id);
