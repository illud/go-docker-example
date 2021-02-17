package users

import (
	"github.com/gin-gonic/gin"
	usersModel "github.com/saturnavt/dockerGo/models/users"
)

func GetUsers(c *gin.Context) {
	var usersmodel []usersModel.User

	usersmodel = append(usersmodel, usersModel.User{Username: "Saturnavt", Password: "Metal is the law"})
	//encodeUsers, err := json.MarshalIndent(usersmodel, "", "  ")
	// if err != nil {
	// 	log.Fatal("Eror: ", err)
	// }

	c.JSON(200, gin.H{
		"message": usersmodel,
	})
}
