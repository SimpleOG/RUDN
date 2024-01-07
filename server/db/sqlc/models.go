// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import ()

type DisciplineOrTypeOfAcademicWork struct {
	ID                                      int32  `json:"id"`
	Block                                   string `json:"block"`
	Component                               string `json:"component"`
	NVRUP                                   string `json:"n_v_RUP"`
	DopInfo                                 string `json:"dop_info"`
	NameOfTheDisciplineOrTypeOfAcademicWork string `json:"name_of_the_discipline_or_type_of_academic_work"`
}

type EducationalProgram struct {
	ID                  int32  `json:"id"`
	TheCodeOfTheOOPRUDN string `json:"the_code_of_the_OOP_RUDN"`
	DirectionCode       string `json:"direction_code"`
	NameOfTheProgram    string `json:"name_of_the_program"`
}

type InformationAboutPP struct {
	ID                int32  `json:"id"`
	Department        string `json:"department"`
	Post              string `json:"post"`
	TermsOfAttraction string `json:"terms_of_attraction"`
	FullName          string `json:"full_name"`
	ASpecialFeature   string `json:"a_special_feature"`
}

type KW struct {
	ID                     int32  `json:"id"`
	SemesterOrModule       string `json:"semester_or_Module"`
	WeeksPerSemesterModule int32  `json:"weeks_per_semester_module"`
	TypeOfEducationalWork  string `json:"type_of_educational_work"`
	LectureHours           int32  `json:"lecture_hours"`
	LaboratoriesHours      int32  `json:"laboratories_hours"`
	PractiseHours          int32  `json:"practise_hours"`
	TypeOfPAOrGIA          string `json:"type_of_PA_or_GIA"`
	CourseWorks            string `json:"course_works"`
	CourseProjects         string `json:"course_projects"`
	CourseUchAveZEOnRUP    string `json:"course_Uch_ave_ZE_on_RUP"`
	PrZEOnRUP              string `json:"pr_ZE_on_RUP"`
	NIRZEByRUP             string `json:"NIR_ZE_by_RUP"`
}

type TheAmountOfTeachingWorkOfTheTeachingStaff struct {
	ID                                              int32   `json:"id"`
	Lectures                                        float64 `json:"lectures"`
	PracticeOrSeminars                              float64 `json:"practice_or_Seminars"`
	LabWorksOrClinicalClasses                       float64 `json:"Lab_works_or_Clinical_classes"`
	CurrentControl                                  float64 `json:"current_control"`
	InterimCertificationPOForBRS                    float64 `json:"interim_certification_PO_for_BRS"`
	RegistrationOfPAResults                         float64 `json:"registration_of_PA_results"`
	OngoingConsultationsOnTheDiscipline             float64 `json:"ongoing_consultations_on_the_discipline"`
	CourseWorks                                     float64 `json:"course_works"`
	CourseProjects                                  float64 `json:"course_projects"`
	EducationalPractice                             float64 `json:"educational_practice"`
	ProcPedagogicalAndPreGraduatePractices          float64 `json:"proc_pedagogical_and_pre_graduate_practices"`
	NIR                                             float64 `json:"NIR"`
	PracticesIncludingResearchOfDigitalMagistracies float64 `json:"practices_including_research_of_digital_magistracies"`
	ReviewingTheAbstractsOfGraduateStudents         float64 `json:"reviewing_the_abstracts_of_graduate_students"`
	CandidatesExam                                  float64 `json:"candidates_exam"`
	ScientificGuidance                              float64 `json:"scientific_guidance"`
	TheLeadershipOfTheWRCOrTheNKR                   float64 `json:"the_leadership_of_the_WRC_or_the_NKR"`
	ReviewOfTheWRC                                  float64 `json:"review_of_the_WRC"`
	GEK                                             float64 `json:"GEK"`
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
	RF          string `json:"RF"`
	Foreign     string `json:"foreign"`
	Standard    string `json:"standard"`
	Calculated  string `json:"calculated"`
	PK          string `json:"PK"`
}

type Together struct {
	ProgramName    string `json:"program_name"`
	DisciplineName string `json:"discipline_name"`
	TeacherName    string `json:"teacher_name"`
	GroupName      string `json:"group_name"`
	KWID           int32  `json:"k_w_id"`
	AmountID       int32  `json:"amount_id"`
}
