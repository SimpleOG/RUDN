-- name: CreateGroup :one
INSERT INTO groups (
    name,
    code,
    number
) VALUES (
             $1, $2,$3
         ) RETURNING *;

-- name: GetGroup :one
SELECT * FROM groups
WHERE name = $1 LIMIT 1;

-- name: ListAllGroups :many
SELECT * FROM groups
ORDER BY name;



