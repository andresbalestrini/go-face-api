package face_controller

import(
    "net/http"
    "golang.org/x/oauth2"
    "github.com/gin-gonic/gin"
	"github.com/andresbalestrini/go-face-api/service/face_service"
	"fmt"
)

func Dataprofile(c *gin.Context) {
 	
	 // generate loginURL
	fbConfig := &oauth2.Config {
 		// ClientId: FBAppID(string), ClientSecret : FBSecret(string)
 		// Example - ClientId: "1234567890", ClientSecret: "red2drdff6e2321e51aedcc94e19c76ee"

 		ClientID:     "1787876351474916", // change this to yours
 		ClientSecret: "6101caa72d74674b02667262c834ccb8",
 		RedirectURL:  "http://localhost:9090/profile", // change this to your webserver adddress
 		Scopes:       []string{"public_profile"},
 		Endpoint: oauth2.Endpoint{
 			AuthURL:  "https://www.facebook.com/dialog/oauth",
 			TokenURL: "https://graph.facebook.com/oauth/access_token",
 		},
 	}

	 code := c.Query("code")
	 notacces := c.Query("error")

	 if code != "" { 

		 accessToken := face_service.Getaccesstoken(fbConfig.ClientID, code, fbConfig.ClientSecret, fbConfig.RedirectURL)
		 
		 response, err := http.Get("https://graph.facebook.com/me?fields=id,name,picture,gender,link,locale,cover,first_name,last_name,age_range&access_token=" + accessToken.Token)
		
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}	

		str := face_service.Readhttpbody(response)

		c.JSON(http.StatusOK,str)

	} else if notacces == "access_denied" {

		c.String(http.StatusUnauthorized, "Permisos denegados por el usuario")

	} else {  	
	 //AuthCodeURL devuelve una URL a la página de consentimiento del proveedor de OAuth 2.0 que pide permisos para los ámbitos (scopes) requeridos explícitamente.
 	 url := fbConfig.AuthCodeURL("")

	 //Página de inicio mostrará un botón para iniciar sesión en Facebook
     c.Redirect(http.StatusTemporaryRedirect,url)
	} 	
 }

 func Publish(c *gin.Context) {
	   // generate loginURL
	fbConfig := &oauth2.Config {
 		// ClientId: FBAppID(string), ClientSecret : FBSecret(string)
 		// Example - ClientId: "1234567890", ClientSecret: "red2drdff6e2321e51aedcc94e19c76ee"

 		ClientID:     "1787876351474916", // change this to yours
 		ClientSecret: "6101caa72d74674b02667262c834ccb8",
 		RedirectURL:  "http://localhost:9090/likes", // change this to your webserver adddress
 		Scopes:       []string{"user_likes"},
 		Endpoint: oauth2.Endpoint{
 			AuthURL:  "https://www.facebook.com/dialog/oauth",
 			TokenURL: "https://graph.facebook.com/oauth/access_token",
 		},
 	}

	 code := c.Query("code")
	 notacces := c.Query("error")

	  if code != "" { 

		 accessToken := face_service.Getaccesstoken(fbConfig.ClientID, code, fbConfig.ClientSecret, fbConfig.RedirectURL)
		 fmt.Printf(accessToken.Token)
		 
		 response, err := http.Get("https://graph.facebook.com/me/likes?access_token=" + accessToken.Token)		 
		 //response, err := http.PostForm("https://graph.facebook.com/me/feed?message=hello, world&amp;access_token=" + accessToken.Token , nil)
		
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}	

		str := face_service.Readhttpbody(response)

		c.JSON(http.StatusOK,str)

	} else if notacces == "access_denied" {

		c.String(http.StatusUnauthorized, "Permisos denegados por el usuario")

	} else {  	
	 //AuthCodeURL devuelve una URL a la página de consentimiento del proveedor de OAuth 2.0 que pide permisos para los ámbitos (scopes) requeridos explícitamente.
 	 url := fbConfig.AuthCodeURL("")

	 //Página de inicio mostrará un botón para iniciar sesión en Facebook
     c.Redirect(http.StatusTemporaryRedirect,url)
	} 	
	 

 }