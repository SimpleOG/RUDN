CREATE TABLE together
(
    program_id    integer not null,
    discipline_id integer not null,
    teacher_id    integer not null,
    group_id      integer not null,
    k_w_id        integer not null,
    amount_id     integer not null

);
CREATE TABLE program_group
(
    name_of_the_program varchar not null,
    group_name          varchar not null
);
CREATE TABLE discipline_group
(
    discipline_name varchar not null,
    group_name      varchar not null
);
CREATE TABLE teacher_group
(
    teacher_name varchar not null,
    group_name   varchar not null
);
CREATE TABLE group_kw
(
    kw_id      integer not null,
    group_name varchar not null
);
CREATE TABLE group_hours_discipline
(
    discpline_name varchar not null,
    group_name     varchar not null,
    amount_id      integer not null
);



