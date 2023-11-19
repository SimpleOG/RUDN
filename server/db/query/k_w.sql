-- name: Create_k_w :one
INSERT INTO k_w (
  "Semester ; Module" ,
  "Weeks per semester (module)"  ,
  "Type_of_educational_work"     ,
  "lecture_hours"                ,
  "laboratories_hours"           ,
  "practise_hours"               ,
  "Type of PA or GIA"            ,
  "Course. works"                ,
  "Course. Course. projects"     ,
  "Course. Uch. ave. (ZE on RUP)",
  "Pr. pr. (ZE on RUP)"          ,
  "NIR (ZE by RUP)")
VALUES ($1, $2, $3, $4,$5,$6,$7,$8,$9,$10,$11,$12)
RETURNING *;

-- name: Get_k_w :one
SELECT *
FROM k_w
WHERE Type_of_educational_work = $1
LIMIT 1;

-- name: ListAll_k_w :many
SELECT *
FROM k_w
ORDER BY Type_of_educational_work;



