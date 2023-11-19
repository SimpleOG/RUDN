-- name: Create_the_contingent_of_students :one
INSERT INTO the_contingent_of_students (
                     "code",
                     "Group number",
                     "Of groups",
                     "subgroups",
                     "total_people",
                     "RF",
                     "Foreign",
                     "Standard",
                     "Calculated",
                     "ПК")
VALUES ($1, $2, $3, $4,$5,$6,$7,$8,$9,$10)
RETURNING *;

-- name: Get_the_contingent_of_students :one
SELECT *
FROM the_contingent_of_students
WHERE "Group number" = $1
LIMIT 1;

-- name: ListAll_the_contingent_of_students :many
SELECT *
FROM the_contingent_of_students
ORDER BY "Group number";



