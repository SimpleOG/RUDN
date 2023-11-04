-- name: CreateTeacher :one
INSERT INTO teachers (
    full_name,
    age
) VALUES (
             $1, $2
         ) RETURNING *;

-- name: GetTeacher :one
SELECT * FROM teachers
WHERE id = $1 LIMIT 1;

-- name: ListAllTeachers :many
SELECT * FROM teachers
ORDER BY id;
--LIMIT $2
-- OFFSET $3;


