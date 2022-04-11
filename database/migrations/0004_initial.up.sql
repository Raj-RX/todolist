CREATE TABLE task(
        taskID UUID REFERENCES users(id),
        taskName TEXT NOT NULL,
        status boolean NOT NULL DEFAULT TRUE,
        completed_upTo TIMESTAMP WITH TIME ZONE,
        archived_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
