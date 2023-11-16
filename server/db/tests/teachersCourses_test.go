package db

import (
	"context"
	"github.com/stretchr/testify/require"
	db "rudnWebApp/db/sqlc"
	"testing"
)

func CreateRandomTeachersCourse(t *testing.T) db.TeachersCourse {
	arg1 := CreateRandomTeacher(t)
	arg2 := CreateRandomCourse(t)
	params := db.CreateTeachersCourseParams{
		TeachersName: arg1.FullName,
		CourseName:   arg2.Name,
	}
	course, err := testQueries.CreateTeachersCourse(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, course)
	require.Equal(t, course.TeachersName, arg1.FullName)
	require.Equal(t, course.CourseName, arg2.Name)
	return course
}
func TestCreateRandomTeachersCourse(t *testing.T) {
	CreateRandomTeachersCourse(t)
}
func TestGetTeachersCourse(t *testing.T) {
	arg := CreateRandomTeachersCourse(t)
	params := db.GetTeachersCourseParams{
		TeachersName: arg.TeachersName,
		CourseName:   arg.CourseName,
	}
	getCourse, err := testQueries.GetTeachersCourse(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, getCourse)
	require.Equal(t, getCourse.TeachersName, arg.TeachersName)
	require.Equal(t, getCourse.CourseName, arg.CourseName)
}

func TestListAllTeachersCourses(t *testing.T) {
	gorutineCreate(func() { CreateRandomTeachersCourse(t) })
	teacher := CreateRandomTeacher(t)
	list, err := testQueries.ListAllTeachersCourses(context.Background(), teacher.FullName)
	require.NoError(t, err)
	for _, v := range list {
		require.Equal(t, teacher.FullName, v.TeachersName)
		require.NotEmpty(t, v)
	}

}
