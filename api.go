package main

import (
  "net/http"
  // http.ResponseWriter, r *http.Request"log"
)


func index(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "/home", 302)
}

func home(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "home", nil)
  return

}

func aboutUs(w http.ResponseWriter, r *http.Request){
  viewrender.HTML(w, http.StatusOK, "aboutUs", nil)
  return

}

func attorneys(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "attorneys", nil)
  return

}

func realEstate(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "realEstate", nil)
  return

}

func business(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "business", nil)
  return

}

func civilLitigation(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "civilLitigation", nil)
  return

}

func resources(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "resources", nil)
  return

}

func contactUs(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "contactUs", nil)
  return

}

func criminalLitigation(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "criminalLitigation", nil)
  return

}

func lawFirms(w http.ResponseWriter, r *http.Request){

  viewrender.HTML(w, http.StatusOK, "lawFirms", nil)
  return

}
