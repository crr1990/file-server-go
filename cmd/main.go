package main

import (
	"net/http"
	"log"
	"os"
	"io"
)

type  ReturnData struct{
	Code int
	Message string
}

func main() {
	http.HandleFunc("/upload", uploader)


	err := http.ListenAndServe(":9090",nil)

	if err != nil {
		log.Fatal(err)
	}
}

func uploader(w http.ResponseWriter, r *http.Request)  {

	//判断请求方式
	if r.Method == "POST" {
		//设置内存大小
		r.ParseMultipartForm(32 << 20);
		//获取上传的第一个文件
		file, header, err := r.FormFile("file");
		defer file.Close();
		if err != nil {
			log.Fatal(err);
		}
		//创建上传目录
		os.Mkdir("./upload", os.ModePerm);
		//创建上传文件
		cur, err := os.Create("./upload/" + header.Filename);
		defer cur.Close();
		if err != nil {
			log.Fatal(err);
		}
		//把上传文件数据拷贝到我们新建的文件
		io.Copy(cur, file);
	}

}
