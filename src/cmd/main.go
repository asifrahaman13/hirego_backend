// cmd/main.go
package main

import (
    "fmt"
    "net/http"
    "github.com/asifrahaman13/hirego/src/internal/application"
    "github.com/asifrahaman13/hirego/src/internal/infrastructure"
    "github.com/asifrahaman13/hirego/src/internal/interfaces/route"
)

func main() {
    userRepository := infrastructure.NewUserRepository()
    userService := application.NewUserService(userRepository)
    routes := route.SetupRoutes(userService)

    // Create a new HTTP server mux
    mux := http.NewServeMux()

    // Register the routes
    for _, route := range routes.Routes {
        mux.HandleFunc(route.Path, route.Handler)
    }

    // Start the HTTP server on port 8080
    server := &http.Server{
        Addr:    ":8080",
        Handler: mux,
    }

    fmt.Println("Starting server on port 8080...")
    if err := server.ListenAndServe(); err != nil {
        fmt.Printf("Error starting server: %s\n", err)
    }
}
