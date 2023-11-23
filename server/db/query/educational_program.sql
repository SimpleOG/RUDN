-- name: Create_EducationalProgram :one
INSERT INTO "educational_program" ("the_code_of_the_OOP_RUDN",
                                   "direction_code",
                                   "name_of_the_program")

VALUES ($1, $2, $3)
RETURNING *;

-- name: Get_EducationalProgram :one
SELECT *
FROM educational_program
WHERE "name_of_the_program" = $1
LIMIT 1;


