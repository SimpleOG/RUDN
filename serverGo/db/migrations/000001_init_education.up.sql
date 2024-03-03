CREATE TABLE "educational_program"
(
    id                         serial  not null primary key,
    the_form_of_education      varchar not null,
    level_of_op                varchar not null,
    "the_code_of_the_oop_rudn" varchar NOT NULL,
    "direction_code"           varchar NOT NULL,
    "name_of_the_program"      varchar NOT NULL
);
CREATE TABLE "discipline_or_type_of_academic_work"
(
    id                                                serial  not null primary key,
    "block"                                           varchar NOT NULL,
    "component"                                       varchar NOT NULL,
    "n_v_rup"                                         varchar NOT NULL,
    "dop_info"                                        varchar NOT NULL,
    "name_of_the_discipline_or_type_of_academic_work" varchar NOT NULL
);
CREATE TABLE "information_about_PPS"
(
    id                    serial  not null primary key,
    "department"          varchar NOT NULL,
    "post"                varchar NOT NULL,
    "terms_of_attraction" varchar NOT NULL,
    "full_name"           varchar NOT NULL,
    "a_special_feature"   varchar NOT NULL

);
CREATE TABLE "the_contingent_of_students"
(
    id             serial  not null primary key,
    "group_name"   varchar not null,
    "code"         varchar NOT NULL,
    "group_number" varchar NOT NULL,
    "of_groups"    varchar NOT NULL,
    "subgroups"    varchar NOT NULL,
    "total_people" varchar NOT NULL,
    "rf"           varchar NOT NULL,
    "foreign"      varchar NOT NULL,
    "standard"     varchar NOT NULL,
    "calculated"   varchar NOT NULL,
    "pk"           varchar NOT NULL
);

CREATE TABLE "k_w"
(
    id                          serial  not null primary key,
    "semester_or_module"        varchar NOT NULL,
    "weeks_per_semester_module" integer NOT NULL,
    "type_of_educational_work"  varchar NOT NULL,
    "lecture_hours"             integer NOT NULL,
    "laboratories_hours"        integer NOT NULL,
    "practise_hours"            integer NOT NULL,
    "type_of_pa_or_gia"         varchar NOT NULL,
    "course_works"              varchar NOT NULL,
    "course_projects"           varchar NOT NULL,
    "course_uch_ave_ze_on_rup"  varchar NOT NULL,
    "pr_ze_on_rup"              varchar NOT NULL,
    "nir_ze_by_rup"             varchar NOT NULL

);

CREATE TABLE "the_amount_of_teaching_work_of_the_teaching_staff"
(
    id                                                     serial not null primary key,
    "lectures"                                             float  not null,
    "practice_or_seminars"                                 float  not null,
    "lab_works_or_clinical_classes"                        float  not null,
    "current_control"                                      float  not null,
    "interim_certification_po_for_brs"                     float  not null,
    "registration_of_pa_results"                           float  not null,
    "ongoing_consultations_on_the_discipline"              float  not null,
    "course_works"                                         float  not null,
    "course_projects"                                      float  not null,
    "educational_practice"                                 float  not null,
    "proc_pedagogical_and_pre_graduate_practices"          float  not null,
    "nir"                                                  float  not null,
    "practices_including_research_of_digital_magistracies" float  not null,
    "reviewing_the_abstracts_of_graduate_students"         float  not null,
    "candidates_exam"                                      float  not null,
    "scientific_guidance"                                  float  not null,
    "the_leadership_of_the_wrc_or_the_nkr"                 float  not null,
    "review_of_the_wrc"                                    float  not null,
    "gek"                                                  float  not null,
    "total"                                                float  not null

);
CREATE TABLE semester
(
    id                         serial  not null primary key,
    semester_type              varchar not null,
    auditorium_work            float   not null,
    pairs_per_week             float   not null,
    extracurricular_activities float   not null
);

CREATE TABLE "together"
(
    program_id    integer references educational_program ("id")                             not null,
    discipline_id integer references discipline_or_type_of_academic_work ("id")             not null,
    teacher_id    integer references "information_about_PPS" ("id")                         not null,
    group_id      integer references the_contingent_of_students ("id")                      not null,
    k_w_id        integer references k_w (id)                                               not null,
    amount_id     integer references the_amount_of_teaching_work_of_the_teaching_staff (id) not null
);