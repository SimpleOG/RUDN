-- name: Create_Semester :one

INSERT INTO "semester" (
     auditorium_work, pairs_per_week, extracurricular_activities,semester_type
   )
VALUES ($1, $2, $3,$4)
RETURNING *;

-- name: Get_Semester :one

SELECT *
FROM "semester"
WHERE "id" = $1
LIMIT 1;

SELECT  "the_code_of_the_oop_rudn" , "direction_code" , "name_of_the_program" , "full_name" , "lectures" , "practice_or_seminars" , "lab_works_or_clinical_classes" , "total"
from discipline_or_type_of_academic_work d
         join together t on d.id = t.discipline_id
         join  k_w kw on t.k_w_id = kw.id
         join the_amount_of_teaching_work_of_the_teaching_staff taotwotts on t.amount_id = taotwotts.id
         join "information_about_PPS" iaP on iaP.id = t.teacher_id and iap.full_name= $1
         join the_contingent_of_students tcos on t.group_id = tcos.id
         join educational_program ep on t.program_id = ep.id
         join semester s on s.id=t.semestr_id and semester_type=$2

;