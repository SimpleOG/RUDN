-- name: CreateEducationalProgram :one
INSERT INTO "educational_program" (
    The_code_of_the_OOP_RUDN,
    Direction_code,
    Name_of_the_program,
    discipline
) VALUES (
             $1, $2,$3,$4
         ) RETURNING *;

-- name: GetEducationalProgram :one
SELECT * FROM educational_program
WHERE Name_of_the_program = $1 LIMIT 1;

-- name: ListAllCourses :many
SELECT * FROM educational_program
ORDER BY The_code_of_the_OOP_RUDN;



