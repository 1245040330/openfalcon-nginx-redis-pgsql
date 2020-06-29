package util

import (
	"github.com/imroc/req"
	"io"
	"log"
	"os"
)

var (
	Info *log.Logger
	Warning *log.Logger
	Error * log.Logger
)
func init(){
	errFile,err:=os.OpenFile("errors.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("打开errFile日志文件失败：",err)
	}
	infoFile,err:=os.OpenFile("info.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0666)
	if err!=nil{
		log.Fatalln("打开infoFile日志文件失败：",err)
	}
	Info = log.New(io.MultiWriter(os.Stderr,infoFile),"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	Warning = log.New(os.Stdout,"Warning:",log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Ltime | log.Lshortfile)

}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {
	heade := req.Header{
		"Content-Type": contentType,
	}
	r, err := req.Post(url, heade, data)
	if err!=nil{
		Error.Println(err)
	}
	//Info.Println(r.ToString())
	res, err := r.ToString()
	if err!=nil{
		Error.Println(err)
	}
	return res
}

func Get(url string, data interface{}) string {
	r, err := req.Get(url, data)
	if err!=nil{
		Error.Println(err)
	}
	//Info.Println(r.ToString())
	res, err := r.ToString()
	if err!=nil{
		Error.Println(err)
	}
	return res
}