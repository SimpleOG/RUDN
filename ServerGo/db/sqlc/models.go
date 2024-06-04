// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import ()

type DisciplineOrTypeOfAcademicWork struct {
	ID                                      int32  `json:"id"`
	Block                                   string `json:"block"`
	Component                               string `json:"component"`
	NVRup                                   string `json:"n_v_rup"`
	DopInfo                                 string `json:"dop_info"`
	NameOfTheDisciplineOrTypeOfAcademicWork string `json:"name_of_the_discipline_or_type_of_academic_work"`
}

type EducationalProgram struct {
	ID                  int32  `json:"id"`
	TheFormOfEducation  string `json:"the_form_of_education"`
	LevelOfOp           string `json:"level_of_op"`
	TheCodeOfTheOopRudn string `json:"the_code_of_the_oop_rudn"`
	DirectionCode       string `json:"direction_code"`
	NameOfTheProgram    string `json:"name_of_the_program"`
}

type InformationAboutPp struct {
	ID                int32  `json:"id"`
	Department        string `json:"department"`
	Post              string `json:"post"`
	TermsOfAttraction string `json:"terms_of_attraction"`
	FullName          string `json:"full_name"`
	ASpecialFeature   string `json:"a_special_feature"`
}

type KW struct {
	ID                     int32  `json:"id"`
	SemesterOrModule       string `json:"semester_or_module"`
	WeeksPerSemesterModule int32  `json:"weeks_per_semester_module"`
	TypeOfEducationalWork  string `json:"type_of_educational_work"`
	LectureHours           int32  `json:"lecture_hours"`
	LaboratoriesHours      int32  `json:"laboratories_hours"`
	PractiseHours          int32  `json:"practise_hours"`
	TypeOfPaOrGia          string `json:"type_of_pa_or_gia"`
	KwCourseWorks          string `json:"kw_course_works"`
	KwCourseProjects       string `json:"kw_course_projects"`
	CourseUchAveZeOnRup    string `json:"course_uch_ave_ze_on_rup"`
	PrZeOnRup              string `json:"pr_ze_on_rup"`
	NirZeByRup             string `json:"nir_ze_by_rup"`
}

type Semester struct {
	ID                        int32   `json:"id"`
	SemesterType              string  `json:"semester_type"`
	AuditoriumWork            float64 `json:"auditorium_work"`
	PairsPerWeek              float64 `json:"pairs_per_week"`
	ExtracurricularActivities float64 `json:"extracurricular_activities"`
}

type TheAmountOfTeachingWorkOfTheTeachingStaff struct {
	ID                                              int32   `json:"id"`
	Lectures                                        float64 `json:"lectures"`
	PracticeOrSeminars                              float64 `json:"practice_or_seminars"`
	LabWorksOrClinicalClasses                       float64 `json:"lab_works_or_clinical_classes"`
	CurrentControl                                  float64 `json:"current_control"`
	InterimCertificationPoForBrs                    float64 `json:"interim_certification_po_for_brs"`
	RegistrationOfPaResults                         float64 `json:"registration_of_pa_results"`
	OngoingConsultationsOnTheDiscipline             float64 `json:"ongoing_consultations_on_the_discipline"`
	CourseWorks                                     float64 `json:"course_works"`
	CourseProjects                                  float64 `json:"course_projects"`
	EducationalPractice                             float64 `json:"educational_practice"`
	ProcPedagogicalAndPreGraduatePractices          float64 `json:"proc_pedagogical_and_pre_graduate_practices"`
	Nir                                             float64 `json:"nir"`
	PracticesIncludingResearchOfDigitalMagistracies float64 `json:"practices_including_research_of_digital_magistracies"`
	ReviewingTheAbstractsOfGraduateStudents         float64 `json:"reviewing_the_abstracts_of_graduate_students"`
	CandidatesExam                                  float64 `json:"candidates_exam"`
	ScientificGuidance                              float64 `json:"scientific_guidance"`
	TheLeadershipOfTheWrcOrTheNkr                   float64 `json:"the_leadership_of_the_wrc_or_the_nkr"`
	ReviewOfTheWrc                                  float64 `json:"review_of_the_wrc"`
	Gek                                             float64 `json:"gek"`
	Total                                           float64 `json:"total"`
}

type TheContingentOfStudent struct {
	ID          int32  `json:"id"`
	GroupName   string `json:"group_name"`
	Code        string `json:"code"`
	GroupNumber string `json:"group_number"`
	OfGroups    string `json:"of_groups"`
	Subgroups   string `json:"subgroups"`
	TotalPeople string `json:"total_people"`
	Rf          string `json:"rf"`
	Foreign     string `json:"foreign"`
	Standard    string `json:"standard"`
	Calculated  string `json:"calculated"`
	Pk          string `json:"pk"`
}

type Together struct {
	ProgramID    int32 `json:"program_id"`
	DisciplineID int32 `json:"discipline_id"`
	TeacherID    int32 `json:"teacher_id"`
	GroupID      int32 `json:"group_id"`
	KWID         int32 `json:"k_w_id"`
	AmountID     int32 `json:"amount_id"`
	SemesterID   int32 `json:"semester_id"`
}