-- name: Create_together :one
INSERT INTO "together"(program_id,
                       discipline_id,
                       group_id,
                       teacher_id,
                       k_w_id,
                       amount_id,
                       semestr_id)

VALUES ($1, $2, $3, $4, $5, $6,$7)
RETURNING *;

-- name: Teacher_Info :one
select full_name, department,post,terms_of_attraction,
       round(cast(sum(total) as numeric),2 )                as total,
       round( cast(sum(lectures) as numeric)     ,2 )                     as lectures,
       round( cast(sum("practice_or_seminars") as numeric)  ,2 )          as practice,
       round( cast(sum("lab_works_or_clinical_classes") as numeric) ,2 )  as labs
from "information_about_PPS" i
         join together t  on  t.teacher_id = i.id
         join the_amount_of_teaching_work_of_the_teaching_staff as a on t.amount_id = a.id
where i.full_name=$1
group by  full_name,  department, post, terms_of_attraction ;


