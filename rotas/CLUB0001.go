package rotas

import (
	Helpgin "PARTNER/internal/gin"
	"PARTNER/lib"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CLUB0001_I struct {
	ID        uuid.UUID       `json:"id"`
	DT_INC    time.Time       `json:"dt_inc"`
	ID_DB     uuid.UUID       `json:"id_db"`
	DT_ALT    time.Time       `json:"dt_alt"`
	USER_ALT  uuid.UUID       `json:"user_alt"`
	DESCRICAO string          `json:"descricao"`
	VALOR     decimal.Decimal `json:"valor"`
	ATIVO     string          `json:"ativo"`
	USER_INC  uuid.UUID       `json:"user_inc"`
}
type CLUB0001_U struct {
	DT_ALT    time.Time       `json:"dt_alt"`
	USER_ALT  uuid.UUID       `json:"user_alt"`
	DESCRICAO string          `json:"descricao"`
	VALOR     decimal.Decimal `json:"valor"`
	ATIVO     string          `json:"ativo"`
}

func CLUB0001_carga(c *gin.Context) {
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
   a1.descricao as label, 
   a1.id as value, 
   a1.id as key,
   a1.id as id
   from club0001 a1`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	sql += " and id_db = " + lib.QuotedStr(tk.ID_GRUPO.String())
	lib.API_carga(sql, cc)
}
func CLUB0001_lookup(c *gin.Context) {
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
  a1.descricao as name 
  from club0001 a1`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	sql += " and id_db = " + lib.QuotedStr(tk.ID_GRUPO.String())
	lib.API_lookup(sql, cc)
}
func CLUB0001_lista(c *gin.Context) {
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
  ,B6.email as lk_user_alt
  ,B5.email as lk_user_inc
  FROM club0001 A1 
  LEFT JOIN user00001 B2
    ON (B2.id = A1.id_db)
  LEFT JOIN user00001 B6
    ON (B6.id = A1.user_alt)
  LEFT JOIN user00001 B5
    ON (B5.id = A1.user_inc)
		`
	sql += " where 1=1 "
	sql += cc.GetFiltros()
	sql += " and id_db = " + lib.QuotedStr(tk.ID_GRUPO.String())
	lib.API_lista(sql, cc)
}
func CLUB0001_row_id(c *gin.Context) {
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
                      ,A3.email as lk_user_alt
                      ,A4.email as lk_user_inc
                      FROM club0001 A1 
                      LEFT JOIN user00001 A2
                        ON (A2.id = A1.id_db)
                      LEFT JOIN user00001 A3
                        ON (A3.id = A1.user_alt)
                      LEFT JOIN user00001 A4
                        ON (A4.id = A1.user_inc)
                      WHERE A1.id = ` + lib.QuotedStr(cc.Param("id"))
	sql += " and a1.id_db = " + lib.QuotedStr(tk.ID_GRUPO.String())
	lib.API_lista(sql, cc)
}
func CLUB0001_insert(c *gin.Context) {
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
	var insert CLUB0001_I
	if err := cc.ShouldBindJSON(&insert); err != nil {
		fmt.Println(err.Error())
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	insert.DT_INC = time.Now()
	insert.ID_DB = *tk.ID_GRUPO
	insert.DT_ALT = time.Now()
	insert.USER_ALT = tk.ID
	insert.USER_INC = tk.ID
	Qry, Values := lib.GenerateInsertQuery("CLUB0001", insert)
	lib.API_exec(Qry, Values, cc)
}
func CLUB0001_update_id(c *gin.Context) {
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
	var update CLUB0001_U
	if err := cc.ShouldBindJSON(&update); err != nil {
		fmt.Println(err.Error())
		cc.JSON(400, gin.H{"error": err.Error()})
		return
	}
	update.DT_ALT = time.Now()
	update.USER_ALT = tk.ID
	Qry, Values := lib.GenerateUpdateQuery("CLUB0001", update, "ID = "+lib.QuotedStr(cc.Param("id")))
	lib.API_exec(Qry, Values, cc)
}
func CLUB0001_delete_id(c *gin.Context) {
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
	sql := "delete from CLUB0001 WHERE ID  = $1"
	values := []interface{}{
		cc.Param("id"),
	}
	lib.API_exec(sql, values, cc)
	cc.JSON(200, gin.H{"message": "Success"})
}
