-- name: Create_together :one
INSERT INTO "together"(program_id,
                       discipline_id,
                       group_id,
                       teacher_id,
                       k_w_id,
                       amount_id)

VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetAllInfo :one
select *
from together
         join educational_program ep
              on together.program_id = ep.id and ep.id=$1
         join k_w kw
              on kw.id = together.k_w_id
         join discipline_or_type_of_academic_work as dis
              on dis.id = together.discipline_id
         join "information_about_PPS" as i
              on i.id = together.teacher_id
         join the_amount_of_teaching_work_of_the_teaching_staff as t
              on t.id = together.amount_id
         join the_contingent_of_students as c
              on c.id = together.group_id


