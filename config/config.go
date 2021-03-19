package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	UserName = "saul"
	NameServer = "localhost"
	Password = "123456789"
	DbName = "db_mhsApi"
	Driver = "mysql"
)

func Mysql() (*sql.DB, error){
	UrlDb := fmt.Sprintf("%s:%v@/%s",UserName,Password,DbName)
	db, err := sql.Open(Driver,UrlDb)

	if err != nil{
		log.Fatal(err)
	}
	return db, nil
}