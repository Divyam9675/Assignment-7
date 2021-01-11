package main

import (
	"log"

	routes "Assignment-7/routes"
	configure "Assignment-7/configure"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect DB
	configure.Connect()

	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}
