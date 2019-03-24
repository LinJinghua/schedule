package service

import (
    "fmt"
    "net/http"
)

func upHandler() http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        if ip, err := getRequestIP(req); err != nil {
            fmt.Fprintln(w, err.Error())
        } else {
            AddIP(ip)
        }
    }
}
