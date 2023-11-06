package db

import (
	"context"
	"github.com/stretchr/testify/require"
	db "rudnWebApp/db/sqlc"
	"testing"
)

func CreateRandomCoursesGroup(t *testing.T) db.CoursesGroup {
	course := CreateRandomCourse(t)
	group := CreateRandomGroup(t)
	arg := db.CreateGroupsCourseParams{
		CourseName: course.Name,
		GroupsID:   group.ID,
	}
	coursesGroup, err := testQueries.CreateGroupsCourse(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, coursesGroup)
	require.Equal(t, coursesGroup.CourseName, course.Name)
	require.Equal(t, coursesGroup.GroupsID, group.ID)
	return coursesGroup
}
func TestCreateCoursesGroup(t *testing.T) {
	CreateRandomCoursesGroup(t)
}

func TestGetCoursesGroup(t *testing.T) {
	course := CreateRandomCoursesGroup(t)
	arg := db.GetGroupsCourseParams{
		CourseName: course.CourseName,
		GroupsID:   course.GroupsID,
	}
	getCourse, err := testQueries.GetGroupsCourse(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, getCourse)
	require.Equal(t, getCourse.CourseName, arg.CourseName)
	require.Equal(t, getCourse.GroupsID, arg.GroupsID)

}
func TestListAllCoursesGroup(t *testing.T) {
	gorutineCreate(func() { CreateRandomCoursesGroup(t) }, t)
	list, err := testQueries.ListAllGroupsCourse(context.Background())
	require.NoError(t, err)
	for _, v := range list {
		require.NotEmpty(t, v)
	}
}
