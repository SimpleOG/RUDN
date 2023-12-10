-- name: Create_together :one
INSERT INTO "together"(program_name,
                       discipline_name,
                       group_name,
                       teacher_name,
                       k_w_id,
                       amount_id)

VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: Teacher_Info :one
select teacher_name,
       cast(sum(total) as float)                           as total,
       cast(sum(lectures) as float)                        as lectures,
       cast(sum("practice_or_Seminars") as float)          as practice,
       cast(sum("Lab_works_or_Clinical_classes") as float) as labs
from together as t
         join the_amount_of_teaching_work_of_the_teaching_staff as a on t.teacher_name = $1 and t.amount_id = a.id
GROUP BY teacher_name
;

-- name: Course_Info :many
select program_name,discipline_name,teacher_name,group_name,
       type_of_educational_work,lecture_hours,laboratories_hours,practise_hours,
       "type_of_PA_or_GIA",lectures,"practice_or_Seminars","Lab_works_or_Clinical_classes",
       total
from together t
         join    k_w kw on kw.id = t.k_w_id
         join the_amount_of_teaching_work_of_the_teaching_staff a on t.amount_id = a.id
where t.program_name=$1;