// https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m
package main

import (
	"gin-mongo-api/config"
	"gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDB()

	routes.UserRoute(router)

	router.Run("localhost:8000")
}
