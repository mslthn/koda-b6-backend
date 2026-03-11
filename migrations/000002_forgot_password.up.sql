create table if not exists "forgot_password"(
    "id" serial primary key,
    "email" varchar(50),
    "otp_code" int,
    "created_at" timestamp default CURRENT_TIMESTAMP
);