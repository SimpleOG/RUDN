CREATE TABLE "educational_program"
(
    "The_code_of_the_OOP_RUDN" varchar NOT NULL,
    "Direction_code"           varchar NOT NULL,
    "Name_of_the_program"      integer NOT NULL,
    "discipline"               varchar NOT NULL
);


CREATE TABLE "discipline_or_type_of_academic_work"
(
    "Block"                                           varchar NOT NULL,
    "Component"                                       varchar NOT NULL,
    "№_v_RUP"                                         varchar,
    "Name_of_the_discipline_or_type_of_academic_work" varchar NOT NULL,
    "dop_info"                                        varchar NOT NULL
);

CREATE TABLE "k_w"
(
    "Semester ; Module"             varchar NOT NULL,
    "Weeks per semester (module)"   integer NOT NULL,
    "Type_of_educational_work"      varchar NOT NULL,
    "lecture_hours"                 integer,
    "laboratories_hours"            integer,
    "practise_hours"                integer,
    "Type of PA or GIA"             varchar,
    "Course. works"                 varchar,
    "Course. Course. projects"      varchar,
    "Course. Uch. ave. (ZE on RUP)" varchar,
    "Pr. pr. (ZE on RUP)"           varchar,
    "NIR (ZE by RUP)"               varchar

);


CREATE TABLE "the_contingent_of_students"
(
    "code"         varchar NOT NULL,
    work_id integer not null ,
    "Group number" varchar NOT NULL,
    "Of groups"    varchar NOT NULL,
    "subgroups"    varchar NOT NULL,
    "total_people" varchar NOT NULL,
    "RF"           varchar NOT NULL,
    "Foreign"      varchar NOT NULL,
    "Standard"     varchar NOT NULL,
    "Calculated"   varchar NOT NULL,
    "ПК"           varchar NOT NULL
);

CREATE TABLE "information_about_PPS"
(
    "department"          varchar NOT NULL,
    "post"                varchar NOT NULL,
    "terms of attraction" varchar NOT NULL,
    "full_name"           varchar NOT NULL,
    "A special feature"   varchar

);

CREATE TABLE "the_amount_of_teaching_work_of_the_teaching_staff"
(   "id" serial primary key not null ,
    "Lectures"                                               varchar not null,
    "Practice / Seminars"                                    varchar not null,
    "Lab. works / Clinical classes"                          varchar not null,
    "Current control"                                        varchar not null,
    "Interim certification (PO) for BRS"                     varchar not null,
    "Registration of PA results"                             varchar not null,
    "Ongoing consultations on the discipline"                varchar not null,
    "course_works"                                           varchar not null,
    "Course projects"                                        varchar not null,
    "Educational practice"                                   varchar not null,
    "Proc., pedagogical and pre-graduate practices"          varchar not null,
    "NIR"                                                    varchar not null,
    "Practices (including research) of digital magistracies" varchar not null,
    "Reviewing the abstracts of graduate students"           varchar not null,
    "Candidate''s exam"                                      varchar not null,
    "Scientific guidance"                                    varchar not null,
    "The leadership of the WRC or the NKR"                   varchar not null,
    "Review of the WRC"                                      varchar not null,
    "GEK"                                                    varchar not null

);

ALTER TABLE "the_contingent_of_students" ADD FOREIGN KEY ("work_id") REFERENCES "the_amount_of_teaching_work_of_the_teaching_staff" ("id");

CREATE TABLE "educational_discipline"
(
    "program_name"    varchar NOT NULL,
    "discipline_name" varchar NOT NULL
);
CREATE TABLE "discipline_prep"
(
    "prep_name"       varchar,
    "discipline_name" varchar NOT NULL
);
CREATE TABLE "group_prep"
(
    "group_name" varchar NOT NULL,
    "prep_name"  varchar
);
CREATE TABLE "group_hours"
(
    "group_name" varchar not null,
    "hours"      float4
);
CREATE TABLE "discipline_group"
(
    "discipline_name" varchar NOT NULL,
    "group_name"      varchar NOT NULL
);