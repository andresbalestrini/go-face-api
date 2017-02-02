package controllers

import(
    "net/http"
    "golang.org/x/oauth2"
    "github.com/gin-gonic/gin"

)

func Home(c *gin.Context) {
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
 	//url := fbConfig.AuthCodeURL("")	

 	// Home page will display a button for login to Facebook
	 //Página de inicio mostrará un botón para iniciar sesión en Facebook
	 //http.Redirect(w,r,url,http.StatusTemporaryRedirect)
     c.Redirect(http.StatusTemporaryRedirect,"/ping")

 	//w.Write([]byte("<html><title>Golang Login Facebook Example</title> <body> <a href='" + url + "'><button>Login with Facebook!</button> </a> </body></html>"))
 }