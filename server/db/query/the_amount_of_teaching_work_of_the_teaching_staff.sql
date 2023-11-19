-- name: Create_the_amount_of_teaching_work_of_the_teaching_staff :one
INSERT INTO the_amount_of_teaching_work_of_the_teaching_staff (
                     "Lectures",
                     "Practice / Seminars",
                     "Lab. works / Clinical classes",
                     "Current control",
                     "Interim certification (PO) for BRS",
                     "Registration of PA results",
                     "Ongoing consultations on the discipline",
                     "course_works",
                     "Course projects",
                     "Educational practice",
                     "Proc., pedagogical and pre-graduate practices",
                     "NIR",
                     "Practices (including research) of digital magistracies",
                     "Reviewing the abstracts of graduate students",
                     "Candidate''s exam",
                     "Scientific guidance",
                     "The leadership of the WRC or the NKR",
                     "Review of the WRC",
                     "GEK")
VALUES ($1, $2, $3, $4,$5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
RETURNING *;

-- name: Get_the_amount_of_teaching_work_of_the_teaching_staff :one
SELECT *
FROM the_amount_of_teaching_work_of_the_teaching_staff
WHERE id = $1
LIMIT 1;

-- name: List_the_amount_of_teaching_work_of_the_teaching_staff :many
SELECT *
FROM the_amount_of_teaching_work_of_the_teaching_staff
ORDER BY id;



