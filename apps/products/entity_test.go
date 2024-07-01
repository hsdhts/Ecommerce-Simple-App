package products

import (
	"testing"

	"github.com/stretchr/testify/require"

	"online-shop/infra/response"
)

func TestValidateProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		product := Product{
			Name:  "Celana Jeans",
			Stock: 7,
			Price: 100_000,
		}

		err := product.Validate()
		require.Nil(t, err)
	})

	t.Run("product required", func(t *testing.T) {
		product := Product{
			Name:  "",
			Stock: 10,
			Price: 10_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})

	t.Run("product invalid", func(t *testing.T) {
		product := Product{
			Name:  "cel",
			Stock: 10,
			Price: 10_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductInvalid, err)
	})

	t.Run("stock invalid", func(t *testing.T) {
		product := Product{
			Name:  "celana jeans",
			Stock: 0,
			Price: 10_000,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrStockInvalid, err)
	})

	t.Run("price invalid", func(t *testing.T) {
		product := Product{
			Name:  "celana",
			Stock: 10,
			Price: 0,
		}

		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})
}
