package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Relation_Struct struct {
	Index []struct {
		Id            int                 `json:"id"`
		Datelocations map[string][]string `json:"datesLocations"`
	}
}

func Json_Relation() (Relation_Struct, error) {
	var relation Relation_Struct

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return relation, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return relation, err
	}
	json.Unmarshal(body, &relation)
	return relation, nil
}
