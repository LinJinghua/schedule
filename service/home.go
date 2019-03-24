package service

import (
    "fmt"
    "net/http"
)

func homeHandler() http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "{ \"id\" : %q, \"content\" : %q }", "1", "succ")
    }
}
