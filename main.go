package main

import (
  "net/http"
  "log"
  "fmt"
  "github.com/unrolled/render"
  "github.com/gorilla/mux"
)

var viewrender *render.Render

func main(){
  router = mux.NewRouter()
  viewrender = render.New(render.Options{
    Extensions: []string{".html"},
    IsDevelopment: true,
    Layout: "layout",
    UnEscapeHTML: true,
  })

  routes()

  http.Handle("/", router)
  log.Println("Starting server at :4000")
  err := http.ListenAndServe(":4000", nil)
  fmt.Println(err)
}
