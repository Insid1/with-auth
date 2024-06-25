CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "token"
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    token      varchar(255) NOT NULL,
    user_id    varchar(255) NOT NULL,
    created_at timestamp        DEFAULT now(),
    updated_at timestamp        DEFAULT now()
);
