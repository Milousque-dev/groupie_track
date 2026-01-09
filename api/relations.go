package api

import "fmt"

func GetAllRelations() (RelationIndex, error) {
	var relations RelationIndex
	err := fetchJSON(RELATIONS_URL, &relations)
	if err != nil {
		return RelationIndex{}, err
	}
	return relations, nil
}

func GetRelationByURL(url string) (RelationData, error) {
	var relation RelationData
	err := fetchJSON(url, &relation)
	if err != nil {
		return RelationData{}, err
	}
	return relation, nil
}

func GetRelationByID(id int) (RelationData, error) {
	relations, err := GetAllRelations()
	if err != nil {
		return RelationData{}, err
	}

	for _, rel := range relations.Index {
		if rel.ID == id {
			return rel, nil
		}
	}

	return RelationData{}, fmt.Errorf("les relations pour l'artiste ID %d n'ont pas été trouvée", id)
}

func CountTotalConcerts(relation RelationData) int {
	total := 0
	for _, dates := range relation.DatesLocations {       //"_" permet d'ogniner la clef de la map et de ne parcourir que les valeurs
		total += len(dates)								  //donc on ne prend en compte que les dates et pas les lieux
	}
	return total
}

func GetConcertLocations(relation RelationData) []string {
	locations := make([]string, 0, len(relation.DatesLocations))
	for location := range relation.DatesLocations {
		locations = append(locations, location)
	}
	return locations
}