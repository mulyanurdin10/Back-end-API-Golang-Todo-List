package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mysql_user := os.Getenv("MYSQL_USER")
	mysql_password := os.Getenv("MYSQL_PASSWORD")
	mysql_port := os.Getenv("MYSQL_PORT")
	mysql_host := os.Getenv("MYSQL_HOST")
	mysql_dbname := os.Getenv("MYSQL_DBNAME")

	config := map[string]string{
		"MYSQL_USER":     mysql_user,
		"MYSQL_PASSWORD": mysql_password,
		"MYSQL_PORT":     mysql_port,
		"MYSQL_HOST":     mysql_host,
		"MYSQL_DBNAME":   mysql_dbname,
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		config["MYSQL_USER"],
		config["MYSQL_PASSWORD"],
		config["MYSQL_HOST"],
		config["MYSQL_PORT"],
		config["MYSQL_DBNAME"])

	// var e error
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
