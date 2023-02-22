package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomMenu(t *testing.T) Menu {

	restaurant := createRandomRestaurant(t)

	args := CreateMenuParams{
		RestaurantID: restaurant.ID,
		Name:         gofakeit.Name(),
		Description:  sql.NullString{String: gofakeit.JobDescriptor(), Valid: true},
	}
	menu, err := testQueries.CreateMenu(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, menu)

	require.Equal(t, args.RestaurantID, menu.RestaurantID)
	require.Equal(t, args.Name, menu.Name)
	require.Equal(t, args.Description, menu.Description)

	require.NotZero(t, menu.ID)
	require.NotZero(t, menu.UpdatedAt)

	return menu
}

func TestCreateMenu(t *testing.T) {
	createRandomMenu(t)
}

func TestGetMenu(t *testing.T) {
	menu := createRandomMenu(t)

	menus, err := testQueries.GetMenu(context.Background(), menu.RestaurantID)

	require.NoError(t, err)
	require.NotEmpty(t, menus)

	require.Equal(t, menus[0].RestaurantID, menu.RestaurantID)
	require.Equal(t, menus[0].Name, menu.Name)
	require.Equal(t, menus[0].Description, menu.Description)

	require.Equal(t, menus[0].ID, menu.ID)
	require.Equal(t, menus[0].UpdatedAt, menu.UpdatedAt)
}

func TestDeleteMenu(t *testing.T) {
	menu := createRandomMenu(t)

	err := testQueries.DeleteMenu(context.Background(), menu.ID)

	require.NoError(t, err)
}
