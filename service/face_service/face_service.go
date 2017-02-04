package face_service

import(

    "fmt"
    "net/http"
	"strings"
    "strconv"    
    "github.com/andresbalestrini/go-face-api/model/token"    
)


func Readhttpbody(response *http.Response) string {

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
func Getaccesstoken(client_id string, code string, secret string, callbackUri string) token.AccessToken {
 	var token token.AccessToken
     //fmt.Println("GetAccessToken")
 	//https://graph.facebook.com/oauth/access_token?client_id=YOUR_APP_ID&redirect_uri=YOUR_REDIRECT_URI&client_secret=YOUR_APP_SECRET&code=CODE_GENERATED_BY_FACEBOOK
 	response, err := http.Get("https://graph.facebook.com/oauth/access_token?client_id=" +
 		client_id + "&redirect_uri=" + callbackUri +
 		"&client_secret=" + secret + "&code=" + code)

 	if err == nil {

 		auth := Readhttpbody(response)

 		//var token AccessToken

 		tokenArr := strings.Split(auth, "&")

 		token.Token = strings.Split(tokenArr[0], "=")[1]
 		expireInt, err := strconv.Atoi(strings.Split(tokenArr[1], "=")[1])

 		if err == nil {
 			token.Expiry = int64(expireInt)
 		}

 		return token
 	}

 	//var token AccessToken

 	return token
 }