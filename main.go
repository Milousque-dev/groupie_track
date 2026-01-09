package main

import (
	"fmt"
	"groupie_tracker/handlers"
	"log"
	"net/http"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomeHandler)

	http.HandleFunc("/artist/", func(w http.ResponseWriter, r *http.Request) {
		if strings.TrimPrefix(r.URL.Path, "/artist/") == "" {
			handlers.ErrorHandler(w, r, http.StatusNotFound, "Page non trouvée")
			return
		}
		handlers.ArtistHandler(w, r)
	})

	port := ":8080"
	fmt.Printf("Serveur Groupie Tracker démarré sur http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Erreur lors du démarrage du serveur:", err)
	}
}
