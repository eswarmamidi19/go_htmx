package handlers

import (
	
	"net/http"

	"github.com/eswarmamidi19/chi_htmx/views"
)

func HandleFoo(w http.ResponseWriter , r *http.Request) error{
    return Render(w,r,views.Index())
}