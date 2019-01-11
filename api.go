package main

import (
  "net/http"
  // http.ResponseWriter, r *http.Request"log"
)


func hello(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "helloWorld", nil)
  return

}

func layout(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "layout", nil)
  return

}
