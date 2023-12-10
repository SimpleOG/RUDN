-- name: Create_Program_group :one
INSERT INTO  "program_group"(
    program_name,
    group_name
    )

VALUES ($1, $2)
RETURNING *;

-- name: Get_Program_group :one
SELECT *
FROM Program_group
WHERE program_name = $1 and group_name=$2;



