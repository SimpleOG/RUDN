-- name: CreateGroupsCourse :one
INSERT INTO courses_groups (
    course_name,
    groups_id

) VALUES (
             $1, $2
         ) RETURNING *;

-- name: GetGroupsCourse :one
SELECT * FROM courses_groups
WHERE course_name =$1 and groups_id=$2  LIMIT 1;

-- name: ListAllGroupsCourse :many
SELECT * FROM courses_groups
ORDER BY course_name;
--LIMIT $2
-- OFFSET $3;


