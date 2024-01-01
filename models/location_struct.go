package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Location_Struct struct {
	Index []struct  {
		Id       int      `json:"id"`
		Locations []string `json:"locations"`
	}
}

func Json_Location() (Location_Struct, error) {
	var Location Location_Struct

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return Location, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return Location, err
	}
	json.Unmarshal(body, &Location)
	return Location, nil
}
