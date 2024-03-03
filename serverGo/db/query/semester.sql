-- name: Create_Semester :one

INSERT INTO "semester" (
     auditorium_work, pairs_per_week, extracurricular_activities,semester_type
   )
VALUES ($1, $2, $3,$4)
RETURNING *;

-- name: Get_Semester :one

SELECT *
FROM "semester"
WHERE "id" = $1
LIMIT 1;

