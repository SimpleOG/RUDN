-- name: Create_information_about_PPS :one

INSERT INTO "information_about_PPS" ("department",
                                     "post",
                                     "terms_of_attraction",
                                     "full_name",
                                     "a_special_feature")
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: Get_information_about_PPS :many

select distinct full_name ,department,post,terms_of_attraction from "information_about_PPS";




