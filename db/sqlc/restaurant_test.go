package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestCreateRestaurant(t *testing.T) {
	args := CreateRestaurantParams{
		Name:        gofakeit.Name(),
		Address:     gofakeit.Address().Address,
		PhoneNumber: gofakeit.Phone(),
		Email:       gofakeit.Email(),
	}

	restaurant, err := testQueries.CreateRestaurant(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, restaurant)

	require.Equal(t, args.Name, restaurant.Name)
	require.Equal(t, args.Address, restaurant.Address)
	require.Equal(t, args.PhoneNumber, restaurant.PhoneNumber)
	require.Equal(t, args.Email, restaurant.Email)

	require.NotZero(t, restaurant.ID)
}
