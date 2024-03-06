package main

import (
	"GROUPIE-TRACKER/Handlers"
	"fmt"
	"net/http"
)

func main() {
	//template
	http.HandleFunc("/", Handlers.IndexHandler)
	http.HandleFunc("/liste", Handlers.ListeHandler)
	http.HandleFunc("/info", Handlers.InfoHandler)

	//css
	fs := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//server

	fmt.Println("Serveur démarré sur le port 8080, http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
