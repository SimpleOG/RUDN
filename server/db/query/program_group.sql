-- name: Create_Program_group :one
INSERT INTO  "program_group"(
    name_of_the_program,
    group_name
    )

VALUES ($1, $2)
RETURNING *;

-- name: Get_Program_group :one
SELECT *
FROM Program_group
WHERE name_of_the_program = $1 and group_name=$2;



