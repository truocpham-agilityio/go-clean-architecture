package main

import (
	"fmt"
	"go-clean-architecture/config"
	"go-clean-architecture/infrastructure/datastore"
	"go-clean-architecture/infrastructure/router"
	"go-clean-architecture/registry"
	"log"

	"github.com/labstack/echo"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server is running at http://localhost:" + config.C.Server.Address)

	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}