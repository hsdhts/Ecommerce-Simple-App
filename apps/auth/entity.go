package auth

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"online-shop/infra/response"
	"online-shop/utility"

)

type Role string

const (
	ROLE_Admin Role = "admin"
	ROLE_User  Role = "user"
)

type AuthEntity struct {
	Id        int       `db:"id"`
	Email     string    `db:"email"`
	PublicId  uuid.UUID `db:"public_id"`
	Password  string    `db:"password"`
	Role      Role      `db:"role"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewFormRegisterUser(req RegisterRequestPayload) AuthEntity {
	return AuthEntity{
		Email:     req.Email,
		Password:  req.Password,
		PublicId:  uuid.New(),
		Role:      ROLE_User,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func NewFormLoginUser(req LoginRequestPayload) AuthEntity {
	return AuthEntity{
		Email:    req.Email,
		Password: req.Password}
}

// kebutuhan auth
// validasi object
func (a AuthEntity) Validate() (err error) {
	if err = a.ValidateEmail(); err != nil {
		return
	}

	if err = a.ValidatePassword(); err != nil {
		return
	}

	return
}

// validasi email
func (a AuthEntity) ValidateEmail() (err error) {
	if a.Email == "" {
		return response.ErrEmailRequired
	}

	emails := strings.Split(a.Email, "@")

	if len(emails) != 2 {
		return response.ErrEmailInvalid
	}
	return
}

// validasi password
func (a AuthEntity) ValidatePassword() (err error) {
	if a.Password == "" {
		return response.ErrPasswordRequired
	}

	if len(a.Password) < 6 {
		return response.ErrPasswordInvalidLength
	}
	return
}

func (a AuthEntity) IsExisting() bool {
	return a.Id != 0
}

func (a *AuthEntity) EncryptPassword(salt int) (err error) {
	encryptPass, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)

	if err != nil {
		return
	}

	a.Password = string(encryptPass)
	return nil

}

func (a AuthEntity) VerifyPasswordFromEncrypt(plain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(plain))
}

func (a AuthEntity) VerifyPasswordFromPlain(encrypt string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(encrypt), []byte(a.Password))
}

func (a AuthEntity) GenerateToken(secret string) (tokenString string, err error) {
	return utility.GenerateToken(a.PublicId.String(), string(a.Role), secret)
}
