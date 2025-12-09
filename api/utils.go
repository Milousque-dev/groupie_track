package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ARTISTS_URL   =  "https://groupietrackers.herokuapp.com/api/artists"
	LOCATIONS_URL =  "https://groupietrackers.herokuapp.com/api/locations"
	DATES_URL     =  "https://groupietrackers.herokuapp.com/api/dates"
	RELATIONS_URL =  "https://groupietrackers.herokuapp.com/api/relation"
)

func fetchJSON(url string, target interface{}) error {								// fetchJSON fait une requête GET et décode le JSON dans target
																					// interface{} permet de tout prendre en paramètres
	resp, err := http.Get(url)														// 1. Faire la requête HTTP GET
	if err != nil {
		return fmt.Errorf("erreur lors de la requête HTTP : %w", err)
	}
	defer resp.Body.Close()															// IMPORTANT : fermer la connexion après utilisation
	
	if resp.StatusCode != http.StatusOK {											// 2. Vérifier le code de statut HTTP
		return fmt.Errorf("code de statut HTTP invalide : %d", resp.StatusCode)
	}
	
	err = json.NewDecoder(resp.Body).Decode(target)									// 3. Décoder le JSON directement depuis le Body
	if err != nil {
		return fmt.Errorf("erreur lors du décodage JSON : %w", err)
	}
	
	return nil
}
