package rotas

import (
	Helpgin "PARTNER/internal/gin"
	"PARTNER/lib"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TB0001_I struct {
	DT_ALT      time.Time `json:"dt_alt"`
	ID          uuid.UUID `json:"id"`
	ID_DB       uuid.UUID `json:"id_db"`
	DT_INC      time.Time `json:"dt_inc"`
	USER_INC    uuid.UUID `json:"user_inc"`
	USER_ALT    uuid.UUID `json:"user_alt"`
	NOME        string    `json:"nome"`
	FANTASIA    string    `json:"fantasia"`
	END_LOGR    string    `json:"end_logr"`
	END_NRO     string    `json:"end_nro"`
	END_CIDADE  string    `json:"end_cidade"`
	END_UF      uuid.UUID `json:"end_uf"`
	END_CEP     string    `json:"end_cep"`
	CNPJ_CPF    string    `json:"cnpj_cpf"`
	RG_IE       string    `json:"rg_ie"`
	RG_ORGAO    string    `json:"rg_orgao"`
	RG_UF       uuid.UUID `json:"rg_uf"`
	IE          string    `json:"ie"`
	IM          string    `json:"im"`
	RNTC        string    `json:"rntc"`
	CRC         string    `json:"crc"`
	EMAIL       string    `json:"email"`
	TELEFONE    string    `json:"telefone"`
	CELULAR     string    `json:"celular"`
	TIPO_PESSOA string    `json:"tipo_pessoa"`
	DTA_NASC    time.Time `json:"dta_nasc"`
	SEXO        string    `json:"sexo"`
	END_BAIRRO  string    `json:"end_bairro"`
}
type TB0001_U struct {
	DT_ALT      time.Time `json:"dt_alt"`
	ID          uuid.UUID `json:"id"`
	ID_DB       uuid.UUID `json:"id_db"`
	USER_ALT    uuid.UUID `json:"user_alt"`
	NOME        string    `json:"nome"`
	FANTASIA    string    `json:"fantasia"`
	END_LOGR    string    `json:"end_logr"`
	END_NRO     string    `json:"end_nro"`
	END_CIDADE  string    `json:"end_cidade"`
	END_UF      uuid.UUID `json:"end_uf"`
	END_CEP     string    `json:"end_cep"`
	CNPJ_CPF    string    `json:"cnpj_cpf"`
	RG_IE       string    `json:"rg_ie"`
	RG_ORGAO    string    `json:"rg_orgao"`
	RG_UF       uuid.UUID `json:"rg_uf"`
	IE          string    `json:"ie"`
	IM          string    `json:"im"`
	RNTC        string    `json:"rntc"`
	CRC         string    `json:"crc"`
	EMAIL       string    `json:"email"`
	TELEFONE    string    `json:"telefone"`
	CELULAR     string    `json:"celular"`
	TIPO_PESSOA string    `json:"tipo_pessoa"`
	DTA_NASC    time.Time `json:"dta_nasc"`
	SEXO        string    `json:"sexo"`
	END_BAIRRO  string    `json:"end_bairro"`
}

func TB0001_carga(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	err = tk.ValidateEmp()
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sql := `select 
   a1.nome as label, 
   a1.id as value, 
   a1.id as key,
   a1.id as id
   from tb0001 a1`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	sql += " and id_db = " + lib.QuotedStr(tk.ID_GRUPO.String())
	lib.API_carga(sql, cc)
}
func TB0001_lookup(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	err = tk.ValidateEmp()
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sql := `select 
  a1.id as id,
  a1.nome as name 
  from tb0001 a1`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	sql += " and id_db = " + lib.QuotedStr(tk.ID_GRUPO.String())
	lib.API_lookup(sql, cc)
}
func TB0001_lista(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	err = tk.ValidateEmp()
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sql := `SELECT A1.*                 
  ,B2.email as lk_id_db
  ,B5.email as lk_user_inc
  ,B6.email as lk_user_alt
  ,B29.uf_sigla as lk_end_uf
  ,B20.uf_sigla as lk_rg_uf
  FROM tb0001 A1 
  LEFT JOIN user00001 B2
    ON (B2.id = A1.id_db)
  LEFT JOIN user00001 B5
    ON (B5.id = A1.user_inc)
  LEFT JOIN user00001 B6
    ON (B6.id = A1.user_alt)
  LEFT JOIN cep_ufs B29
    ON (B29.id = A1.end_uf)
  LEFT JOIN cep_ufs B20
    ON (B20.id = A1.rg_uf)
		`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	sql += " and id_db = " + lib.QuotedStr(tk.ID_GRUPO.String())
	lib.API_lista(sql, cc)
}
func TB0001_row_id(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	err = tk.ValidateEmp()
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sql := `SELECT A1.*                 
                      ,A2.email as lk_id_db
                      ,A3.email as lk_user_inc
                      ,A4.email as lk_user_alt
                      ,A5.uf_sigla as lk_end_uf
                      ,A6.uf_sigla as lk_rg_uf
                      FROM tb0001 A1 
                      LEFT JOIN user00001 A2
                        ON (A2.id = A1.id_db)
                      LEFT JOIN user00001 A3
                        ON (A3.id = A1.user_inc)
                      LEFT JOIN user00001 A4
                        ON (A4.id = A1.user_alt)
                      LEFT JOIN cep_ufs A5
                        ON (A5.id = A1.end_uf)
                      LEFT JOIN cep_ufs A6
                        ON (A6.id = A1.rg_uf)
                      WHERE A1.id = ` + lib.QuotedStr(cc.Param("id"))
	sql += " and a1.id_db = " + lib.QuotedStr(tk.ID_GRUPO.String())
	lib.API_lista(sql, cc)
}
func TB0001_insert(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	err = tk.ValidateEmp()
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var insert TB0001_I
	if err := cc.ShouldBindJSON(&insert); err != nil {
		fmt.Println(err.Error())
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	insert.DT_ALT = time.Now()
	insert.ID_DB = *tk.ID_GRUPO
	insert.DT_INC = time.Now()
	insert.USER_INC = tk.ID
	insert.USER_ALT = tk.ID
	Qry, Values := lib.GenerateInsertQuery("TB0001", insert)
	lib.API_exec(Qry, Values, cc)
}
func TB0001_update_id(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	err = tk.ValidateEmp()
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var update TB0001_U
	if err := cc.ShouldBindJSON(&update); err != nil {
		fmt.Println(err.Error())
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	update.DT_ALT = time.Now()
	update.ID_DB = *tk.ID_GRUPO
	update.USER_ALT = tk.ID
	Qry, Values := lib.GenerateUpdateQuery("CLUB0001", update, "ID = "+lib.QuotedStr(cc.Param("id")))
	lib.API_exec(Qry, Values, cc)
}
func TB0001_delete_id(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	err = tk.ValidateEmp()
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	sql := "delete from TB0001 WHERE ID  = $1"
	values := []interface{}{
		cc.Param("id"),
	}
	lib.API_exec(sql, values, cc)
	cc.JSON(200, gin.H{"message": "Success"})
}
