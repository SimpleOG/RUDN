-- name: Create_EducationalProgram :one

INSERT INTO "educational_program" ("the_form_of_education",
                                   "level_of_op",
                                   "the_code_of_the_oop_rudn",
                                   "direction_code",
                                   "name_of_the_program")

VALUES ($1, $2, $3,$4,$5)
RETURNING *;

-- name: Get_EducationalProgram :one

SELECT *
FROM educational_program
WHERE "id" = $1
LIMIT 1;


