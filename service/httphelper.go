package service

import (
    "fmt"
    "net"
    "net/http"
)

func getRequestIP(req *http.Request) (string, error) {
    ip, _, err := net.SplitHostPort(req.RemoteAddr)
    if err != nil {
        return "", fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
    }
    return ip, nil
}


func writeRequestIP(w http.ResponseWriter, req *http.Request) {
    // [Correct way of getting Client's IP Addresses from http.Request (Golang)]
    // (https://stackoverflow.com/questions/27234861/correct-way-of-getting-clients-ip-addresses-from-http-request-golang)

    ip, port, err := net.SplitHostPort(req.RemoteAddr)
    if err != nil {
        //return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
        fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
    }

    userIP := net.ParseIP(ip)
    if userIP == nil {
        //return nil, fmt.Errorf("userip: %q is not IP:port", req.RemoteAddr)
        fmt.Fprintf(w, "userip: %q is not IP:port", req.RemoteAddr)
        return
    }

    // This will only be defined when site is accessed via non-anonymous proxy
    // and takes precedence over RemoteAddr
    // Header.Get is case-insensitive
    forward := req.Header.Get("X-Forwarded-For")

    fmt.Fprintf(w, "<p>IP: %s</p>\n", ip)
    fmt.Fprintf(w, "<p>Port: %s</p>\n", port)
    fmt.Fprintf(w, "<p>Forwarded for: %s</p>\n", forward)
}
