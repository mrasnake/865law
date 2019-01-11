package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

var router *mux.Router

func routes(){

stat1 := http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets/")))

router.PathPrefix("/assets/").Handler(stat1)
router.HandleFunc("/hello", hello)
router.HandleFunc("/layout", layout)
}
