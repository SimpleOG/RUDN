-- name: Create_k_w :one

INSERT INTO "k_w" ("semester_or_module",
                   "weeks_per_semester_module",
                   "type_of_educational_work",
                   "lecture_hours",
                   "laboratories_hours",
                   "practise_hours",
                   "type_of_pa_or_gia",
                   "course_works",
                   "course_projects",
                   "course_uch_ave_ze_on_rup",
                   "pr_ze_on_rup",
                   "nir_ze_by_rup")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: Get_k_w :one

SELECT *
FROM k_w
WHERE "id" = $1
LIMIT 1;

