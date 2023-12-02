-- name: Create_discipline_group :one
INSERT INTO  "discipline_group"(
    discipline_name,
    group_name
    )

VALUES ($1, $2)
RETURNING *;

-- name: Get_discipline_group :one
SELECT *
FROM discipline_group
WHERE
        discipline_name = $1
    and group_name=$2;



