package main

import "net/http"

func main() {
	httpMux := http.NewServeMux()
	httpMux.Handle("/app/",
		http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	httpMux.HandleFunc("/healthz", readinessCheck)
	http.ListenAndServe(":8080", httpMux)

}

func readinessCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Add("Content-Type:", "text/plain; charset=utf-8")
	w.Write([]byte("OK"))
}
