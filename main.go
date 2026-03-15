package main

import "net/http"

func main() {
	httpMux := http.NewServeMux()
	http.ListenAndServe(":8080", httpMux)

}
