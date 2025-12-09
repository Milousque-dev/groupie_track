package api

import "fmt"

func GetAllLocations() (LocationIndex, error) {
	var locations LocationIndex
	err := fetchJSON(LOCATIONS_URL, &locations)
	if err != nil {
		return LocationIndex{}, err
	}
	return locations, nil
}

func GetLocationByURL(url string) (LocationData, error) {
	var location LocationData
	err := fetchJSON(url, &location)
	if err != nil {
		return LocationData{}, err
	}
	return location, nil
}

func GetLocationByID(id int) (LocationData, error) {
	locations, err := GetAllLocations()
	if err != nil {
		return LocationData{}, err
	}

	for _, loc := range locations.Index {
		if loc.ID == id {
			return loc, nil
		}
	}

	return LocationData{}, fmt.Errorf("locations pour l'artiste ID %d non trouvées", id)
}