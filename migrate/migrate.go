package main

import (
	"github.com/camilamedeir0s/golang_challenge/initializers"
	"github.com/camilamedeir0s/golang_challenge/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	initializers.DB.AutoMigrate(&models.User{})
}