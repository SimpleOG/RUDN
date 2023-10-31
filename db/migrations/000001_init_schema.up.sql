CREATE TABLE "teachers" (
                            "id" serial PRIMARY KEY,
                            "full_name" text
);

CREATE TABLE "students" (
                            "id" serial PRIMARY KEY,
                            "full_name" text,
                            "group_id" integer
);

CREATE TABLE "groups" (
                          "id" serial PRIMARY KEY,
                          "name" text,
                          "number" int
);

CREATE TABLE "teachers_groups" (
                                   "teachers_id" integer,
                                   "groups_id" integer
);

ALTER TABLE "groups" ADD FOREIGN KEY ("id") REFERENCES "students" ("group_id");

ALTER TABLE "teachers" ADD FOREIGN KEY ("id") REFERENCES "teachers_groups" ("teachers_id");

ALTER TABLE "groups" ADD FOREIGN KEY ("id") REFERENCES "teachers_groups" ("groups_id");
