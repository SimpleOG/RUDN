-- name: Create_information_about_PPS :one
INSERT INTO information_about_PPS (
                     "department",
                     "post",
                     "terms of attraction",
                     "full_name",
                     "A special feature")
VALUES ($1, $2, $3, $4,$5)
RETURNING *;

-- name: Get_information_about_PPS :one
SELECT *
FROM information_about_PPS
WHERE full_name = $1
LIMIT 1;

-- name: List_information_about_PPS :many
SELECT *
FROM information_about_PPS
ORDER BY full_name;



