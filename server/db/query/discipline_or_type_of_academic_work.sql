-- name: Create_Discipline_or_type_of_academic_work :one
INSERT INTO "discipline_or_type_of_academic_work" (
               "block",
               "component",
               "n_v_RUP",
               "name_of_the_discipline_or_type_of_academic_work",
               "dop_info")
VALUES ($1, $2, $3, $4,$5)
RETURNING *;

-- name: Get_Discipline_or_type_of_academic_work :one
SELECT *
FROM "discipline_or_type_of_academic_work"
WHERE "name_of_the_discipline_or_type_of_academic_work" = $1
LIMIT 1;



