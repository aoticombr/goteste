package rotas

import (
	auth "PARTNER/internal/auth"
	Helpgin "PARTNER/internal/gin"
	"PARTNER/lib"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestBody struct {
	IdEmp   string `json:"id_emp"`
	IdGrupo string `json:"id_grupo"`
}

type Auth struct {
	id      uuid.UUID
	email   string
	status  bool
	id_emp  *uuid.UUID
	id_user *uuid.UUID
}

type Validade struct {
	auth          bool    `json:"auth"`
	emp           bool    `json:"emp"`
	token         string  `json:"token"`
	refresh_token *string `json:"refresh_token"`
	expiresIn     *string `json:"expiresIn"`
}

func PostSignIn(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	user := &auth.User{}
	if err := cc.ShouldBindJSON(&user); err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := user.Validate()
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sql :=
		`select a1.id, a1.email, a1.status, a1.id_emp, a2.id_user as id_grupo 
 	 from USER00001 a1
	 left join emp00001 a2
	   on (a2.id = a1.id_emp)
	 where md5(a1.EMAIL) = ` + lib.QuotedStr(lib.StrToMd5(user.Username))
	sql += ` and a1.SENHA = ` + lib.QuotedStr(lib.StrToMd5(user.Password))
	db, err := lib.GetConexaoPG_SQL()
	if err != nil {
		cc.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})

		return
	}
	defer db.Close()
	rows, err := db.Query(sql)
	if err != nil {
		cc.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()
	results := []Auth{}

	// Iterar sobre as linhas retornadas
	for rows.Next() {
		var result Auth
		err := rows.Scan(&result.id, &result.email, &result.status, &result.id_emp, &result.id_emp)
		if err != nil {
			log.Println(err)
			cc.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		results = append(results, result)
	}
	if len(results) > 0 {
		if results[0].status == false {
			cc.JSON(http.StatusUnauthorized, gin.H{"error": "Usu?rio Bloqueado, fale com seu administrador."})
			return
		}
		jtoken := auth.Token{
			ID:       results[0].id,
			EMAIL:    results[0].email,
			ID_EMP:   results[0].id_emp,
			ID_GRUPO: results[0].id_user,
		}
		jwtToken, err := jtoken.GenerateJWT(lib.GetSecret_key())
		if err != nil {
			cc.JSON(http.StatusUnauthorized, gin.H{"error": "Erro ao gerar o JWT"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"auth": true, "emp": false, "token": jwtToken, "refresh_token": nil, "expiresIn": "7d"})
		return
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "N?o encontrado Usu?rio/Senha invalido."})
		return
	}
	//	if (rows.)
}
func PostSignEmp(c *gin.Context) {
	cc := Helpgin.HelperGin(c)

	fmt.Println("login...")
	var eg auth.GrupoEmp
	if err := c.ShouldBindJSON(&eg); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := eg.Validate()
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	usuario := cc.GetHeader("usuario")
	if usuario == "" {
		cc.JSON(http.StatusUnauthorized, gin.H{"error": "usuario vazio"})
		return
	}
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.ID_EMP = eg.ID_EMP
	tk.ID_GRUPO = eg.ID_GRUPO
	jwtToken, err := tk.GenerateJWT(lib.GetSecret_key())
	if err != nil {
		cc.JSON(http.StatusUnauthorized, gin.H{"error": "Erro ao gerar o JWT"})
		return
	} else {

		cc.JSON(http.StatusOK, gin.H{"auth": true, "emp": true, "token": jwtToken, "refresh_token": nil, "expiresIn": "7d"})
		return
	}

}
func GetSignEmp(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = tk.ValidateEmp()
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}

	sql := `select 
                a1.id as id,
                a1.id as value,                
                a1.id as key,
                a1.id_emp as id_emp,
                a2.id_user as id_grupo,
                concat('(',a2.cpf_cnpj,') ', a2.razao_nome) as label
              from emp00003 a1
              inner join emp00001 a2
                on (a2.id = a1.id_emp)
              where a1.id_user = ` + lib.QuotedStr(tk.ID.String())
	lib.API_lista(sql, cc)

}
