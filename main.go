package main

import(	// indico paquetes propios de go que utilizo

	"fmt" 
	"runtime"
	"github.com/gin-gonic/gin"
)

var(

	Router *gin.Engine

)


func main() {

	startApp()

	
}

func startApp(){ // go busca el main que esta en server.go (este archivo) y busca el main en el caual esta esta funcion y se dirige a la misma

	configureRuntime()
	initGinEngine()
	mapUrlsToControllers()
	Router.Run(":9090")

}

func initGinEngine() {

	Router = gin.Default()
	Router.RedirectFixedPath = false
	Router.RedirectTrailingSlash = false
}

func configureRuntime() {

	numCPU:=runtime.NumCPU()
	fmt.Println(numCPU)
	runtime.GOMAXPROCS(numCPU) // indico que voy a trabajar en paralelo con 4 nucleos

}
