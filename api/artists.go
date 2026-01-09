package api

import (
	"fmt"
	"strings"
)

func GetAllArtists() ([]Artist, error) {
	var artists []Artist							
	err := fetchJSON(ARTISTS_URL, &artists)	
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func GetArtistByID(id int) (*Artist, error) {
	artists, err := GetAllArtists()
	if err != nil {
		return nil, err
	}

	for _, artist := range artists {
		if artist.ID == id {
			return &artist, nil
		}
	}

	return nil, fmt.Errorf("l'artiste avec l'ID %d n'a pas été trouvé", id)
}

func GetArtistByName(name string) (*Artist, error) {
	artists, err := GetAllArtists()
	if err != nil {
		return nil, err
	}

	for _, artist := range artists {
		if strings.EqualFold(artist.Name, name) {
			return &artist, nil
		}
	}

	return nil, fmt.Errorf("l'artiste '%s' n'a pas été trouvé", name)
}