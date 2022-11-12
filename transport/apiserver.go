package transport

import "net/http"

func Start() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	http.ListenAndServe(":8000", mux)
}
