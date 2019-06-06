package main

import (
	"log"
    "fmt"
	"net/http"
	"html/template"
    "os"
    "strconv"
)

// localhost:PORTにアクセスすると呼び出される
func handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		log.Fatal(err)
	}
}

// 送信ボタンを押すと呼び出される
func test_post(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("textarea")
	fmt.Fprintf(w, text)
}


func main() {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	fmt.Printf("Starting server at Port %d", port)
	// localhost:port
	http.HandleFunc("/", handler)
	// localhost:port/fuga
	http.HandleFunc("/testform", test_post)
	// cssファイルを読み込めるようにする
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	// サーバ起動
    http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

