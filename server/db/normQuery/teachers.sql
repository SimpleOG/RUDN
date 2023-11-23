-- name: CreateTeacher :one
INSERT INTO teachers (
    full_name,
    department,
    post,
    conditions
) VALUES (
             $1, $2,$3,$4
         ) RETURNING *;

-- name: GetTeacher :one
SELECT * FROM teachers
WHERE full_name = $1 LIMIT 1;

-- name: ListAllTeachers :many
SELECT * FROM teachers
ORDER BY full_name;


