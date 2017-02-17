package face_service

import(
    "net/http"	
	"io/ioutil"    
)
func Readbody(response *http.Response) (string, error) {
	// ReadAll lee desde hasta un error o EOF y devuelve los datos que lee. Una llamada exitosa devuelve err == nil, no err == EOF.	
	htmlData, err := ioutil.ReadAll(response.Body)

 	if err != nil {
		 str:=""
		 return str,err
 	}	
	str := string(htmlData)
	return str,err
}