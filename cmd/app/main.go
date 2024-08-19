package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eswarmamidi19/chi_htmx/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main(){
        
       if err:=godotenv.Load(); err!=nil{
       	 log.Fatal(err.Error())
       } 

       port  := os.Getenv("PORT")                  
	 router := chi.NewMux()
	 router.Get("/foo" ,  handlers.MakeHttpHandler(handlers.HandleFoo))
	 fmt.Println("Listining on port " , port)
        


	 http.ListenAndServe(port,router)

}


