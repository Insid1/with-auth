CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "user"
(
    id         uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
    name       varchar(255),
    email      varchar(255) UNIQUE NOT NULL,
    age        int NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp DEFAULT now()
);


INSERT INTO "user" (id, name, email, age) VALUES ('8fe7b3f4-18ae-11ef-a7a6-0242ac150002', 'admin', 'admin@admin.admin', 0) ON CONFLICT DO NOTHING;
