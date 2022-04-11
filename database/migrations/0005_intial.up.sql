ALTER TABLE task
    ADD uniqueId UUID PRIMARY KEY DEFAULT gen_random_uuid();