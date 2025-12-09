package api

type Artist struct {
	ID              int                  `json:"id"`
	Image           string               `json:"image"`
	Name            string               `json:"name"`
	Members         []string             `json:"members"`
	CreationDate    int                  `json:"creationDate"`
	FirstAlbum      string               `json:"firstAlbum"`
	LocationsURL    string               `json:"locations"`
	DatesURL        string               `json:"concertDates"`
	RelationsURL    string               `json:"relations"`
}

type LocationIndex struct {											//wrapper pour gérer la structure index de l'api
    Index           []LocationData       `json:"index"`
}

type LocationData struct {
	ID              int                  `json:"id"`
	Locations       []string             `json:"locations"`
	DatesURL        string               `json:"dates"`
}

type DateIndex struct {												//wrapper aussi
	Index           []DateData           `json:"index"`
}

type DateData struct {
	ID              int                  `json:"id"`
	Dates           []string             `json:"dates"`
}

type RelationIndex struct {											//wrapper smr
	Index           []RelationData       `json:"index"`
}

type RelationData struct {
	ID              int                  `json:"id"`
	DatesLocations  map[string][]string  `json:"datesLocations"`    //map qui permet d'accéder aux dates (valeur[]) a l'aide de la localisation (clef [string])
}