package main

import (
        "fmt"
        "html/template"
        "log"
        "net/http"
        "strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
    // Parse url parameters passed
    r.ParseForm()

    // Print form info on server side
    fmt.Println(r.Form)
    fmt.Println("Path: ", r.URL.Path)
    fmt.Println("Scheme: ", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key: ", k)
        fmt.Println("value: ", strings.Join(v, " "))
    }

    // Write data to response
    fmt.Fprintf(w, "Hello!")
}

func login(w http.ResponseWriter, r *http.Request) {
    // Get request method
    fmt.Println("method: ", r.Method)
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()

        // Logic part of login
        fmt.Println("username: ", template.HTMLEscapeString(r.Form.Get("username"))) // HTMLEscape used to prevent XSS
        fmt.Println("password: ", template.HTMLEscapeString(r.Form.Get("password")))
        template.HTMLEscape(w, []byte(r.Form.Get("username")))
    }
}

func main() {
    // Router rules
    http.HandleFunc("/", sayHelloName)
    http.HandleFunc("/login", login)

    // Setting listening port
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe", err)
    }
}
