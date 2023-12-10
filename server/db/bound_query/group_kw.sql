-- name: Create_group_kw :one
INSERT INTO  "group_kw"(
    kw_id,
    discipline_name,
    group_name
)

VALUES ($1, $2,$3)
RETURNING *;

-- name: Get_group_kw :one
SELECT *
FROM group_kw
WHERE discipline_name=$1
  and group_name=$2;



