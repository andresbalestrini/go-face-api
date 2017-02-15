package face_service

import(
    "net/http"	
	"io/ioutil"    
)
func Readbody(response *http.Response) (string, error){	
	htmlData, err := ioutil.ReadAll(response.Body)

 	if err != nil {
		 str:=""
		 return str,err
 	}	
	str := string(htmlData)
	return str,err
}