CREATE TABLE sessions(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    uid uuid NOT NULL ,
    login_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    log_out  TIMESTAMP WITH TIME ZONE DEFAULT NULL
);
