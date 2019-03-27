package service

import (
    "fmt"
    "net/http"
)

func downHandler() http.HandlerFunc {

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
            getWPool().DelIP(ip)
        } else if v == "rr" {
            getRResultPool().DelIP(ip)
        } else if v == "wr" {
            getWResultPool().DelIP(ip)
        } else {
            getRPool().DelIP(ip)
        }
    }
}
