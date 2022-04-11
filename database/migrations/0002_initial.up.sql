CREATE TABLE cart
(
    id          UUID PRIMARY KEY         DEFAULT gen_random_uuid(),
    cart_name        TEXT NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    archived_at TIMESTAMP WITH TIME ZONE
);
