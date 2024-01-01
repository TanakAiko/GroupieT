package tools

import (
	"fmt"
	md "groupie/models"
)

func GenerateData() md.Api_Info {
	var out md.Api_Info

	artists, err := md.Json_Artist()
	if err != nil {
		fmt.Println(err)
		return out
	}

	location, err := md.Json_Location()
	if err != nil {
		fmt.Println(err)
		return out
	}

	date, err := md.Json_Dates()
	if err != nil {
		fmt.Println(err)
		return out
	}

	relation, err := md.Json_Relation()
	if err != nil {
		fmt.Println(err)
		return out
	}

	out.Artist = artists
	out.Dates = date
	out.Location = location
	out.Relation = relation
	return out
}
