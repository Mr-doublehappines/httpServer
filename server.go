package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {
		//1.接收客户端 request，并将 request 中带的 header 写入 response header
		for k, v := range r.Header {
			w.Header()[k] = v
		}

		//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
		val := os.Getenv("VERSION")
		w.Header().Set("VERSION",val)


		//3.Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
		fmt.Println("clientIP :", r.RemoteAddr)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello !!!"))
		fmt.Println(http.StatusOK)


	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, "200\n")
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte("hello world\n"))
		fmt.Println("健康检查响应成功")

	})
	err := http.ListenAndServe("0.0.0.0:7848", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
