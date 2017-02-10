package main

/*
import (
 	"fmt"
 	//"github.com/antonholmquist/jason"
  "golang.org/x/oauth2"
  "net/http"
 	"strconv"
 	"strings"
 )

 type AccessToken struct {
 	Token  string
 	Expiry int64
 }

 func readHttpBody(response *http.Response) string {

 	fmt.Println("Reading body")

 	bodyBuffer := make([]byte, 5000)
 	var str string

 	count, err := response.Body.Read(bodyBuffer)

 	for ; count > 0; count, err = response.Body.Read(bodyBuffer) {

 		if err != nil {

 		}

 		str += string(bodyBuffer[:count])
 	}

 	return str

 }

 //Converts a code to an Auth_Token
 func GetAccessToken(client_id string, code string, secret string, callbackUri string) AccessToken {
 	//fmt.Println("GetAccessToken")
 	//https://graph.facebook.com/oauth/access_token?client_id=YOUR_APP_ID&redirect_uri=YOUR_REDIRECT_URI&client_secret=YOUR_APP_SECRET&code=CODE_GENERATED_BY_FACEBOOK
 	response, err := http.Get("https://graph.facebook.com/oauth/access_token?client_id=" +
 		client_id + "&redirect_uri=" + callbackUri +
 		"&client_secret=" + secret + "&code=" + code)

 	if err == nil {

 		auth := readHttpBody(response)

 		var token AccessToken

 		tokenArr := strings.Split(auth, "&")

 		token.Token = strings.Split(tokenArr[0], "=")[1]
 		expireInt, err := strconv.Atoi(strings.Split(tokenArr[1], "=")[1])

 		if err == nil {
 			token.Expiry = int64(expireInt)
 		}

 		return token
 	}

 	var token AccessToken

 	return token
 }

 func Home(w http.ResponseWriter, r *http.Request) {
 	//w.Header().Set("Content-Type", "text/html; charset=utf-8")

 	// generate loginURL
 	fbConfig := &oauth2.Config{
 		// ClientId: FBAppID(string), ClientSecret : FBSecret(string)
 		// Example - ClientId: "1234567890", ClientSecret: "red2drdff6e2321e51aedcc94e19c76ee"

 		ClientID:     "1787876351474916", // change this to yours
 		ClientSecret: "6101caa72d74674b02667262c834ccb8",
 		RedirectURL:  "http://localhost:9090/FBLogin", // change this to your webserver adddress
 		Scopes:       []string{"public_profile"},
 		Endpoint: oauth2.Endpoint{
 			AuthURL:  "https://www.facebook.com/dialog/oauth",
 			TokenURL: "https://graph.facebook.com/oauth/access_token",
 		},
 	}
	 //AuthCodeURL devuelve una URL a la página de consentimiento del proveedor de OAuth 2.0 que pide permisos para los ámbitos (scopes) requeridos explícitamente.
 	url := fbConfig.AuthCodeURL("")	

 	// Home page will display a button for login to Facebook
	 //Página de inicio mostrará un botón para iniciar sesión en Facebook
	 http.Redirect(w,r,url,http.StatusTemporaryRedirect)

 	//w.Write([]byte("<html><title>Golang Login Facebook Example</title> <body> <a href='" + url + "'><button>Login with Facebook!</button> </a> </body></html>"))
 }

 func FBLogin(w http.ResponseWriter, r *http.Request) {
 	// grab the code fragment

 	w.Header().Set("Content-Type", "application/json")
 	code := r.FormValue("code")

 	ClientId := "1787876351474916" // change this to yours
 	ClientSecret := "6101caa72d74674b02667262c834ccb8"
 	RedirectURL := "http://localhost:9090/FBLogin"

 	accessToken := GetAccessToken(ClientId, code, ClientSecret, RedirectURL)
	fmt.Printf(accessToken.Token)

 	response, err := http.Get("https://graph.facebook.com/me?fields=id,name,picture,gender,link,locale,cover,first_name,last_name,age_range&access_token=" + accessToken.Token)

 	// handle err. You need to change this into something more robust
 	// such as redirect back to home page with error message
 	if err != nil {
 		w.Write([]byte(err.Error()))
 	}

 	str := readHttpBody(response)
//	 fmt.Printf("Valor de str")
//	 fmt.Printf(str)

 	// dump out all the data
	w.Write([]byte(str))
 }

 func main() {

 	mux := http.NewServeMux()
 	mux.HandleFunc("/", Home)
 	mux.HandleFunc("/FBLogin", FBLogin)

 	http.ListenAndServe(":9090", mux)
 }
 */

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
