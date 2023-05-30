package lib

import (
	HelpGin "PARTNER/internal/gin"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func API_carga(sql string, c *HelpGin.CustomGin) {
	c.Header("Content-Type", "application/json")
	db, err := GetConexaoPG_SQL()
	if err != nil {
		SalvarErro("erro metodo xx carga:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer db.Close()
	rows, err := db.Query(sql)
	if err != nil {
		SalvarErro("erro metodo xx carga:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()
	// Criar uma slice para armazenar os resultados
	// Criar uma slice para armazenar os resultados
	results := []Carga{}

	// Iterar sobre as linhas retornadas
	for rows.Next() {
		var result Carga
		err := rows.Scan(&result.LABEL, &result.VALUE, &result.KEY, &result.ID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		results = append(results, result)
	}

	// Verificar erros de iteração
	err = rows.Err()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Serializar o resultado para JSON
	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Escrever o resultado na resposta
	c.Data(http.StatusOK, "application/json", jsonData)
}

func API_lookup(sql string, c *HelpGin.CustomGin) {
	c.Header("Content-Type", "application/json")
	db, err := GetConexaoPG_SQL()
	if err != nil {
		SalvarErro("erro metodo xx carga:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer db.Close()
	rows, err := db.Query(sql)
	if err != nil {
		SalvarErro("erro metodo xx carga:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()
	// Criar uma slice para armazenar os resultados
	// Criar uma slice para armazenar os resultados
	results := []Lookup{}

	// Iterar sobre as linhas retornadas
	for rows.Next() {
		var result Lookup
		err := rows.Scan(&result.ID, &result.NAME)
		if err != nil {
			//	log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		results = append(results, result)
	}

	// Verificar erros de iteração
	err = rows.Err()
	if err != nil {
		//	log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Serializar o resultado para JSON
	jsonData, err := json.Marshal(results)
	if err != nil {
		//log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Escrever o resultado na resposta
	c.Data(http.StatusOK, "application/json", jsonData)
}

func API_lista(sql string, c *HelpGin.CustomGin) {
	c.Header("Content-Type", "application/json")
	db, err := GetConexaoPG_SQL()
	if err != nil {
		SalvarErro("erro metodo xx carga:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})

		return
	}
	defer db.Close()
	rows, err := db.Query(sql)
	if err != nil {
		SalvarErro("erro metodo xx carga:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Criar uma slice para armazenar os resultados
	results := make([]map[string]interface{}, 0)

	// Iterar sobre as linhas retornadas
	for rows.Next() {
		columnPointers := make([]interface{}, len(columns))
		columnValues := make([]interface{}, len(columns))
		for i := range columnValues {
			columnPointers[i] = &columnValues[i]
		}

		err := rows.Scan(columnPointers...)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		rowData := make(map[string]interface{})
		for i, column := range columns {
			columnType := columnTypes[i]
			value := columnValues[i]

			// Converter o valor para o tipo apropriado
			convertedValue, err := convertValue(value, columnType)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
				return
			}

			rowData[column] = convertedValue
		}
		results = append(results, rowData)
	}

	// Verificar erros de iteração
	err = rows.Err()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Serializar o resultado para JSON
	jsonData, err := json.Marshal(results)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	// Escrever o resultado na resposta
	c.Data(http.StatusOK, "application/json", jsonData)
}
func API_exec(sql string, values []interface{}, c *HelpGin.CustomGin) {
	c.Header("Content-Type", "application/json")
	db, err := GetConexaoPG_SQL()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao conectar"})
		return
	}
	defer db.Close()
	stmt, err := db.Prepare(sql)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao preparar"})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(values...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao executar"})
	}
	c.JSON(200, gin.H{"message": "Success"})

}
