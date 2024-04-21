package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "snake/docs"
)

// @title Snake Game Swagger API
// @version 1.0
// @description This is snake game api document.
func main(){
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run("0.0.0.0:8080")
}