package main

import (
	"github.com/gin-gonic/gin"
	usersRoute "github.com/saturnavt/dockerGo/routes/users"
)

func main() {
	r := gin.Default()

	r.GET("/users", usersRoute.GetUsers)

	r.Run(":8080")
}
