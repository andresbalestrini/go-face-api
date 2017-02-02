package main

import(

	"github.com/gin-gonic/gin"    
	"net/http"	
)


func mapUrlsToControllers() {

	Router.GET("/ping", Ping)    

}

func Ping(c *gin.Context) { 

	c.String(http.StatusOK, "pong")


}
