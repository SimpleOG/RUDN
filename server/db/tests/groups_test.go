package db

import (
	"context"
	"github.com/stretchr/testify/require"
	db "rudnWebApp/db/sqlc"
	"rudnWebApp/util"
	"testing"
)

func CreateRandomGroup(t *testing.T) db.Group {
	arg := db.CreateGroupParams{
		Name:   util.RandomName(),
		Code:   util.RandomString(3),
		Number: int32(util.RandomInt(1, 3)),
	}
	group, err := testQueries.CreateGroup(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, group)
	require.NotZero(t, group.ID)
	require.Equal(t, group.Name, arg.Name)
	require.Equal(t, group.Number, arg.Number)
	require.Equal(t, group.Code, arg.Code)
	return group
}
func TestCreateRandomGroup(t *testing.T) {
	CreateRandomGroup(t)
}

func TestGetGroup(t *testing.T) {
	group := CreateRandomGroup(t)
	getGroup, err := testQueries.GetGroup(context.Background(), group.Name)
	require.NoError(t, err)
	require.NotEmpty(t, getGroup)
	require.NotZero(t, getGroup.ID)
	require.Equal(t, group.Name, getGroup.Name)
	require.Equal(t, group.ID, getGroup.ID)
	require.Equal(t, group.Code, getGroup.Code)
	require.Equal(t, group.Number, getGroup.Number)

}
func TestListAllGroups(t *testing.T) {
	gorutineCreate(func() { CreateRandomGroup(t) })
	list, err := testQueries.ListAllGroups(context.Background())
	require.NoError(t, err)
	for _, v := range list {
		require.NotEmpty(t, v)
		require.NotZero(t, v.ID)
	}
}
