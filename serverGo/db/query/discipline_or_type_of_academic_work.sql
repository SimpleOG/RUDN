-- name: Create_Discipline_or_type_of_academic_work :one

INSERT INTO "discipline_or_type_of_academic_work" (
              "block"                                          ,
              "component"                                      ,
              "n_v_rup"                                        ,
              "dop_info"                                       ,
              "name_of_the_discipline_or_type_of_academic_work")
VALUES ($1, $2, $3, $4,$5)
RETURNING *;

-- name: Get_Discipline_or_type_of_academic_work :one

SELECT *
FROM "discipline_or_type_of_academic_work"
WHERE "id" = $1
LIMIT 1;


-- name: List_All_Teacher_Disciplines :many

SELECT  type_of_educational_work,name_of_the_discipline_or_type_of_academic_work,total,group_name from discipline_or_type_of_academic_work d
join together t on d.id = t.discipline_id
join k_w kw on t.k_w_id = kw.id
join the_amount_of_teaching_work_of_the_teaching_staff taotwotts on t.amount_id = taotwotts.id
join "information_about_PPS" iaP on iaP.id = t.teacher_id and iap.full_name='Разумчик Р.В.'
join the_contingent_of_students tcos on t.group_id = tcos.id
join semester s on s.id = t.semestr_id and s.semester_type='Весенний';
