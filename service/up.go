package service

import (
    "fmt"
    "net/http"
)

func upHandler() http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        var ip string
        if ip = req.Form.Get("ip"); ip == "" {
            var err error
            if ip, err = getRequestIP(req); err != nil {
                fmt.Fprintln(w, err.Error())
                return
            }
        }
        if v := req.Form.Get("type"); v == "w" {
            getWPool().AddIP(ip)
        } else if v == "rr" {
            getRResultPool().AddIP(ip)
        } else if v == "wr" {
            getWResultPool().AddIP(ip)
        } else {
            getRPool().AddIP(ip)
        }
    }
}
