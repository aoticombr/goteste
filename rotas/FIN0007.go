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
type FIN0007_I struct {
 ID uuid.UUID `json:"id"`
 ID_DB uuid.UUID `json:"id_db"`
 DT_INC time.Time `json:"dt_inc"`
 DT_ALT time.Time `json:"dt_alt"`
 USER_INC uuid.UUID `json:"user_inc"`
 USER_ALT uuid.UUID `json:"user_alt"`
 ID_FORNECEDOR uuid.UUID `json:"id_fornecedor"`
 ID_CATEGORIA uuid.UUID `json:"id_categoria"`
 DESCRICAO string `json:"descricao"`
 COD_BARRA string `json:"cod_barra"`
 COD_PIX string `json:"cod_pix"`
 DT_VCTO time.Time `json:"dt_vcto"`
 DT_PRRGDO time.Time `json:"dt_prrgdo"`
 VLR_TITULO decimal.Decimal `json:"vlr_titulo"`
 VLR_OUT_MENOS decimal.Decimal `json:"vlr_out_menos"`
 VLR_OUT_MAIS decimal.Decimal `json:"vlr_out_mais"`
 VLR_PAGO decimal.Decimal `json:"vlr_pago"`
 VLR_REST decimal.Decimal `json:"vlr_rest"`
 NUMERO string `json:"numero"`
 RASTREIO string `json:"rastreio"`
 ID_EMP uuid.UUID `json:"id_emp"`
 ID_TIPO_TITULO uuid.UUID `json:"id_tipo_titulo"`
 DT_CONTA time.Time `json:"dt_conta"`
} 
type FIN0007_U struct {
 DT_ALT time.Time `json:"dt_alt"`
 USER_ALT uuid.UUID `json:"user_alt"`
 ID_FORNECEDOR uuid.UUID `json:"id_fornecedor"`
 ID_CATEGORIA uuid.UUID `json:"id_categoria"`
 DESCRICAO string `json:"descricao"`
 COD_BARRA string `json:"cod_barra"`
 COD_PIX string `json:"cod_pix"`
 DT_VCTO time.Time `json:"dt_vcto"`
 DT_PRRGDO time.Time `json:"dt_prrgdo"`
 NUMERO string `json:"numero"`
 ID_TIPO_TITULO uuid.UUID `json:"id_tipo_titulo"`
 DT_CONTA time.Time `json:"dt_conta"`
} 

func FIN0007_carga(c *gin.Context) {           
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
   sql:=`select 
   a1.descricao as label, 
   a1.id as value, 
   a1.id as key,
   a1.id as id
   from fin0007 a1`
  sql += " where 1=1 "                                               
  sql += cc.GetFiltros()                                        
    sql += " and id_db = "+lib.QuotedStr(tk.ID_GRUPO.String())
    sql += " and id_emp = "+lib.QuotedStr(tk.ID_EMP.String())
  lib.API_carga(sql, cc)                                            
}                                                                     
func FIN0007_lookup(c *gin.Context) {           
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
	sql:=`select 
  a1.id as id,
  a1.descricao as name 
  from fin0007 a1`
  sql += " where 1=1 "                                               
  sql += cc.GetFiltros()                                        
    sql += " and id_db = "+lib.QuotedStr(tk.ID_GRUPO.String())
    sql += " and id_emp = "+lib.QuotedStr(tk.ID_EMP.String())
  lib.API_lookup(sql, cc)                                            
}                                                                     
func FIN0007_lista(c *gin.Context) {           
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
	sql:=`SELECT A1.*                 
  ,B2.email as lk_id_db
  ,B5.email as lk_user_inc
  ,B6.email as lk_user_alt
  ,B7.nome as lk_id_fornecedor
  ,B8.descricao as lk_id_categoria
  ,B3.razao_nome as lk_id_emp
  ,B9.descricao as lk_id_tipo_titulo
  FROM fin0007 A1 
  LEFT JOIN user00001 B2
    ON (B2.id = A1.id_db)
  LEFT JOIN user00001 B5
    ON (B5.id = A1.user_inc)
  LEFT JOIN user00001 B6
    ON (B6.id = A1.user_alt)
  LEFT JOIN tb0001 B7
    ON (B7.id = A1.id_fornecedor)
  LEFT JOIN fin0002 B8
    ON (B8.id = A1.id_categoria)
  LEFT JOIN emp00001 B3
    ON (B3.id = A1.id_emp)
  LEFT JOIN fin0004 B9
    ON (B9.id = A1.id_tipo_titulo)
		`
  sql += " where 1=1 "                                               
  sql += cc.GetFiltros()                                        
    sql += " and id_db = "+lib.QuotedStr(tk.ID_GRUPO.String())
    sql += " and id_emp = "+lib.QuotedStr(tk.ID_EMP.String())
  lib.API_lista(sql, cc)                                            
}                                                                     
func FIN0007_row_id(c *gin.Context) {           
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
 sql:= `SELECT A1.*                 
                      ,A2.email as lk_id_db
                      ,A3.email as lk_user_inc
                      ,A4.email as lk_user_alt
                      ,A5.nome as lk_id_fornecedor
                      ,A6.descricao as lk_id_categoria
                      ,A7.razao_nome as lk_id_emp
                      ,A8.descricao as lk_id_tipo_titulo
                      FROM fin0007 A1 
                      LEFT JOIN user00001 A2
                        ON (A2.id = A1.id_db)
                      LEFT JOIN user00001 A3
                        ON (A3.id = A1.user_inc)
                      LEFT JOIN user00001 A4
                        ON (A4.id = A1.user_alt)
                      LEFT JOIN tb0001 A5
                        ON (A5.id = A1.id_fornecedor)
                      LEFT JOIN fin0002 A6
                        ON (A6.id = A1.id_categoria)
                      LEFT JOIN emp00001 A7
                        ON (A7.id = A1.id_emp)
                      LEFT JOIN fin0004 A8
                        ON (A8.id = A1.id_tipo_titulo)
                      WHERE A1.id = `+lib.QuotedStr(cc.Param("id"))
    sql += " and a1.id_db = "+lib.QuotedStr(tk.ID_GRUPO.String())
    sql += " and a1.id_emp = "+lib.QuotedStr(tk.ID_EMP.String())
 lib.API_lista(sql, cc)                            
}                                                                     
func FIN0007_insert(c *gin.Context) {           
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
var insert FIN0007_I                                                                     
if err := cc.ShouldBindJSON(&insert); err != nil {                                        
	fmt.Println(err.Error())                                                                
	cc.JSON(400, gin.H{"error": err.Error()})                                               
	return                                                                                  
	}                                                                                       
 insert.ID_DB = *tk.ID_GRUPO 
 insert.DT_INC = time.Now() 
 insert.DT_ALT = time.Now() 
 insert.USER_INC = tk.ID 
 insert.USER_ALT = tk.ID 
 insert.ID_EMP = *tk.ID_EMP 
Qry, Values := lib.GenerateInsertQuery("FIN0007", insert)                                
lib.API_exec(Qry, Values, cc)                                                             
}                                                                                         
func FIN0007_update_id(c *gin.Context) {           
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
  var update FIN0007_U              
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
func FIN0007_delete_id(c *gin.Context) {           
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
  sql := "delete from FIN0007 WHERE ID  = $1"
  values := []interface{}{                                                       
    cc.Param("id"),                                                                 
  }                                                                             
  lib.API_exec(sql,values, cc)                                                
  cc.JSON(200, gin.H{"message": "Success"})                                   
}                                                                             
