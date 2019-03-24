package service

import (
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
    n := negroni.Classic()
    // n := negroni.New(negroni.NewRecovery())
    mx := mux.NewRouter()

    initRoutes(mx)

    n.UseHandler(mx)
    return n
}

func initRoutes(mx *mux.Router) {
    mx.NotFoundHandler = NotImplementedHandler()
    mx.MethodNotAllowedHandler = NotImplementedHandler()
    mx.HandleFunc("/down", downHandler())
    mx.HandleFunc("/up", upHandler())
    mx.HandleFunc("/get", getHandler())
    mx.HandleFunc("/", homeHandler()).Methods("GET")
}
