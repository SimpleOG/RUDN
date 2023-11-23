-- name: Create_k_w :one
INSERT INTO "k_w" ("semester_or_Module",
                   "weeks_per_semester_module",
                   "type_of_educational_work",
                   "lecture_hours",
                   "laboratories_hours",
                   "practise_hours",
                   "type_of_PA_or_GIA",
                   "course_works",
                   "course_projects",
                   "course_Uch_ave_ZE_on_RUP",
                   "pr_ZE_on_RUP",
                   "NIR_ZE_by_RUP")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING *;

-- name: Get_k_w :one
SELECT *
FROM k_w
WHERE "type_of_educational_work" = $1
LIMIT 1;

