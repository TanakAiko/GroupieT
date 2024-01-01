package tools

import (
	"fmt"
	md "groupie/models"
	"strconv"
)

type CustomError struct {
	Message string
}

// Implement the error interface with Error()
func (e *CustomError) Error() string {
	return fmt.Sprintf("CustomError: %s", e.Message)
}

func GetArtist(ID string, data md.Api_Info) (md.OneArtist, error) {
	var artist md.OneArtist
	id, err := strconv.Atoi(ID)
	if err != nil || 0 >= id || id > len(data.Artist) {
		return artist, &CustomError{Message: "invalid Id"}
	}

	for i := 0; i < len(data.Artist); i++ {
		if data.Artist[i].Id == id {
			artist.Artist = data.Artist[i]
			artist.Date = data.Dates.Index[i].Dates
			artist.Locations = data.Location.Index[i].Locations
			artist.Datelocations = data.Relation.Index[i].Datelocations
			break
		}
	}

	return artist, nil
}
