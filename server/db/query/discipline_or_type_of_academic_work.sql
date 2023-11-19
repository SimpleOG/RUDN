-- name: CreateDiscipline_or_type_of_academic_work :one
INSERT INTO "discipline_or_type_of_academic_work" (
               "Block",
               "Component",
               "№_v_RUP",
               "Name_of_the_discipline_or_type_of_academic_work",
               "dop_info")
VALUES ($1, $2, $3, $4,$5)
RETURNING *;

-- name: GetDiscipline_or_type_of_academic_work :one
SELECT *
FROM discipline_or_type_of_academic_work
WHERE Name_of_the_discipline_or_type_of_academic_work = $1
LIMIT 1;

-- name: ListAllCourses :many
SELECT *
FROM discipline_or_type_of_academic_work
ORDER BY Name_of_the_discipline_or_type_of_academic_work;



