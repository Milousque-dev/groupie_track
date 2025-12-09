package main

import (
	"fmt"
	"log"
	"net/http"
	"groupie_tracker/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handlers.HomeHandler)

	port := ":8080"
	fmt.Printf("🚀 Serveur démarré sur http://localhost%s\n", port)
	fmt.Println("Appuyez sur Ctrl+C pour arrêter")
	
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("❌ Erreur serveur:", err)
	}
}
