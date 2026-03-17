CREATE TABLE if not EXISTS "user_role"(
    "id" serial PRIMARY KEY,
    "role" VARCHAR(80),
    "user_id" INT NOT NULL,
    CONSTRAINT fk_user_roles_user
        FOREIGN KEY ("user_id")
        REFERENCES "user"("id")
);
