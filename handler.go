package main

import (
  "net/http"
  "html/template"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("hello.html")
  t.Execute(w, "Hello!")
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }

  http.HandleFunc("/hello", helloHandler)

  server.ListenAndServe()
}
