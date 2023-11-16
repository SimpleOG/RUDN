CREATE TABLE "teachers" (
                            "id" serial NOT NULL,
                            "full_name" varchar UNIQUE PRIMARY KEY NOT NULL,
                            "department" varchar  NOT NULL,
                            "post" varchar  NOT NULL,
                            "conditions"  varchar  NOT NULL
);

CREATE TABLE "groups" (
                          "id" serial PRIMARY KEY NOT NULL,
                          "code" varchar NOT NULL,
                          "number" integer  NOT NULL,
                          "name" varchar  NOT NULL
);

CREATE TABLE "Educational program" (
                           "id" serial NOT NULL,
                           "name" varchar PRIMARY KEY NOT NULL,
                           "lecture_hours" integer NOT NULL,
                           "laboratories_hours" integer NOT NULL,
                           "practise_hours" integer NOT NULL
);

CREATE TABLE "teachers_courses" (
                                    "teachers_name" varchar NOT NULL,
                                    "course_name" varchar NOT NULL
);

CREATE TABLE "students" (
                            "id" serial PRIMARY KEY NOT NULL,
                            "full_name" varchar NOT NULL ,
                            "group_id" integer NOT NULL
);

CREATE TABLE "courses_groups" (
                                   "course_name" varchar NOT NULL,
                                   "groups_id" integer NOT NULL

);

CREATE INDEX ON "teachers" ("full_name");

CREATE INDEX ON "groups" ("name");

CREATE INDEX ON "courses" ("name");

ALTER TABLE "teachers_courses" ADD FOREIGN KEY ("teachers_name") REFERENCES "teachers" ("full_name");

ALTER TABLE "teachers_courses" ADD FOREIGN KEY ("course_name") REFERENCES "courses" ("name");

ALTER TABLE "students" ADD FOREIGN KEY ("group_id") REFERENCES "groups" ("id");

ALTER TABLE "courses_groups" ADD FOREIGN KEY ("course_name") REFERENCES "courses" ("name");


ALTER TABLE "courses_groups" ADD FOREIGN KEY ("groups_id") REFERENCES "groups" ("id");


