package main
 
import (
	"fmt"
	"time"
    "net/http"
)
 
func logRequests(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    	now := time.Now()
		method := r.Method
		path := r.URL.Path
		fmt.Printf("%v %s %s\n", now.Format("Jan 2, 2006 at 3:04pm (MST)"), method, path)
	    handler.ServeHTTP(w, r)
    })
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello world!"))
}
 
func main() {
    app := logRequests(http.HandlerFunc(helloWorld)) 
    http.ListenAndServe(":8080", app)
}
