package main

import (
	"fmt"
	"html/template"
	"net/http"
	handler "test/handlers"
)

func main() {

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		Data := struct{ ResponseBody string }{ResponseBody: ""}
		tmpl.Execute(w, Data)
	})

	http.HandleFunc("/submit", handler.GetMethod)
	http.HandleFunc("/send", handler.PostMethod)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":5000", nil)

}
