package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//ToDo: This can be part of env variables.
const (
	DBUSER     = "root"
	DBPASSWORD = "root"
	DBNAME     = "city"
)

//NewMysqlConnection for client connection
func NewDbConnection() *gorm.DB {

	strConnection := DBUSER + ":" + DBPASSWORD + "@/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"

	//Open Connection for GORM instance
	client, err := gorm.Open("mysql", strConnection)

	if err != nil {
		fmt.Println("----", err)
		panic("Error! Can't create connection to database")

	}
	return client

}
