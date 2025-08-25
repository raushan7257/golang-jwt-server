package handlers

import (
    "fmt"
    "net/http"
)
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! ")
  
}