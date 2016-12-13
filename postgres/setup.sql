CREATE TABLE "public"."time_session" (
    "id" text,
    "name" text,
    "duration" int,
    "created_at" timestamp,
    PRIMARY KEY ("id")
);

INSERT INTO time_session VALUES('be1ae26e-c0b4-11e6-a4a6-cec0c932ce01', 'Build Website', 3700, (now() - INTERVAL '29 days'));
INSERT INTO time_session VALUES('be1ae4e4-c0b4-11e6-a4a6-cec0c932ce01', 'Develop Tooling', 37000, now() - INTERVAL '10 days');
INSERT INTO time_session VALUES('be1ae5e8-c0b4-11e6-a4a6-cec0c932ce01', 'Design Logo', 1000, now() - INTERVAL '23 days');
INSERT INTO time_session VALUES('be1ae6c4-c0b4-11e6-a4a6-cec0c932ce01', 'Refactor Stuff', 700, now() - INTERVAL '2 days');
INSERT INTO time_session VALUES('be1ae7a0-c0b4-11e6-a4a6-cec0c932ce01', 'Study Go', 600, now() - INTERVAL '1 days');
INSERT INTO time_session VALUES('be1ae868-c0b4-11e6-a4a6-cec0c932ce01', 'Deploy Time Tracker', 100, now() - INTERVAL '3 days');
INSERT INTO time_session VALUES('be1aeba6-c0b4-11e6-a4a6-cec0c932ce01', 'Unninstal Eclipse', 33700, now());
