CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "user" (
                        id uuid PRIMARY KEY DEFAULT uuid_generate_v1(),
                        name varchar(255),
                        email varchar(255) UNIQUE NOT NULL,
                        age int
)
