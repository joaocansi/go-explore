package main

import (
	"api/src/internal/users"
	"api/src/pkg"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := pkg.NewDatabaseConnection()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	users.CreateRoutes(router, db)

	router.Run()
}
