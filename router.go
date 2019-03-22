package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

var router *mux.Router

func routes(){

stat1 := http.StripPrefix("/assets/", http.FileServer(http.Dir("./public/assets/")))

router.PathPrefix("/assets/").Handler(stat1)
router.HandleFunc("/home", home)
router.HandleFunc("/", index)
router.HandleFunc("/aboutUs", aboutUs)
router.HandleFunc("/attorneys", attorneys)
router.HandleFunc("/realEstate", realEstate)
router.HandleFunc("/business", business)
router.HandleFunc("/civilLitigation", civilLitigation)
router.HandleFunc("/criminalLitigation", criminalLitigation)
router.HandleFunc("/lawFirms", lawFirms)
router.HandleFunc("/resources", resources)
router.HandleFunc("/contactUs", contactUs)
router.HandleFunc("/api/email", emailForm)
}
