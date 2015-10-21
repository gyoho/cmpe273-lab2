package main


import (
    // Standard library packages
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "log"

    // Third party packages
    "github.com/julienschmidt/httprouter"
)

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func postHello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        log.Fatal(err)
    }
    defer req.Body.Close()

    var contents map[string]interface{}
    err = json.Unmarshal(body, &contents)
    if err != nil {
        log.Fatal(err)
    }

    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(200)

    jsonStr := `{
                    "greeting" : "Hello, ` +  contents["name"].(string) + `!"
                }`

    fmt.Fprintf(rw, "%s\n", jsonStr)
}

func main() {
    mux := httprouter.New()

    mux.GET("/hello/:name", hello)

    mux.POST("/hello", postHello)

    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }

    server.ListenAndServe()
}
