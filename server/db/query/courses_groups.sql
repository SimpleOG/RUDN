-- name: CreateGroupsCourse :one
INSERT INTO courses_groups (
    course_name,
    groups_id

) VALUES (
             $1, $2
         ) RETURNING *;

-- name: GetGroupsCourse :one
SELECT * FROM courses_groups
WHERE course_name =$1 and groups_id=$2  LIMIT 1;

-- name: ListAllGroupCourses :many
SELECT * FROM courses_groups join groups as g on groups_id=$1 and
                                                 g.id=courses_groups.groups_id
    join courses c on c.name = courses_groups.course_name

ORDER BY course_name ;

-- name: ListAllCourseGroups :many
SELECT * FROM courses_groups join courses c on courses_groups.course_name=$1 and
                               c.name = courses_groups.course_name
                             join groups as g  on g.id = courses_groups.groups_id

ORDER BY course_name;
