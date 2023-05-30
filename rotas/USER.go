package rotas

import (
	Helpgin "PARTNER/internal/gin"
	"PARTNER/lib"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type USER_I struct {
	ID     uuid.UUID `json:"id"`
	EMAIL  string    `json:"email"`
	SENHA  string    `json:"senha"`
	STATUS string    `json:"status"`
}
type USER_U struct {
	ID     uuid.UUID `json:"id"`
	EMAIL  string    `json:"email"`
	SENHA  string    `json:"senha"`
	STATUS string    `json:"status"`
}

func USER_carga(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := `select 
   a1.email as label, 
   a1.id as value, 
   a1.id as key,
   a1.id as id
   from user00001 a1`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	lib.API_carga(sql, cc)
}
func USER_lookup(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := `select 
  a1.id as id,
  a1.email as name 
  from user00001 a1`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	lib.API_lookup(sql, cc)
}
func USER_lista(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := `SELECT A1.*                 
  FROM user00001 A1 
		`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	lib.API_lista(sql, cc)
}
func USER_row_id(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := `SELECT A1.*                 
                      FROM user00001 A1 
                      WHERE A1.id = ` + lib.QuotedStr(cc.Param("id"))
	lib.API_lista(sql, cc)
}
func USER_insert(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	var insert USER_I
	if err := cc.ShouldBindJSON(&insert); err != nil {
		fmt.Println(err.Error())
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	Qry, Values := lib.GenerateInsertQuery("USER", insert)
	lib.API_exec(Qry, Values, cc)
}
func USER_update_id(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	var update USER_U
	if err := cc.ShouldBindJSON(&update); err != nil {
		fmt.Println(err.Error())
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	Qry, Values := lib.GenerateUpdateQuery("CLUB0001", update, "ID = "+lib.QuotedStr(cc.Param("id")))
	lib.API_exec(Qry, Values, cc)
}
func USER_delete_id(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := "delete from USER00001 WHERE ID  = $1"
	values := []interface{}{
		cc.Param("id"),
	}
	lib.API_exec(sql, values, cc)
	cc.JSON(200, gin.H{"message": "Success"})
}
