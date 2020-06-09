package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// r 里面包含着httpserver httpserver里面含有这cookie方法
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")

	// 访客
	// 无值的时候  没有登录的
	if err1 != nil || err2 != nil {
		// 需要赋值的参数
		p := &HomePage{Name: "孙潇"}
		// 将html认定为模板文件
		t, e := template.ParseFiles("D:/Github/Streaming-video/src/webserver/template/home.html")
		if e != nil {
			log.Printf("parsing template home.html error:%s", e)
			return
		}
		// t为模板 p写入到w传递给模板t
		t.Execute(w, p)
		return
	}

	// 用户
	// 有值得时候  登陆了的
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "D:/Github/Streaming-video/src/webserver/template/userhome.html", http.StatusFound)
		return
	}
}

func userHomeHandler(w http.ResponseWriter,r *http.Request,ps httprouter.Params){
	// 获取值
	cname,err:=r.Cookie("username")	//Cookie值用于登陆过的
	fname:=r.FormValue("username")	
	var p *UserPage

	// 未登录的
	if err!=nil  {
		http.Redirect(w,r,"/",http.StatusFound)
		return
	}

	// 登录了的
	//Cookie值
	if len(cname.Value) !=0 {
		p=&UserPage{Name: cname.Value}
	//填充值
	}else if len(fname)!=0{
		p=&UserPage{Name: fname}
	}

	// 注册页面
	// p=&UserPage{Name:"孙潇2"}
	t,e:=template.ParseFiles("D:/Github/Streamingmedia/src/webserver/template/userhome.html")

	if e !=nil {
		log.Printf("parsing template home.html error:%s",e)
		return
	}
	t.Execute(w,p)
	return
}

func apiHandler(w http.ResponseWriter,r *http.Request,ps httprouter.Params){
	// 方式get  转发一般是post方式的
	if r.Method!=http.MethodPost {
		re,_:=json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w,string(re))
		return
	}

	// body
	// ioutil流失读写
	res,_:=ioutil.ReadAll(r.Body)
	apibody:=&ApiBody{}
	// 转化为非json形式
	if err:=json.Unmarshal(res,apibody);err!=nil{
		re,_:=json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w,string(re))
		return
	}
	// 透传的调用
	// request
	// 第一个参数请求的body
	// 第二个参数响应
	// 第三个参数请求
	request(apibody,w,r)

	defer r.Body.Close()
}

func proxyHandler(w http.ResponseWriter,r *http.Request,ps httprouter.Params){
	// 目标url
	u,_:=url.Parse("http://127.0.0.1:9000/")
	// 创建一个proxy代理（进行替换域名）
	proxy:=httputil.NewSingleHostReverseProxy(u) 
	// 地址进行转发
	proxy.ServeHTTP(w,r)
}
