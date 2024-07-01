package auth

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // Import the postgres driver
	"github.com/stretchr/testify/require"

	"online-shop/external/database"
	"online-shop/infra/response"
	"online-shop/internal/config"

)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

	db, err := database.ConnectionPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestRegister_Success(t *testing.T) {
	ctx := context.Background()

	req := RegisterRequestPayload{
		Email:    fmt.Sprintf("%v@test.com", uuid.NewString()),
		Password: "password123",
	}

	err := svc.register(ctx, req)
	require.Nil(t, err)
}

func TestRegister_Fail(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		// Preparation for duplicate email
		email := fmt.Sprintf("%v@test.com", uuid.NewString())
		req := RegisterRequestPayload{
			Email:    email,
			Password: "password123",
		}

		// Register the user for the first time
		err := svc.register(context.Background(), req)
		require.Nil(t, err)

		// Try to register the same user again
		err = svc.register(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrEmailAlreadyUsed, err)
	})
}

func TestLoginSuccess(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		email := fmt.Sprintf("%v@test.com", uuid.NewString())
	password := "password123"
	req := RegisterRequestPayload{
		Email:    email,
		Password: password,
	}

	err := svc.register(context.Background(), req)
	require.Nil(t, err)

	reqLogin := LoginRequestPayload{
		Email:    email,
		Password: password,
	}

	token, err := svc.login(context.Background(), reqLogin)
	require.Nil(t, err)
	require.NotEmpty(t, token)
	log.Printf(token)
	})

}
