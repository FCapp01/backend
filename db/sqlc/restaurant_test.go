package db

import (
	"context"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func createRandomRestaurant(t *testing.T) Restaurant {
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

	return restaurant
}

func TestCreateRestaurant(t *testing.T) {
	createRandomRestaurant(t)
}

func TestGetRestaurant(t *testing.T) {
	restaurant := createRandomRestaurant(t)

	res, err := testQueries.GetRestaurant(context.Background(), restaurant.ID)

	require.NoError(t, err)
	require.NotEmpty(t, res)

	require.Equal(t, res.ID, restaurant.ID)
	require.Equal(t, res.Name, restaurant.Name)
	require.Equal(t, res.Address, restaurant.Address)
	require.Equal(t, res.PhoneNumber, restaurant.PhoneNumber)
	require.Equal(t, res.Email, restaurant.Email)
}

func TestDeleteRestaurant(t *testing.T) {
	restaurant := createRandomRestaurant(t)

	err := testQueries.DeleteRestaurant(context.Background(), restaurant.ID)

	require.NoError(t, err)
}
