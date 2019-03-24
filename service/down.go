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
            req.ParseForm()
            if v := req.Form.Get("type"); v == "w" {
                getWPool().DelIP(ip)
            } else {
                getRPool().DelIP(ip)
            }
        }
    }
}
