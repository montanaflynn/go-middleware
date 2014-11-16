package main
 
import (
	"log"
	"time"
	"net/http"
)
 
func delay(handler http.Handler, length string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeout, err := time.ParseDuration(length)
		if err != nil {
			log.Println("Could not parse delay length, skipping delay.")
			handler.ServeHTTP(w, r)
		}
		time.Sleep(timeout)
		handler.ServeHTTP(w, r)
	})
}
 
func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}
 
func main() {
	app := delay(http.HandlerFunc(helloWorld), "5s") 
	http.ListenAndServe(":8080", app)
}
