package database

import (
	"testing"

	"github.com/stretchr/testify/require"

	"online-shop/internal/config"

)

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)

	if err != nil {
		panic(err)
	}
}

func TestLoadConnectionPostgres(t *testing.T) {
	t.Run("success", func (t *testing.T) {
		db, err := ConnectionPostgres(config.Cfg.DB)
		require.NoError(t, err)
		require.NotNil(t, db)

	})
}