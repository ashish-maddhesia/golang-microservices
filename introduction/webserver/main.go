package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", hello).Methods("GET")
	

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
func hello (w http.ResponseWriter, r *http.Request){
	fmt.Println("might this will be better")
	
}
	
