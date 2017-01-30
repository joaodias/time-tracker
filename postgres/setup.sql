CREATE TABLE "public"."time_session" (
    "id" text,
    "name" text,
    "duration" int,
    "created_at" timestamp,
    "user_id" text,
    "want_calendar" boolean,
    "initial_timestamp" text,
    PRIMARY KEY ("id")
);

CREATE TABLE "public"."auth_user" (
    "id" text,
    "name" text,
    "email" text,
    "access_token" text,
    "created_at" timestamp,
    "last_login" timestamp,
    PRIMARY KEY ("email")
);
