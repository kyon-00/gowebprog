package main

import (
	"net/http"
	"text/template"
)

func process(w http.ResponseWriter, r *http.Request){
	t,_ := template.ParseFiles("tmpl.html")
	t.Execute(w, "Hello world")
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process",process)
	server.ListenAndServe()
}