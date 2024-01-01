package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//"time"

type Artist_Struct struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func Json_Artist() ([]Artist_Struct, error) {
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	var artists []Artist_Struct
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}
	json.Unmarshal(body, &artists)
	return artists, nil
}
