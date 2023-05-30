package rotas

import (
	Helpgin "PARTNER/internal/gin"
	"PARTNER/lib"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type USEREMP_I struct {
	ID      uuid.UUID `json:"id"`
	ID_EMP  uuid.UUID `json:"id_emp"`
	ID_USER uuid.UUID `json:"id_user"`
}
type USEREMP_U struct {
	ID      uuid.UUID `json:"id"`
	ID_EMP  uuid.UUID `json:"id_emp"`
	ID_USER uuid.UUID `json:"id_user"`
}

func USEREMP_carga(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := `select 
   a1.id as label, 
   a1.id as value, 
   a1.id as key,
   a1.id as id
   from emp00003 a1`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	lib.API_carga(sql, cc)
}
func USEREMP_lookup(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := `select 
  a1.id as id,
  a1.id as name 
  from emp00003 a1`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	lib.API_lookup(sql, cc)
}
func USEREMP_lista(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := `SELECT A1.*                 
  ,B2.razao_nome as lk_id_emp
  ,B3.email as lk_id_user
  FROM emp00003 A1 
  LEFT JOIN emp00001 B2
    ON (B2.id = A1.id_emp)
  LEFT JOIN user00001 B3
    ON (B3.id = A1.id_user)
		`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	lib.API_lista(sql, cc)
}
func USEREMP_row_id(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := `SELECT A1.*                 
                      ,A2.razao_nome as lk_id_emp
                      ,A3.email as lk_id_user
                      FROM emp00003 A1 
                      LEFT JOIN emp00001 A2
                        ON (A2.id = A1.id_emp)
                      LEFT JOIN user00001 A3
                        ON (A3.id = A1.id_user)
                      WHERE A1.id = ` + lib.QuotedStr(cc.Param("id"))
	lib.API_lista(sql, cc)
}
func USEREMP_insert(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	var insert USEREMP_I
	if err := cc.ShouldBindJSON(&insert); err != nil {
		fmt.Println(err.Error())
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	Qry, Values := lib.GenerateInsertQuery("USEREMP", insert)
	lib.API_exec(Qry, Values, cc)
}
func USEREMP_update_id(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	var update USEREMP_U
	if err := cc.ShouldBindJSON(&update); err != nil {
		fmt.Println(err.Error())
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	Qry, Values := lib.GenerateUpdateQuery("CLUB0001", update, "ID = "+lib.QuotedStr(cc.Param("id")))
	lib.API_exec(Qry, Values, cc)
}
func USEREMP_delete_id(c *gin.Context) {
	cc := Helpgin.HelperGin(c)
	tk, err := cc.GetToken(lib.GetSecret_key())
	if err != nil {
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tk.Check()
	sql := "delete from EMP00003 WHERE ID  = $1"
	values := []interface{}{
		cc.Param("id"),
	}
	lib.API_exec(sql, values, cc)
	cc.JSON(200, gin.H{"message": "Success"})
}
