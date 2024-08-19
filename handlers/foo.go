package handlers

import "net/http"

func HandleFoo(w http.ResponseWriter , r *http.Request) error{
   
   _,err:= w.Write([]byte("Bar"));
   return err
}