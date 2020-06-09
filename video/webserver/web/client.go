package main
// client:用http的client来代理发送我们真正的apirequest
import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var httpClient *http.Client

func init(){
	httpClient = &http.Client{}
}

// 我们采用额的方式是重新组装http请求的方法
// 我们把最原先的http请求进行分解 提取其中的参数
// 然后进行组装成新的http请求，进行发送http请求
// 就可以进行api透传了
func request(b *ApiBody,w http.ResponseWriter,r *http.Request){
	var resp *http.Response
	var err error

	switch b.Method {
		case http.MethodGet:
			// 第一个参数是方式  
			// 第二个参数是地址
			// 第三个参数是可选的身体部分 参数之类的
			req,_:=http.NewRequest("GET",b.Url,nil)
			// 进行复制透传
			req.Header = r.Header
			// 发送请求
			resp,err=httpClient.Do(req)
			if err!=nil{
				log.Printf("err")
				return
			}
			normalResponse(w,resp)
		case http.MethodPost:
			req,_:=http.NewRequest("POST",b.Url,bytes.NewBuffer([]byte(b.ReqBody)))
			req.Header=r.Header
			resp,err=httpClient.Do(req)
			if err!=nil{
				log.Printf("err")
				return
			}
			normalResponse(w,resp)
		case http.MethodDelete:
			req,_:=http.NewRequest("DELETE",b.Url,bytes.NewBuffer([]byte(b.ReqBody)))
			req.Header=r.Header
			resp,err=httpClient.Do(req)
			if err!=nil{
				log.Printf("err")
				return
			}
			normalResponse(w,resp)
		default:
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w,"Bad api equest")
	}
}

func normalResponse(w http.ResponseWriter,r *http.Response){
	res,err:=ioutil.ReadAll(r.Body)
	if err!=nil{
		re,_:=json.Marshal(ErrorInternalFaults)
		w.WriteHeader(500)
		io.WriteString(w,string(re))
		return
	}

	w.WriteHeader(r.StatusCode)
	io.WriteString(w,string(res))
}