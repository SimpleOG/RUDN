package db

import (
	"context"
	"github.com/stretchr/testify/require"
	db "rudnWebApp/db/sqlc"
	"rudnWebApp/util"
	"testing"
)

func CreateRandomTeacher(t *testing.T) db.Teacher {
	arg := db.CreateTeacherParams{
		FullName:   util.RandomName(),
		Department: util.RandomDepartment(),
	}
	teacher, err := testQueries.CreateTeacher(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, teacher)
	require.NotZero(t, teacher.ID)
	require.Equal(t, arg.FullName, teacher.FullName)
	require.Equal(t, arg.Department, teacher.Department)
	return teacher
}
func TestCreateRandomTeacher(t *testing.T) {
	CreateRandomTeacher(t)
}
func TestGetTeacher(t *testing.T) {
	teacher := CreateRandomTeacher(t)
	get, err := testQueries.GetTeacher(context.Background(), teacher.FullName)
	require.NoError(t, err)
	require.NotEmpty(t, get)
	require.Equal(t, teacher.FullName, get.FullName)
	require.Equal(t, teacher.ID, get.ID)
	require.Equal(t, teacher.Department, get.Department)
}
func TestListAllTeachers(t *testing.T) {
	gorutineCreate(func() { CreateRandomTeacher(t) })
	list, err := testQueries.ListAllTeachers(context.Background())
	require.NoError(t, err)
	for _, v := range list {
		require.NotEmpty(t, v)
		require.NotZero(t, v.ID)
	}
}
