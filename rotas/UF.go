package rotas

import (
	Helpgin "PARTNER/internal/gin"
	"PARTNER/lib"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UF_I struct {
	UF_DNE   string    `json:"uf_dne"`
	UF_NOME  string    `json:"uf_nome"`
	UF_SIGLA string    `json:"uf_sigla"`
	IBGE_NFE string    `json:"ibge_nfe"`
	ID       uuid.UUID `json:"id"`
}
type UF_U struct {
	UF_DNE   string    `json:"uf_dne"`
	UF_NOME  string    `json:"uf_nome"`
	UF_SIGLA string    `json:"uf_sigla"`
	IBGE_NFE string    `json:"ibge_nfe"`
	ID       uuid.UUID `json:"id"`
}

func UF_carga(c *gin.Context) {
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
   a1.uf_sigla as label, 
   a1.id as value, 
   a1.id as key,
   a1.id as id
   from cep_ufs a1`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	lib.API_carga(sql, cc)
}
func UF_lookup(c *gin.Context) {
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
  a1.uf_sigla as name 
  from cep_ufs a1`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	lib.API_lookup(sql, cc)
}
func UF_lista(c *gin.Context) {
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
  FROM cep_ufs A1 
		`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	lib.API_lista(sql, cc)
}
func UF_row_id(c *gin.Context) {
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
                      FROM cep_ufs A1 
                      WHERE A1.id = ` + lib.QuotedStr(cc.Param("id"))
	lib.API_lista(sql, cc)
}
func UF_insert(c *gin.Context) {
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
	var insert UF_I
	if err := cc.ShouldBindJSON(&insert); err != nil {
		fmt.Println(err.Error())
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	Qry, Values := lib.GenerateInsertQuery("UF", insert)
	lib.API_exec(Qry, Values, cc)
}
func UF_update_id(c *gin.Context) {
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
	var update UF_U
	if err := cc.ShouldBindJSON(&update); err != nil {
		fmt.Println(err.Error())
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	Qry, Values := lib.GenerateUpdateQuery("CLUB0001", update, "ID = "+lib.QuotedStr(cc.Param("id")))
	lib.API_exec(Qry, Values, cc)
}
func UF_delete_id(c *gin.Context) {
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
	sql := "delete from CEP_UFS WHERE ID  = $1"
	values := []interface{}{
		cc.Param("id"),
	}
	lib.API_exec(sql, values, cc)
	cc.JSON(200, gin.H{"message": "Success"})
}
