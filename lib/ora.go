package lib

import (
	"database/sql"
	"fmt"

	_ "github.com/sijms/go-ora/v2"
)

func GetConexao() (*sql.DB, error) {
	url := fmt.Sprintf("oracle://%s:%s@%s:%d/%s",
		GetOra_user(), GetOra_pass(), GetOra_ip(), GetOra_portInt(), GetOra_schema())

	fmt.Println(url)
	db, err := sql.Open("oracle", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}
