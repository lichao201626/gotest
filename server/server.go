package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// myfs.WriteF("a/b/c/name.sx", "asdfal ,sdjf 中国")
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	h := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", h)) // 启动静态文件服务
	http.HandleFunc("/", hand)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Data struct {
	T string
}

func hand(res http.ResponseWriter, req *http.Request) {

	// aw := "asdbs asaas"
	x := Data{T: "sssss"}
	t, _ := template.ParseFiles("../src/index.html")
	t.Execute(res, x)
	// res.Write([]byte(aw))
}

/* func Handler {
	ServeHTTP(ResponseWriter, *Request) {

	}
} */
