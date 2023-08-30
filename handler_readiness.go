package main

import "net/http"
//this is the func signature you have to use if you want to desin a http handler the go std lib expects 
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, struct{} {})
}
