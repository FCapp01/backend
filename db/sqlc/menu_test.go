package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCreateMenu(t *testing.T) {
	args := CreateMenuParams{
		RestaurantID: 1,
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
}
