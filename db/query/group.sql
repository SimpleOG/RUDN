-- name: CreateGroup :one
INSERT INTO groups (
    name,
    number
) VALUES (
             $1, $2
         ) RETURNING *;

-- name: GetGroup :one
SELECT * FROM groups
WHERE id = $1 LIMIT 1;

-- name: ListAllGroups :many
SELECT * FROM groups
ORDER BY id;
--LIMIT $2
   -- OFFSET $3;

-- name: ListAllTeachersGroups :many
SELECT * FROM teachers_groups as tg
INNER  join teachers as t on tg.teachers_id = t.id
INNER join groups as g on tg.groups_id = g.id
WHERE t.id=$1;
