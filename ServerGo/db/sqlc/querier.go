// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	Create_Discipline_or_type_of_academic_work(ctx context.Context, arg Create_Discipline_or_type_of_academic_workParams) (DisciplineOrTypeOfAcademicWork, error)
	Create_EducationalProgram(ctx context.Context, arg Create_EducationalProgramParams) (EducationalProgram, error)
	Create_Semester(ctx context.Context, arg Create_SemesterParams) (Semester, error)
	Create_information_about_pps(ctx context.Context, arg Create_information_about_ppsParams) (InformationAboutPp, error)
	Create_k_w(ctx context.Context, arg Create_k_wParams) (KW, error)
	Create_the_amount_of_teaching_work_of_the_teaching_staff(ctx context.Context, arg Create_the_amount_of_teaching_work_of_the_teaching_staffParams) (TheAmountOfTeachingWorkOfTheTeachingStaff, error)
	Create_the_contingent_of_students(ctx context.Context, arg Create_the_contingent_of_studentsParams) (TheContingentOfStudent, error)
	Create_together(ctx context.Context, arg Create_togetherParams) (Together, error)
	Get_Discipline_or_type_of_academic_work(ctx context.Context, id int32) (DisciplineOrTypeOfAcademicWork, error)
	Get_EducationalProgram(ctx context.Context, id int32) (EducationalProgram, error)
	Get_Semester(ctx context.Context, id int32) (Semester, error)
	Get_information_about_pps(ctx context.Context) ([]Get_information_about_ppsRow, error)
	Get_k_w(ctx context.Context, id int32) (KW, error)
	Get_the_amount_of_teaching_work_of_the_teaching_staff(ctx context.Context, id int32) (TheAmountOfTeachingWorkOfTheTeachingStaff, error)
	Get_the_contingent_of_students(ctx context.Context, id int32) (TheContingentOfStudent, error)
	List_All_Teacher_Disciplines(ctx context.Context, arg List_All_Teacher_DisciplinesParams) ([]List_All_Teacher_DisciplinesRow, error)
	Teacher_Info(ctx context.Context, fullName string) (Teacher_InfoRow, error)
}

var _ Querier = (*Queries)(nil)