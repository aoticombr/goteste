package lib

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CustomContext struct {
	*gin.Context
}

func HelperContext(c *gin.Context) *CustomContext {
	return &CustomContext{
		Context: c,
	}
}

// Novo método adicional para CustomContext
func (cc *CustomContext) SetHeader(key string, value string) {
	cc.Request.Header.Del(key)
	cc.Request.Header.Set(key, value)
}

func (cc *CustomContext) GetUser() *Token {
	usuario := cc.GetHeader("usuario")
	if usuario == "" {
		return nil
	}
	var token Token
	err := json.Unmarshal([]byte(usuario), &token)
	if err != nil {
		cc.JSON(http.StatusUnauthorized, gin.H{"error": "falha ao ler o json do usuario"})
		return nil
	}
	return &token
}

func (cc *CustomContext) GetBodyArrayJson(field string) []interface{} {
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

func (cc *CustomContext) GetBodyValueJson(field string) []interface{} {
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

func (cc *CustomContext) ValidaToken(valid bool) bool {
	if !valid {
		return valid
	}
	authorizationHeader := cc.GetHeader("Authorization")
	//fmt.Println("Authorization", authorizationHeader)
	if authorizationHeader == "" {
		cc.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorização inválido"})
		return false
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetSecret_key()), nil
	})
	if err != nil || !token.Valid {
		cc.JSON(http.StatusInternalServerError, gin.H{"error": "Token de autorização inválido"})
		return false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		cc.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter os dados do token"})
		return false
	}
	uID, _ := uuid.Parse(claims["id"].(string))

	jtoken := Token{
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
	jsonData, err := json.Marshal(jtoken)
	if err != nil {
		log.Fatal(err)
	}
	cc.SetHeader("usuario", string(jsonData))
	return true
}

func (cc *CustomContext) ValidaEmpresa(valid bool) bool {
	if !valid {
		return valid
	}
	usuario := cc.GetHeader("usuario")
	if usuario == "" {
		cc.JSON(http.StatusUnauthorized, gin.H{"error": "usuario vazio"})
		return false
	}
	var token Token
	err := json.Unmarshal([]byte(usuario), &token)
	if err != nil {
		cc.JSON(http.StatusUnauthorized, gin.H{"error": "falha ao ler o json do usuario"})
		return false
	}
	if token.ID_EMP == nil {
		cc.JSON(http.StatusUnauthorized, gin.H{"error": "escolha uma empresa"})
		return false
	}
	return true
}

func (cc *CustomContext) GetConvertQry(filt []interface{}) string {
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

func (cc *CustomContext) GetFiltros() string {
	filtros := cc.GetBodyArrayJson("filtros")

	return cc.GetConvertQry(filtros)
}
