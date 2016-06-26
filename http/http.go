package main

import (
    "fmt"
    "net/http"
)

var i = 1
func handler(w http.ResponseWriter, r *http.Request) {
  //  fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    fmt.Println(i)
    i++
    fmt.Fprintf(w, "Hi there, I love Body: %s!", (r.Body))
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}