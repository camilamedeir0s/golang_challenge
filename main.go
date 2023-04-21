package main

import (
	"github.com/gin-gonic/gin"
	"github.com/camilamedeir0s/golang_challenge/initializers"
	"github.com/camilamedeir0s/golang_challenge/controllers"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	r := gin.Default()
	r.POST("/user", controllers.UsersCreate)
	r.POST("/user/createWithArray", controllers.UsersArrayCreate)
	r.GET("/users", controllers.UsersIndex)
	r.GET("/user/:username", controllers.GetUser)
	r.PUT("/user/:username", controllers.UserUpdate)
	r.DELETE("/user/:username", controllers.UserDelete)

	r.Run()
}