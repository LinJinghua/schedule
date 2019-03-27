package service

import (
    "fmt"
    "net/http"
)

func getHandler() http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        if v := req.Form.Get("type"); v == "w" {
            fmt.Fprint(w, getWPool().GetIP())
        } else if v == "rr" {
            fmt.Fprint(w, getRResultPool().GetIP())
        } else if v == "wr" {
            fmt.Fprint(w, getWResultPool().GetIP())
        } else {
            fmt.Fprint(w, getRPool().GetIP())
        }
    }
}
