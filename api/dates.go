package api

import "fmt"

func GetAllDates() (DateIndex, error) {
	var dates DateIndex
	err := fetchJSON(DATES_URL, &dates)
	if err != nil {
		return DateIndex{}, err
	}
	return dates, nil
}

func GetDateByURL(url string) (DateData, error) {
	var date DateData
	err := fetchJSON(url, &date)
	if err != nil {
		return DateData{}, err
	}
	return date, nil
}

func GetDateByID(id int) (DateData, error) {
	dates, err := GetAllDates()
	if err != nil {
		return DateData{}, err
	}

	for _, date := range dates.Index {
		if date.ID == id {
			return date, nil
		}
	}

	return DateData{}, fmt.Errorf("les dates pour l'artiste ID %d n'ont pas été trouvées", id)
}