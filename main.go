package main

import(
	"fmt"
	"net/http"
)

func mainHandler() http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello World!! I'm working! Again now!")
	})
}

func main(){
	http.HandleFunc("/", mainHandler())
	http.ListenAndServe(":8080", nil)
}
