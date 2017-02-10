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
	//"fmt"
	//"io/ioutil"
)
var (
 // generate loginURL
	fbConfig = &oauth2.Config {
 		// ClientId: FBAppID(string), ClientSecret : FBSecret(string)
 		// Example - ClientId: "1234567890", ClientSecret: "red2drdff6e2321e51aedcc94e19c76ee"

 		ClientID:     "1787876351474916", // change this to yours
 		ClientSecret: "6101caa72d74674b02667262c834ccb8",
 		RedirectURL:  "http://localhost:9090/permiso", // change this to your webserver adddress
 		Scopes:       []string{"public_profile","publish_actions","user_relationships","user_friends"},
 		Endpoint: oauth2.Endpoint{
 			AuthURL:  "https://www.facebook.com/dialog/oauth",
 			TokenURL: "https://graph.facebook.com/oauth/access_token",
 		},
 	}
)

	 //fmt.Printf("code: %s",code)
	 //fmt.Printf("notacces: %s",notacces)
	 /*
	 if code != "" {

		tok, err := fbConfig.Exchange(oauth2.NoContext, code)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		}
		fmt.Printf("token: %s",tok.AccessToken)
		client := fbConfig.Client(oauth2.NoContext, tok)
		/*
		resp, err1 := client.Get("https://graph.facebook.com/me?fields=id,name,gender,locale,first_name,last_name,age_range&access_token=" + tok.AccessToken)
		

		if err1 != nil {
			c.String(http.StatusBadRequest, err1.Error())
		}
		

		//haciendo post
		message := "usando mi api"
		// Values asigna una clave de cadena a una lista de valores.
		form := url.Values{}
		// Add agrega el valor a la clave. Se añade a cualquier valor existente asociado con clave.
    	form.Add("message", message)
    	form.Add("access_token", tok.AccessToken)
		// ==> https://graph.facebook.com/me/feed?message='message'&access_token='tok.AccessToken'
		req, err2 := http.NewRequest("POST", "https://graph.facebook.com/me/feed?", strings.NewReader(form.Encode()))
		if err2 !=nil {
			c.String(http.StatusBadRequest,err2.Error())
		}

		resp2, err3 := client.Do(req)

		if err3 !=nil {
			c.String(http.StatusBadRequest,err3.Error())
		}
		
		str, err4 :=	face_service.Readbody(resp2) 
		if err4 != nil{
			c.String(http.StatusBadRequest,err4.Error())
		}

		fmt.Printf("valor str: %s",str)
		c.JSON(http.StatusOK,str)

	 } else if notacces == "access_denied" {

		c.String(http.StatusUnauthorized, "Permisos denegados por el usuario")

	} else {		
	 url := fbConfig.AuthCodeURL("state")
	 c.Redirect(http.StatusTemporaryRedirect,url)
	}
	*/

 func Permissions(c *gin.Context){
	 code := c.Query("code")
	 notacces := c.Query("error")
	 if code != ""{
		// obtengo el token para retornarlo
		tok, err := fbConfig.Exchange(oauth2.NoContext, code)		
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}	
		var acctoken token.AccessToken
		acctoken.Token = tok.AccessToken
		c.JSON(http.StatusOK,acctoken)		
		 return
	 }else if notacces == "access_denied"{
		 c.String(http.StatusUnauthorized, "Permisos denegados por el usuario")
		 return
	 } else {
		url := fbConfig.AuthCodeURL("state")
	 	c.Redirect(http.StatusTemporaryRedirect,url)		
	 }	 	 
 }

 func Publish(c *gin.Context){	 

	var newstate message.Data	 
	
	failed := c.BindJSON(&newstate)	

	if failed != nil {
		c.String(http.StatusBadRequest, failed.Error())
		return
	}

	// creo cliente
	//client := fbConfig.Client(oauth2.NoContext, newstate.Token)
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

/*
 func Family(c *gin.Context){
	
	client := fbConfig.Client(oauth2.NoContext, tok)
	
	//haciendo Get
	resp2, err3 := client.Get("https://graph.facebook.com/me/family?")

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
 */