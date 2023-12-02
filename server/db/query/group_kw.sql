-- name: Create_group_kw :one
INSERT INTO  "group_kw"(
    kw_id,
    group_name
)

VALUES ($1, $2)
RETURNING *;

-- name: Get_group_kw :one
SELECT *
FROM group_kw
WHERE kw_id = $1
  and group_name=$2;



