package service

import (
    "fmt"
    "net/http"
)

func downHandler() http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        if ip, err := getRequestIP(req); err != nil {
            fmt.Fprintln(w, err.Error())
        } else {
            DelIP(ip)
        }
    }
}
