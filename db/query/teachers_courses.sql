-- name: CreateTeachersCourse :one
INSERT INTO teachers_courses (
    teachers_name,
    course_name
) VALUES (
             $1, $2
         ) RETURNING *;

-- name: GetTeachersCourse :one
SELECT * FROM teachers_courses
WHERE teachers_name =$1 and course_name=$2 LIMIT 1;

-- name: ListAllTeachersCourses :many
SELECT * FROM teachers_courses
ORDER BY course_name;
--LIMIT $2
-- OFFSET $3;


