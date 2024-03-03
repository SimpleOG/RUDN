-- name: Create_the_contingent_of_students :one

INSERT INTO "the_contingent_of_students"("group_name",
                                         "code",
                                         "group_number",
                                         "of_groups",
                                         "subgroups",
                                         "total_people",
                                         "rf",
                                         "foreign",
                                         "standard",
                                         "calculated",
                                         "pk")
VALUES ($1  ||'-'|| $2, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: Get_the_contingent_of_students :one

SELECT *
FROM the_contingent_of_students
WHERE "id" = $1
LIMIT 1;





