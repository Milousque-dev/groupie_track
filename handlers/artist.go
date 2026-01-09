package handlers

import (
	"groupie_tracker/api"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(path)
	if err != nil || id <= 0 {
		ErrorHandler(w, r, http.StatusBadRequest, "ID d'artiste invalide")
		return
	}

	artist, err := api.GetArtistByID(id)
	if err != nil {
		ErrorHandler(w, r, http.StatusNotFound, "Artiste non trouvé")
		return
	}

	relation, err := api.GetRelationByID(id)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Erreur lors de la récupération des concerts")
		return
	}

	data := struct {
		Artist   *api.Artist
		Relation api.RelationData
	}{
		Artist:   artist,
		Relation: relation,
	}

	tmpl, err := template.ParseFiles("templates/artist.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Erreur lors du chargement du template")
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Erreur lors de l'affichage de la page")
		return
	}
}
