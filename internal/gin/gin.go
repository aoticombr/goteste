package gin

import (
	auth "PARTNER/internal/auth"
	apierrors "PARTNER/internal/errors"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomGin struct {
	*gin.Context
}

func HelperGin(c *gin.Context) *CustomGin {
	return &CustomGin{
		Context: c,
	}
}

func (cc *CustomGin) SetHeader(key string, value string) {
	cc.Request.Header.Del(key)
	cc.Request.Header.Set(key, value)
}

func (cc *CustomGin) GetAutorization() (string, error) {
	authorizationHeader := cc.GetHeader("Authorization")
	//fmt.Println("Authorization", authorizationHeader)
	if authorizationHeader == "" {
		return "", apierrors.NewValidationError("Authorization", "required")
	}
	return authorizationHeader, nil
}

func (cc *CustomGin) GetToken(secretKey string) (*auth.Token, error) {
	authorization, err := cc.GetAutorization()
	if err != nil {
		return nil, err
	}
	tokenString := strings.Replace(authorization, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return nil, apierrors.NewValidationError("Authorization", "invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, apierrors.NewValidationError("Token", "erro read")
	}
	uID, _ := uuid.Parse(claims["id"].(string))

	jtoken := auth.Token{
		ID:    uID,
		EMAIL: claims["email"].(string),
	}
	if idEmp, ok := claims["id_emp"].(string); ok {
		value, _ := uuid.Parse(idEmp)
		jtoken.ID_EMP = &value
	}

	if idGrupo, ok := claims["id_grupo"].(string); ok {
		value, _ := uuid.Parse(idGrupo)
		jtoken.ID_GRUPO = &value
	}
	return &jtoken, nil
}

func (cc *CustomGin) GetBodyArrayJson(field string) []interface{} {
	var requestBody map[string]interface{}

	// BindJSON para mapear o corpo JSON para um mapa
	if err := cc.ShouldBindJSON(&requestBody); err != nil {
		cc.JSON(400, gin.H{"error": "Falha ao analisar o corpo JSON"})
		return nil
	}

	// Acessar o valor do campo "filtros"
	filtrosInterface, ok := requestBody[field].([]interface{})
	if !ok {
		return nil
	}

	values := make([]interface{}, len(filtrosInterface))
	for i, value := range filtrosInterface {
		values[i] = value
	}

	return values
}

func (cc *CustomGin) GetBodyValueJson(field string) []interface{} {
	var requestBody map[string]interface{}

	// BindJSON para mapear o corpo JSON para um mapa
	if err := cc.ShouldBindJSON(&requestBody); err != nil {
		cc.JSON(400, gin.H{"error": "Falha ao analisar o corpo JSON"})
		return nil
	}

	// Acessar o valor do campo "filtros"
	filtrosInterface, ok := requestBody[field].([]interface{})
	if !ok {
		return nil
	}

	values := make([]interface{}, len(filtrosInterface))
	for i, value := range filtrosInterface {
		values[i] = value
	}

	return values
}
func (cc *CustomGin) GetConvertQry(filt []interface{}) string {
	qr := ""

	for _, item := range filt {
		json := item.(map[string]interface{})
		campo := json["campo"].(string)
		tipo := json["tipo"].(string)
		valor := json["valor"].(string)

		if valor != "" {
			switch tipo {
			case "=":
				qr += fmt.Sprintf(" and upper(%s) = upper('%s') ", campo, valor)
			case "==":
				qr += fmt.Sprintf(" and %s = '%s'", campo, valor)
			case "===":
				qr += fmt.Sprintf(" and %s = %s", campo, valor)
			case "=%":
				qr += fmt.Sprintf(" and upper(%s) like upper('%s%%')", campo, valor)
			case "%=":
				qr += fmt.Sprintf(" and upper(%s) like upper('%%%s')", campo, valor)
			case "%=%":
				qr += fmt.Sprintf(" and upper(%s) like upper('%%%s%%') ", campo, valor)
			case ">=":
				qr += fmt.Sprintf(" and %s >= '%s' ", campo, valor)
			case "<=":
				qr += fmt.Sprintf(" and %s <= '%s' ", campo, valor)
			default:
				qr += fmt.Sprintf(" and upper(%s) like upper('%s%%') ", campo, valor)
			}
		}
	}

	return qr
}

func (cc *CustomGin) GetFiltros() string {
	filtros := cc.GetBodyArrayJson("filtros")

	return cc.GetConvertQry(filtros)
}
