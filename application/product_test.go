package application_test

import (
	"github.com/phelliperodrigues/go-hexagonal/application"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hellp"
	product.Status = application.DISABLED
	product.Prince = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Prince = 0

	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}
