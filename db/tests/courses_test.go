package db

import (
	"context"
	"github.com/stretchr/testify/require"
	db "rudnWebApp/db/sqlc"
	"rudnWebApp/util"
	"testing"
)

func CreateRandomCourse(t *testing.T) db.Course {
	arg := db.CreateCourseParams{
		Name:              util.RandomString(5),
		LectureHours:      int32(util.RandomInt(1, 5)),
		LaboratoriesHours: int32(util.RandomInt(1, 5)),
		PractiseHours:     int32(util.RandomInt(1, 5)),
	}
	course, err := testQueries.CreateCourse(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, course)
	require.Equal(t, arg.Name, course.Name)
	require.Equal(t, arg.LectureHours, course.LectureHours)
	require.Equal(t, arg.LaboratoriesHours, course.LaboratoriesHours)
	require.Equal(t, arg.PractiseHours, course.PractiseHours)
	require.NotZero(t, course.ID)
	return course
}
func TestCreateRandomCourse(t *testing.T) {
	CreateRandomCourse(t)
}
func TestGetCourse(t *testing.T) {
	course := CreateRandomCourse(t)
	getCourse, err := testQueries.GetCourse(context.Background(), course.Name)
	require.NoError(t, err)
	require.NotEmpty(t, getCourse)
	require.Equal(t, course.Name, getCourse.Name)
	require.Equal(t, course.LectureHours, getCourse.LectureHours)
	require.Equal(t, course.LaboratoriesHours, getCourse.LaboratoriesHours)
	require.Equal(t, course.PractiseHours, getCourse.PractiseHours)
	require.Equal(t, course.ID, getCourse.ID)

}
func TestListAllCourses(t *testing.T) {
	gorutineCreate(func() { CreateRandomCourse(t) }, t)
	list, err := testQueries.ListAllCourses(context.Background())
	require.NoError(t, err)
	for _, v := range list {
		require.NotEmpty(t, v)
		require.NotZero(t, v.ID)
	}
}
