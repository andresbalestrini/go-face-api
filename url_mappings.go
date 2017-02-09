package main

import(
	"github.com/andresbalestrini/go-face-api/controllers/face_controller"
	"github.com/gin-gonic/gin"    
	"net/http"	
)


func mapUrlsToControllers() {

	Router.GET("/ping", Ping)
	Router.GET("/profile", face_controller.Dataprofile)
}

func Ping(c *gin.Context) { 

	c.String(http.StatusOK, "pong")
}
