package main
 
import (
	"net/http"
)
 
func enableCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
		}
		handler.ServeHTTP(w, r)
	})
}
 
func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
 
func main() {
	app := enableCORS(http.HandlerFunc(helloWorld)) 
	http.ListenAndServe(":8080", app)
}
