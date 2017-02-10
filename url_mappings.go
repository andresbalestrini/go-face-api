package main

import(
	"github.com/andresbalestrini/go-face-api/controllers/face_controller"
	"github.com/gin-gonic/gin"    
	"net/http"	
)


func mapUrlsToControllers() {

	Router.GET("/ping", Ping)
	Router.GET("/permiso", face_controller.Permissions)
	//Router.GET("/familia", face_controller.Family)
	Router.POST("/publicar", face_controller.Publish)
}

func Ping(c *gin.Context) { 

	c.String(http.StatusOK, "pong")
}
