package handlers

import (
	"groupie_tracker/api"
	"html/template"
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound, "Page non trouvée")
		return
	}

	artists, err := api.GetAllArtists()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Erreur lors de la récupération des artistes")
		return
	}

	relations, err := api.GetAllRelations()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, "Erreur lors de la récupération des relations")
		return
	}

	filteredArtists := applyFilters(artists, relations.Index, r)

	data := struct {
		Artists   []api.Artist
		Relations []api.RelationData
	}{
		Artists:   filteredArtists,
		Relations: relations.Index,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
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

func applyFilters(artists []api.Artist, relations []api.RelationData, r *http.Request) []api.Artist {

	creationMin := r.URL.Query().Get("creation_min")
	creationMax := r.URL.Query().Get("creation_max")
	members := r.URL.Query()["members"]

	if creationMin == "" && creationMax == "" && len(members) == 0 {
		return artists
	}

	var filtered []api.Artist
	for _, artist := range artists {
		if creationMin != "" {
			min, _ := strconv.Atoi(creationMin)
			if artist.CreationDate < min {
				continue
			}
		}
		if creationMax != "" {
			max, _ := strconv.Atoi(creationMax)
			if artist.CreationDate > max {
				continue
			}
		}

		if len(members) > 0 {
			memberCount := len(artist.Members)
			matched := false
			for _, m := range members {
				count, _ := strconv.Atoi(m)

				if count == 5 && memberCount >= 5 {
					matched = true
					break
				}
				if memberCount == count {
					matched = true
					break
				}
			}
			if !matched {
				continue
			}
		}

		filtered = append(filtered, artist)
	}

	return filtered
}