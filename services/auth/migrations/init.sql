CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "token"
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    token      varchar(255) NOT NULL,
    user_id    varchar(255) NOT NULL,
    created_at timestamp        DEFAULT now(),
    updated_at timestamp        DEFAULT now()
);


-- INSERT INTO "token" (id, token, user_id)
-- VALUES ('8fe7b3f4-18ae-11ef-a7a6-0242ac150002', '1', '8fe7b3f4-18ae-11ef-a7a6-0242ac150002')
-- ON CONFLICT DO NOTHING;
