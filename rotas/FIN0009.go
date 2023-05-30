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
type FIN0009_I struct {
 ID_REF string `json:"id_ref"`
 DT_PGTO time.Time `json:"dt_pgto"`
 ID_FORM_PGTO uuid.UUID `json:"id_form_pgto"`
 ID uuid.UUID `json:"id"`
 ID_DB uuid.UUID `json:"id_db"`
 DT_INC time.Time `json:"dt_inc"`
 USER_INC uuid.UUID `json:"user_inc"`
 ID_TITULO uuid.UUID `json:"id_titulo"`
 VLR_JUROS decimal.Decimal `json:"vlr_juros"`
 VLR_DESC decimal.Decimal `json:"vlr_desc"`
 VLR_PAGO decimal.Decimal `json:"vlr_pago"`
} 
type FIN0009_U struct {
 ID_REF string `json:"id_ref"`
 DT_PGTO time.Time `json:"dt_pgto"`
 ID_FORM_PGTO uuid.UUID `json:"id_form_pgto"`
 ID uuid.UUID `json:"id"`
 ID_DB uuid.UUID `json:"id_db"`
 DT_INC time.Time `json:"dt_inc"`
 USER_INC uuid.UUID `json:"user_inc"`
 ID_TITULO uuid.UUID `json:"id_titulo"`
 VLR_JUROS decimal.Decimal `json:"vlr_juros"`
 VLR_DESC decimal.Decimal `json:"vlr_desc"`
 VLR_PAGO decimal.Decimal `json:"vlr_pago"`
} 

func FIN0009_carga(c *gin.Context) {           
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
   a1.id as label, 
   a1.id as value, 
   a1.id as key,
   a1.id as id
   from fin0009 a1`
  sql += " where 1=1 "                                               
  sql += cc.GetFiltros()                                        
  lib.API_carga(sql, cc)                                            
}                                                                     
func FIN0009_lookup(c *gin.Context) {           
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
  a1.id as name 
  from fin0009 a1`
  sql += " where 1=1 "                                               
  sql += cc.GetFiltros()                                        
  lib.API_lookup(sql, cc)                                            
}                                                                     
func FIN0009_lista(c *gin.Context) {           
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
  ,B11.descricao as lk_id_form_pgto
  ,B2.email as lk_id_db
  ,B4.email as lk_user_inc
  ,B5.descricao as lk_id_titulo
  FROM fin0009 A1 
  LEFT JOIN fin0003 B11
    ON (B11.id = A1.id_form_pgto)
  LEFT JOIN user00001 B2
    ON (B2.id = A1.id_db)
  LEFT JOIN user00001 B4
    ON (B4.id = A1.user_inc)
  LEFT JOIN fin0007 B5
    ON (B5.id = A1.id_titulo)
		`
  sql += " where 1=1 "                                               
  sql += cc.GetFiltros()                                        
  lib.API_lista(sql, cc)                                            
}                                                                     
func FIN0009_row_id(c *gin.Context) {           
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
                      ,A2.descricao as lk_id_form_pgto
                      ,A3.email as lk_id_db
                      ,A4.email as lk_user_inc
                      ,A5.descricao as lk_id_titulo
                      FROM fin0009 A1 
                      LEFT JOIN fin0003 A2
                        ON (A2.id = A1.id_form_pgto)
                      LEFT JOIN user00001 A3
                        ON (A3.id = A1.id_db)
                      LEFT JOIN user00001 A4
                        ON (A4.id = A1.user_inc)
                      LEFT JOIN fin0007 A5
                        ON (A5.id = A1.id_titulo)
                      WHERE A1.id = `+lib.QuotedStr(cc.Param("id"))
 lib.API_lista(sql, cc)                            
}                                                                     
func FIN0009_insert(c *gin.Context) {           
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
var insert FIN0009_I                                                                     
if err := cc.ShouldBindJSON(&insert); err != nil {                                        
	fmt.Println(err.Error())                                                                
	cc.JSON(400, gin.H{"error": err.Error()})                                               
	return                                                                                  
	}                                                                                       
Qry, Values := lib.GenerateInsertQuery("FIN0009", insert)                                
lib.API_exec(Qry, Values, cc)                                                             
}                                                                                         
func FIN0009_update_id(c *gin.Context) {           
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
  var update FIN0009_U              
  if err := cc.ShouldBindJSON(&update); err != nil {                       
    fmt.Println(err.Error())                                               
    cc.JSON(400, gin.H{"error": err.Error()})                              
  return                                                                 
  }                                                                        
 Qry, Values := lib.GenerateUpdateQuery("CLUB0001", update, "ID = "+lib.QuotedStr(cc.Param("id")))  
 lib.API_exec(Qry, Values, cc)                                         
}                                                                     
func FIN0009_delete_id(c *gin.Context) {           
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
  sql := "delete from FIN0009 WHERE ID  = $1"
  values := []interface{}{                                                       
    cc.Param("id"),                                                                 
  }                                                                             
  lib.API_exec(sql,values, cc)                                                
  cc.JSON(200, gin.H{"message": "Success"})                                   
}                                                                             
