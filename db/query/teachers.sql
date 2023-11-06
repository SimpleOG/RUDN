-- name: CreateTeacher :one
INSERT INTO teachers (
    full_name,
    department
) VALUES (
             $1, $2
         ) RETURNING *;

-- name: GetTeacher :one
SELECT * FROM teachers
WHERE full_name = $1 LIMIT 1;

-- name: ListAllTeachers :many
SELECT * FROM groups
ORDER BY name;
--LIMIT $2
-- OFFSET $3;


