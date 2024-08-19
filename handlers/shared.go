package handlers

import (
	"log/slog"
	"net/http"
)


type Apifunc func(http.ResponseWriter , *http.Request) error

func MakeHttpHandler(h Apifunc) http.HandlerFunc{

  return func(w http.ResponseWriter , r *http.Request){
     if err:=h(w,r); err!=nil {
     	 slog.Error(err.Error())
     }
  }
}