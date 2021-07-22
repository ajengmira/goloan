package models

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *gorm.DB

func GetConn() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbName := os.Getenv("db_name")
	dbPassword := os.Getenv("db_pass")
	dbUsername := os.Getenv("db_user")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")

	/*
		conString := fmt.Sprintf("%s%s%s%s%s%s%s%s%s%s", dbUsername,
			":", dbPassword, "@tcp(", dbHost, ":", dbPort, ")/", dbName, "?parseTime=true")

		fmt.Println(conString)
		var db, errr = gorm.Open("mysql", conString)
	*/

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, errr := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errr != nil {
		panic(errr)
	}

	return db
}
