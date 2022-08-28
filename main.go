package main

import (
	"testcode/config"
	"testcode/factory"
	"testcode/migration"
	"testcode/routes"
)

func main() {
	dbConn := config.InitDB()
	migration.InitMigrate(dbConn)
	presenter := factory.InitFactory(dbConn)
	e := routes.New(presenter)
	e.Logger.Fatal(e.Start(":3030"))
}
