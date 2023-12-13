CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "username" varchar NOT NULL,
    "password" varchar NOT NULL
);

CREATE TABLE "user_geo" (
    "user_id" int references "users"("id"),
    "started_at" timestamp NOT NULL,
    "finished_at" timestamp NOT NULL,
    "location" jsonb NOT NULL
);