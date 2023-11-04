CREATE TABLE "teachers" (
                            "id" serial PRIMARY KEY,
                            "full_name" text NOT NULL,
                            "age" integer NOT NULL
);

CREATE TABLE "students" (
                            "id" serial PRIMARY KEY,
                            "full_name" text NOT NULL,
                            "group_id" integer NOT NULL
);
CREATE TABLE "user" (

                        "username" varchar PRIMARY KEY,
                        "hashed_password" varchar NOT NULL,
                        "full_name" varchar NOT NULL,
                        "email" varchar UNIQUE NOT NULL
);
ALTER TABLE "teachers" ADD FOREIGN KEY ("full_name") REFERENCES "user" ("username");
ALTER TABLE "students" ADD FOREIGN KEY ("full_name") REFERENCES "user" ("username");
CREATE TABLE "groups" (
                          "id" serial PRIMARY KEY,
                          "name" text NOT NULL,
                          "number" int NOT NULL
);

CREATE TABLE "teachers_groups" (
                                   "teachers_id" integer NOT NULL,
                                   "groups_id" integer NOT NULL
);

ALTER TABLE "students" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "teachers_groups" ADD FOREIGN KEY ("teachers_id") REFERENCES "teachers" ("id");

ALTER TABLE "teachers_groups" ADD FOREIGN KEY ("groups_id") REFERENCES "groups" ("id");
