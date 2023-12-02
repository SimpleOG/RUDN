-- name: Create_group_hours_discipline :one
INSERT INTO "group_hours_discipline"(
         group_name,
         discpline_name,
         amount_id)

VALUES ($1, $2, $3)
RETURNING *;

-- name: Get_group_hours_discipline :one
SELECT *
FROM group_hours_discipline
WHERE group_name = $1 and discpline_name=$2;



