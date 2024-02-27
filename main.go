

package main

import (
    "fmt"
    "net/http"
    "GROUPIE-TRACKER/Handlers"
)

func main() {
    http.HandleFunc("/", Handlers.IndexHandler)
    http.HandleFunc("/liste", Handlers.ListeHandler)
    http.HandleFunc("/info", Handlers.InfoHandler)
    fmt.Println("Serveur démarré sur le port 8080, http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
