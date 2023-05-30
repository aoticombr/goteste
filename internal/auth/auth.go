package auth

import (
	apierrors "PARTNER/internal/errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	Username string `json:"user"`
	Password string `json:"pass"`
}

func (u User) Validate() error {
	if u.Username == "" {
		return apierrors.NewValidationError("User", "required")
	}
	if u.Password == "" {
		return apierrors.NewValidationError("Pass", "required")
	}
	return nil
}

func (u User) ValidateUser(c *gin.Context) {

}

type GrupoEmp struct {
	ID_EMP   *uuid.UUID `json:"id_emp"`
	ID_GRUPO *uuid.UUID `json:"id_grupo"`
}

func (ge GrupoEmp) Validate() error {
	if ge.ID_EMP == nil {
		return apierrors.NewValidationError("id_emp", "required")
	}
	if ge.ID_GRUPO == nil {
		return apierrors.NewValidationError("id_grupo", "required")
	}
	return nil
}
