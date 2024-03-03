-- name: Create_the_amount_of_teaching_work_of_the_teaching_staff :one

INSERT INTO "the_amount_of_teaching_work_of_the_teaching_staff" ("lectures",
                                                                 "practice_or_seminars",
                                                                 "lab_works_or_clinical_classes",
                                                                 "current_control",
                                                                 "interim_certification_po_for_brs",
                                                                 "registration_of_pa_results",
                                                                 "ongoing_consultations_on_the_discipline",
                                                                 "course_works",
                                                                 "course_projects",
                                                                 "educational_practice",
                                                                 "proc_pedagogical_and_pre_graduate_practices",
                                                                 "nir",
                                                                 "practices_including_research_of_digital_magistracies",
                                                                 "reviewing_the_abstracts_of_graduate_students",
                                                                 "candidates_exam",
                                                                 "scientific_guidance",
                                                                 "the_leadership_of_the_wrc_or_the_nkr",
                                                                 "review_of_the_wrc",
                                                                 "gek",
                                                                 "total")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
RETURNING *;

-- name: Get_the_amount_of_teaching_work_of_the_teaching_staff :one

SELECT *
FROM the_amount_of_teaching_work_of_the_teaching_staff
WHERE id = $1
LIMIT 1;





