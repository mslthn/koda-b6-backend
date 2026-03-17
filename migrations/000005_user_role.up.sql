CREATE table if not EXISTS "user_role"(
    "id" serial,
    "role" VARCHAR(80),
    "user_id" INT,
    CONSTRAINT user
        FOREIGN KEY ("user_id")
        REFERENCES "user"("id")
);