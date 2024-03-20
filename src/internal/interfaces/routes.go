// internal/interfaces/routes.go
package interfaces

import "net/http"

type Route struct {
    Path    string
    Handler http.HandlerFunc
    Method  string
}

type Routes struct {
    Routes []*Route
}
