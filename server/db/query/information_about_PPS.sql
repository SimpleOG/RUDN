-- name: Create_information_about_PPS :one
INSERT INTO "information_about_PPS" ("department",
                                     "post",
                                     "terms_of_attraction",
                                     "full_name",
                                     "a_special_feature")
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: Get_information_about_PPS :one
SELECT *
FROM "information_about_PPS"
WHERE full_name = $1
LIMIT 1;



