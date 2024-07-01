package utility

import (
	"log"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		publicId := uuid.NewString()
		secret := "Ini secret"
		tokenString, err := GenerateToken(publicId, secret, "user")
		require.Nil(t, err)
		require.NotEmpty(t, tokenString)
		log.Printf(tokenString)
	})
}

func TestVerifyToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		publicId := uuid.NewString()
		role := "user"
		secret := "Ini secret"
		tokenString, err := GenerateToken(publicId, secret, role)
		require.Nil(t, err)
		require.NotEmpty(t, tokenString)

		jwtId, jwtRole, err := ValidateToken(tokenString, secret)
		require.Nil(t, err)
		require.NotEmpty(t, jwtId)
		require.NotEmpty(t, jwtRole)

		require.Equal(t, publicId, jwtId)
		require.Equal(t, role, jwtRole)

		log.Printf("Token: %s", tokenString)
	})
}
