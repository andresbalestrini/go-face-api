package face_controller

import(
    "net/http"
	"strings"
	"net/url"
    "golang.org/x/oauth2"
    "github.com/gin-gonic/gin"
	"github.com/andresbalestrini/go-face-api/service/face_service"
	"github.com/andresbalestrini/go-face-api/model/message"
	"github.com/andresbalestrini/go-face-api/model/token"
	"github.com/andresbalestrini/go-face-api/model/profile"
	"encoding/json"
	"fmt"
)
var (
 // generate loginURL
	fbConfig = &oauth2.Config {
 		// ClientId: FBAppID(string), ClientSecret : FBSecret(string)
 		// Example - ClientId: "1234567890", ClientSecret: "red2drdff6e2321e51aedcc94e19c76ee"

 		ClientID:     "1787876351474916", // change this to yours
 		ClientSecret: "6101caa72d74674b02667262c834ccb8",
 		RedirectURL:  "http://localhost:9090/permiso", // change this to your webserver adddress
 		Scopes:       []string{"public_profile","publish_actions","user_relationships","user_location","user_birthday"},
 		Endpoint: oauth2.Endpoint{
 			AuthURL:  "https://www.facebook.com/dialog/oauth",
 			TokenURL: "https://graph.facebook.com/oauth/access_token",
 		},
 	}
	oauthStateString = "facestate"
)

 func Permissions(c *gin.Context){
	 code := c.Query("code")
	 notacces := c.Query("error")
	 state := c.Query("state")
 	if state != oauthStateString && code != "" {
		 c.String(http.StatusBadRequest,"invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		 return	 
	} else if code != ""{
		// obtengo el token para retornarlo
		tok, err := fbConfig.Exchange(oauth2.NoContext, code)		
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}	
		var acctoken token.Token
		acctoken.AccessToken = tok.AccessToken
		c.JSON(http.StatusOK,acctoken)		
		 return
	 } else if notacces == "access_denied"{
		 c.String(http.StatusUnauthorized, "Permisos denegados por el usuario")
		 return
	 } else {
		// Estado es un símbolo para proteger al usuario de ataques CSRF(). 
		// Siempre debe proporcionar una cadena distinta de cero y validar que coincide con el parámetro de consulta de estado en su devolución
		// de llamada de redireccionamiento.
		url := fbConfig.AuthCodeURL(oauthStateString)
	 	c.Redirect(http.StatusTemporaryRedirect,url)		
	 }	 	 
 }

 func Publish(c *gin.Context){	 

	var newstate message.Data
	
	failed := c.BindJSON(&newstate)	

	if failed != nil {
		c.String(http.StatusBadRequest, failed.Error())
		return
	} else if newstate.Token == ""{
		c.String(http.StatusBadRequest,"Error, debe suministrar un Token de acceso")
		return
	}

	// creamos un cliente http
	client := http.Client{}
	//haciendo post
	// Values asigna una clave de cadena a una lista de valores.
	form := url.Values{}
	// Add agrega el valor a la clave. Se añade a cualquier valor existente asociado con clave.
	form.Add("message", newstate.Message)
	form.Add("access_token", newstate.Token)
	// ==> https://graph.facebook.com/me/feed?message='message'&access_token='tok.AccessToken'
	req, err2 := http.NewRequest("POST", "https://graph.facebook.com/me/feed?", strings.NewReader(form.Encode()))
	if err2 !=nil {
		c.String(http.StatusBadRequest,err2.Error())
		return
	}

	resp2, err3 := client.Do(req)

	if err3 !=nil {
		c.String(http.StatusBadRequest,err3.Error())
		return
	}
	
	str, err4 := face_service.Readbody(resp2) 
	if err4 != nil{
		c.String(http.StatusBadRequest,err4.Error())
		return
	}

	c.JSON(http.StatusOK,str)
 }


func Profile(c *gin.Context){

	var accesstoken token.AccessToken
	
	failed := c.BindJSON(&accesstoken)	

	if failed != nil {
		c.String(http.StatusBadRequest, failed.Error())
		return
	} else if accesstoken.Token == ""{
		c.String(http.StatusBadRequest,"Error, debe suministrar un Token de acceso")
		return
	}
	
	resp, err := http.Get("https://graph.facebook.com/me?fields=name,gender,locale,birthday&access_token=" + accesstoken.Token)

	if err != nil{
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	var perfil profile.Dataprofile	
	json.NewDecoder(resp.Body).Decode(&perfil)

	resp1, err1 := http.Get("https://graph.facebook.com/me/family?access_token=" + accesstoken.Token)

	if err1 != nil{
		c.String(http.StatusBadRequest, err1.Error())
		return
	}

	var familia profile.Arrayfamily
	// Un Decoder lee y decodifica los valores JSON de un flujo de entrada.
	// NewDecoder devuelve un nuevo decodificador que lee desde resp1.Body
	// Decode lee el siguiente valor codificado JSON de su entrada y lo almacena en familia.
	// Es decir que los valos JSON del body son decodificados en un Decoder por NewDecoder y luego con Decode se almacenan en familia
	json.NewDecoder(resp1.Body).Decode(&familia)

	var response profile.Response

	response.Family = familia
	response.Profile = perfil

	c.JSON(http.StatusOK, response)
}