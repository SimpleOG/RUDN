-- name: CreateCourse :one
INSERT INTO courses (
    name,
    lecture_hours,
    laboratories_hours,
    practise_hours
) VALUES (
             $1, $2,$3,$4
         ) RETURNING *;

-- name: GetCourse :one
SELECT * FROM courses
WHERE name = $1 LIMIT 1;

-- name: ListAllCourses :many
SELECT * FROM courses
ORDER BY name;



