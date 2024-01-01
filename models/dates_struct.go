package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Date_Struct struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	}
}

func Json_Dates() (Date_Struct, error) {
	var date Date_Struct

	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return date, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return date, err
	}
	json.Unmarshal(body, &date)
	return date, nil
}
