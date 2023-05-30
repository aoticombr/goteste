package lib

import (
	sql "database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

func GetConexaoPG_SQL() (*sql.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Getpg_ip(),
		5432,
		Getpg_user(),
		Getpg_pass(),
		Getpg_database())

	db, err := sql.Open("postgres", connStr)
	SalvarLog(connStr)
	if err != nil {
		SalvarErro("GetConexaoPG:", err)
		return nil, err
	}
	return db, nil
}

func convertValue(value interface{}, columnType *sql.ColumnType) (interface{}, error) {
	// Verificar o tipo da coluna e converter o valor correspondente
	//fmt.Println(value, columnType.DatabaseTypeName())
	//fmt.Println(columnType.DatabaseTypeName())
	switch columnType.DatabaseTypeName() {
	case "UUID":
		uuidBytes, ok := value.([]byte)
		if !ok {
			return nil, fmt.Errorf("valor de coluna UUID inválido: %v", value)
		}
		uuidStr := fmt.Sprintf("%s", uuidBytes)
		parsedUUID, err := uuid.Parse(uuidStr)
		if err != nil {
			return nil, fmt.Errorf("falha ao converter valor de coluna UUID: %v", err)
		}
		return parsedUUID, nil
	case "MONEY":
		moneyBytes, ok := value.([]byte)
		if !ok {
			return nil, fmt.Errorf("valor de coluna MONEY inválido: %v", value)
		}
		// Remover o símbolo do dinheiro ($) e converter os bytes para string
		moneyStr := string(moneyBytes[1:]) // Ignorar o primeiro byte que representa o símbolo do dinheiro
		return moneyStr, nil
	case "NUMERIC":
		numericBytes, ok := value.([]byte)
		if !ok {
			return nil, fmt.Errorf("valor de coluna NUMERIC inválido: %v", value)
		}
		numericStr := string(numericBytes)
		numericValue, err := decimal.NewFromString(numericStr)
		if err != nil {
			return nil, fmt.Errorf("falha ao converter valor de coluna NUMERIC: %v", err)
		}
		return numericValue, nil
	}

	// Se o tipo não for reconhecido, retornar o valor original
	return value, nil
}

func CountRows(r sql.Rows) int {
	count := 0
	for r.Next() {
		count++
	}
	return count
}

func generatePlaceholders(count int) string {
	placeholders := make([]string, count)
	for i := range placeholders {
		placeholders[i] = "?"
	}
	return strings.Join(placeholders, ", ")
}

func GenerateInsertQuery(tableName string, obj interface{}) (string, []interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var columns []string
	var placeholders []string
	var values []interface{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		// Verificar se o campo é do tipo time.Time
		if field.Type == reflect.TypeOf(time.Time{}) {
			// Converter para time.Time
			timeValue := value.(time.Time)
			// Formatar a data no formato esperado pelo PostgreSQL
			formattedTime := timeValue.Format("2006-01-02")
			// Adicionar o valor formatado à lista de valores
			values = append(values, formattedTime)
		} else {
			values = append(values, value)
		}

		columns = append(columns, field.Name)
		placeholders = append(placeholders, fmt.Sprintf("$%d", len(placeholders)+1))
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		tableName, strings.Join(columns, ", "), strings.Join(placeholders, ", "))

	return query, values
}

func GenerateUpdateQuery(tableName string, obj interface{}, condition string) (string, []interface{}) {
	fmt.Println("GenerateUpdateQuery...1")
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	fmt.Println("GenerateUpdateQuery...2")
	//var assignments []string
	var placeholders []string
	var values []interface{}
	fmt.Println("GenerateUpdateQuery...3")
	fmt.Println(obj)
	fmt.Println("GenerateUpdateQuery...3")

	fmt.Println("GenerateUpdateQuery...3")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		//	assignments = append(assignments, fmt.Sprintf("%s = ?", field.Name))
		if field.Type == reflect.TypeOf(time.Time{}) {
			// Converter para time.Time
			timeValue := value.(time.Time)
			// Formatar a data no formato esperado pelo PostgreSQL
			formattedTime := timeValue.Format("2006-01-02")
			// Adicionar o valor formatado à lista de valores
			values = append(values, formattedTime)
		} else {
			values = append(values, value)
		}
		placeholders = append(placeholders, fmt.Sprintf("%s = $%d", field.Name, len(placeholders)+1))
	}
	fmt.Println("GenerateUpdateQuery...4")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, strings.Join(placeholders, ", "), condition)

	return query, values
}
