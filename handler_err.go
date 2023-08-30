package main

import "net/http"
//this is the func signature you have to use if you want to desin a http handler the go std lib expects 
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "enthokkeyo thett patti mwonu")
}
