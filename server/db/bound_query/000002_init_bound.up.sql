CREATE TABLE program_group
(
    program_name varchar not null,
    group_name   varchar not null
);
CREATE TABLE discipline_group
(
    discipline_name varchar not null,
    group_name      varchar not null
);

CREATE TABLE teacher_group
(
    teacher_name    varchar not null,
    group_name      varchar not null,
    discipline_name varchar not null
);
CREATE TABLE group_kw

(
    discipline_name varchar not null,
    kw_id           integer not null,
    group_name      varchar not null
);
CREATE TABLE group_hours_discipline
(
    discipline_name varchar not null,
    group_name      varchar not null,
    amount_id     integer not null
);



