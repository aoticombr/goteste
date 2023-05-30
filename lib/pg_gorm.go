package lib

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConexaoPG_GORM() (*gorm.DB, error) {

	//connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	//	Getpg_ip(), Getpg_port(), Getpg_user(), Getpg_pass(), Getpg_database())
	// "host=localhost port=5432 user=user password=password dbname=database sslmode=disable")
	//if err != nil {

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=UTC",
		Getpg_ip(),
		Getpg_user(),
		Getpg_pass(),
		Getpg_database())

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
