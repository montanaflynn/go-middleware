# go-middleware

An attempt to collect and present all the middlewares that are compatible with http.Handler.

### Middlewares

- [Enable CORS](https://github.com/montanaflynn/go-middleware/blob/master/enable-cors.go)
- [Logging](https://github.com/montanaflynn/go-middleware/blob/master/log-requests.go)

### Usage

Here's an example of enabling CORS using two middlewares:

```go
package main
 
import (
    "fmt"
    "time"
    "net/http"
)
 
func enableCORS(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        if r.Method == "OPTIONS" {
            w.WriteHeader(200)
        }
        handler.ServeHTTP(w, r)
    })
}

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
    app := logRequests(enableCORS(http.HandlerFunc(helloWorld)))
    http.ListenAndServe(":8080", app)
}
```

As you can see if you're using multiple middlewares things can quickly become ugly. That's where [Alice](http://www.github.com/justinas/alice) comes in. Alice is a super lightweight wrapper that let's you write the same thing in a more concise manner. Notice the only difference is the import and in the main():

```go
package main
 
import (
    "fmt"
    "time"
    "net/http"
    "github.com/justinas/alice"
)
 
func enableCORS(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        if r.Method == "OPTIONS" {
            w.WriteHeader(200)
        }
        handler.ServeHTTP(w, r)
    })
}

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
    finalHandler := http.HandlerFunc(helloWorld)
    app := alice.New(enableCORS, logRequests).Then(finalHandler)
    http.ListenAndServe(":8080", app)
}
```
