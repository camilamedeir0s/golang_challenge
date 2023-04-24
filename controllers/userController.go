package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/camilamedeir0s/golang_challenge/initializers"
	"github.com/camilamedeir0s/golang_challenge/models"
	"reflect"
)

func UsersCreate(c *gin.Context) {
	length := c.Request.ContentLength
    if length == 0 {
        c.JSON(404, gin.H{"error": "Request não possui body"})
        return
    }

	var user models.User

	err := c.BindJSON(&user)
	if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"usuario criado": user,
	})
}

func UsersArrayCreate(c *gin.Context) {
	length := c.Request.ContentLength
    if length == 0 {
        c.JSON(404, gin.H{"error": "Request não possui body"})
        return
    }
	
	var users []models.User
	err := c.BindJSON(&users)
    
	if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	result := initializers.DB.Create(&users)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"Todos os usuários criados": users,
	})
}

func UsersIndex(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"usuarios": users,
	})
}

func GetUser(c *gin.Context){
	username := c.Param("username")
	var user models.User
	result := initializers.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		c.JSON(404, gin.H{
			"Not found": "Usuário não encontrado"})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UserUpdate(c *gin.Context){
	username := c.Param("username")

	var user models.User
	result := initializers.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"Not found": "Usuário não encontrado"})
		return
	}

	err := c.ShouldBindJSON(&user)
	if err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

	initializers.DB.Model(*&user).UpdateColumns(user)
	c.JSON(200, gin.H{
		"message": "usuário atualizado",
	})
}

func UserDelete(c *gin.Context){
	username := c.Param("username")

	if reflect.TypeOf(username).Kind() != reflect.String {
        c.JSON(400, gin.H{"error": "username inválido"})
        return
    }

	var user models.User
	result := initializers.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		c.JSON(404, gin.H{
			"Not found": "Usuário não encontrado"})
		return
	}

	initializers.DB.Delete(&user)
	
	c.JSON(200, gin.H{
		"message": "usuário apagado",
	})
}
