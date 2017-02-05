package client
import(

	"net/http"

	"io/ioutil"

)


type Comm struct {

	Method string

	Url string

	Id string
}

func (*Comm)Conn(s *Comm)(body []byte ,err error){

	req,err := http.NewRequest(s.Method,s.Url+"/"+s.Id,nil)

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body,err = ioutil.ReadAll(res.Body)

	return

}



