package application_test

import (
	"github.com/phelliperodrigues/go-hexagonal/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := buildProduct()

	err := product.Enable()
	require.Nil(t, err)

	product.Prince = 0

	err = product.Enable()
	require.Equal(t, "the price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := buildProduct()
	product.Prince = 0
	err := product.Disable()
	require.Nil(t, err)

	product.Prince = 10

	err = product.Disable()
	require.Equal(t, "the price must be zero in order to have the product disable", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := buildProduct()
	id := uuid.NewV4().String()
	product.ID = id

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Status = ""
	_, err = product.IsValid()
	require.Nil(t, err)
	require.Equal(t, application.DISABLED, product.GetStatus())

	product.Name = ""
	_, err = product.IsValid()
	require.Equal(t, "Name: non zero value required", err.Error())

	product.Prince = -1
	_, err = product.IsValid()
	require.Equal(t, "the price must be greater or equals zero", err.Error())

}

func TestProduct_GetID(t *testing.T) {
	product := buildProduct()
	id := uuid.NewV4().String()
	product.ID = id
	require.Equal(t, id, product.GetID())
}

func TestProduct_GetName(t *testing.T) {
	product := buildProduct()
	require.Equal(t, "Hello", product.GetName())
}

func TestProduct_GetPrice(t *testing.T) {
	product := buildProduct()
	require.Equal(t, 10.0, product.GetPrice())
}

func TestProduct_GetStatus(t *testing.T) {
	product := buildProduct()
	require.Equal(t, application.DISABLED, product.GetStatus())
}

func buildProduct() application.Product {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Prince = 10
	return product
}
