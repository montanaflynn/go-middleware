# go-middleware

An attempt to collect and present all the middlewares that are compatible with http.Handler.

- [Enable CORS](https://github.com/montanaflynn/go-middleware/blob/master/enable-cors.go)
- [Log Requests](https://github.com/montanaflynn/go-middleware/blob/master/log-requests.go)
- [Delay](https://github.com/montanaflynn/go-middleware/blob/master/delay.go)

### Usage

Here's an example of enabling CORS using all three middlewares together:

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
    app := delay(logRequests(enableCORS(http.HandlerFunc(helloWorld))), "5s")
    http.ListenAndServe(":8080", app)
}
```

As you can see if you're using multiple middlewares things can quickly become ugly. That's where [Alice](http://www.github.com/justinas/alice) comes in. Alice is a super lightweight wrapper that let's you chain middleware together in a nice way. If you are using 3 or more middlewares I highly suggest taking a look.

### Todos

- Add more middlewares and functionality to existing ones
- Write examples for HttpRouter, Gin, Goji, Gorilla, etc...
- Create a middleware library that can be imported 
