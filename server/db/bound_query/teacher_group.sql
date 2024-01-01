-- name: Create_teacher_group :one
INSERT INTO  "teacher_group"(
    teacher_name,
    discipline_name,
    group_name
)

VALUES ($1, $2,$3)
RETURNING *;

-- name: Get_teacher_group :one
SELECT *
FROM teacher_group
WHERE teacher_name = $1
  and group_name=$2 and discipline_name=$3;


