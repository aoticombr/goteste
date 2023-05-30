package auth

import (
	apierrors "PARTNER/internal/errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Token struct {
	ID       uuid.UUID  `json:"id"`
	EMAIL    string     `json:"email"`
	ID_EMP   *uuid.UUID `json:"id_emp"`
	ID_GRUPO *uuid.UUID `json:"id_grupo"`
}

func (t Token) GenerateJWT(secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"id":       t.ID,
		"email":    t.EMAIL,
		"id_emp":   t.ID_EMP,
		"id_grupo": t.ID_GRUPO,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(), // Defina o tempo de expiração do JWT, nesse caso, 1 hora.
		"iat":      time.Now().Unix(),                         // Defina o tempo de emissão do JWT.
	}

	JWTtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := JWTtoken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (t Token) ValidateEmp() error {
	if t.ID_EMP == nil {
		return apierrors.NewValidationError("Empresa", "required")
	}

	return nil
}

func (t Token) Check() error {
	return nil
}
