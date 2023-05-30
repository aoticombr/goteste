package rotas

import (
	Helpgin "PARTNER/internal/gin"
	"PARTNER/lib"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type EMPRESAS_I struct {
	ID_USER          uuid.UUID `json:"id_user"`
	CPF_CNPJ         string    `json:"cpf_cnpj"`
	RAZAO_NOME       string    `json:"razao_nome"`
	FANTASIA_APELIDO string    `json:"fantasia_apelido"`
	RG_IE            string    `json:"rg_ie"`
	ENDERECO         string    `json:"endereco"`
	NUMERO           string    `json:"numero"`
	COMPLEMENTO      string    `json:"complemento"`
	BAIRRO           string    `json:"bairro"`
	ESTADO           string    `json:"estado"`
	CEP              string    `json:"cep"`
	ID               uuid.UUID `json:"id"`
	ATIVO            string    `json:"ativo"`
}
type EMPRESAS_U struct {
	ID_USER          uuid.UUID `json:"id_user"`
	CPF_CNPJ         string    `json:"cpf_cnpj"`
	RAZAO_NOME       string    `json:"razao_nome"`
	FANTASIA_APELIDO string    `json:"fantasia_apelido"`
	RG_IE            string    `json:"rg_ie"`
	ENDERECO         string    `json:"endereco"`
	NUMERO           string    `json:"numero"`
	COMPLEMENTO      string    `json:"complemento"`
	BAIRRO           string    `json:"bairro"`
	ESTADO           string    `json:"estado"`
	CEP              string    `json:"cep"`
	ID               uuid.UUID `json:"id"`
	ATIVO            string    `json:"ativo"`
}

func EMPRESAS_carga(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := `select 
   a1.razao_nome as label, 
   a1.id as value, 
   a1.id as key,
   a1.id as id
   from emp00001 a1`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	sql += " and id_user = " + lib.QuotedStr(tk.ID.String())
	lib.API_carga(sql, cc)
}
func EMPRESAS_lookup(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := `select 
  a1.id as id,
  a1.razao_nome as name 
  from emp00001 a1`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	sql += " and id_user = " + lib.QuotedStr(tk.ID.String())
	lib.API_lookup(sql, cc)
}
func EMPRESAS_lista(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := `SELECT A1.*                 
  ,B2.email as lk_id_user
  FROM emp00001 A1 
  LEFT JOIN user00001 B2
    ON (B2.id = A1.id_user)
		`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	sql += " and id_user = " + lib.QuotedStr(tk.ID.String())
	lib.API_lista(sql, cc)
}
func EMPRESAS_row_id(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := `SELECT A1.*                 
                      ,A2.email as lk_id_user
                      FROM emp00001 A1 
                      LEFT JOIN user00001 A2
                        ON (A2.id = A1.id_user)
                      WHERE A1.id = ` + lib.QuotedStr(cc.Param("id"))
	sql += " and a1.id_user = " + lib.QuotedStr(tk.ID.String())
	lib.API_lista(sql, cc)
}
func EMPRESAS_insert(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	var insert EMPRESAS_I
	if err := cc.ShouldBindJSON(&insert); err != nil {
		fmt.Println(err.Error())
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	insert.ID_USER = tk.ID
	Qry, Values := lib.GenerateInsertQuery("EMPRESAS", insert)
	lib.API_exec(Qry, Values, cc)
}
func EMPRESAS_update_id(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	var update EMPRESAS_U
	if err := cc.ShouldBindJSON(&update); err != nil {
		fmt.Println(err.Error())
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	update.ID_USER = tk.ID
	Qry, Values := lib.GenerateUpdateQuery("CLUB0001", update, "ID = "+lib.QuotedStr(cc.Param("id")))
	lib.API_exec(Qry, Values, cc)
}
func EMPRESAS_delete_id(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := "delete from EMP00001 WHERE ID  = $1"
	values := []interface{}{
		cc.Param("id"),
	}
	lib.API_exec(sql, values, cc)
	cc.JSON(200, gin.H{"message": "Success"})
}
