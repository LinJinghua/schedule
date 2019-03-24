package service

import (
    "fmt"
    "net/http"
)

func getHandler() http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprint(w, GetIPFree())
    }
}
