-- NOTE Create table users
CREATE TABLE IF NOT EXISTS "public"."users" (
    "id" uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    "username" varchar(15) NOT NULL,
    "name" varchar(50) NOT NULL,
    "password" varchar(255) NOT NULL,
    "created_at" timestamptz(6),
    "updated_at" timestamptz(6)
);