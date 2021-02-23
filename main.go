package main

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gops01/views"
	_ "gops01/docs"
)


// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://razeen.me

// @contact.name lupuxiao
// @contact.url https://razeen.me
// @contact.email me@razeen.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1

func main() {

	r := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", views.HandleHello)
		v1.POST("/login", views.HandleLogin)
		v1Auth := v1.Use(views.HandleAuth)
		{
			v1Auth.POST("/upload", views.HandleUpload)
			v1Auth.GET("/list", views.HandleList)
			v1Auth.GET("/file/:id", views.HandleGetFile)
			v1Auth.POST("/json", views.HandleJSON)
		}
	}


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}