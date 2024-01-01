package tools

import (
	"fmt"
	md "groupie/models"
	"strings"
)

func GetDataToPrint(dataSearch string, data md.Api_Info) []md.SearchResult {
	var searchResult []md.SearchResult
	lowerDataSearch := strings.ToLower(dataSearch)
	for _, artist := range data.Artist {
		var singleSearchResult md.SearchResult
		lowerArtistName := strings.ToLower(artist.Name)
		lowerArtistFirstAlbum := strings.ToLower(artist.FirstAlbum)
		lowerArtistCreationDate := strings.ToLower(fmt.Sprint(artist.CreationDate))
		if strings.Contains(lowerArtistName, lowerDataSearch) {
			singleSearchResult.Artist = artist
			singleSearchResult.ToFind = artist.Name
			singleSearchResult.Label = "Artist/band"
			searchResult = append(searchResult, singleSearchResult)
			continue
		}

		if strings.Contains(lowerArtistFirstAlbum, lowerDataSearch) {
			singleSearchResult.Artist = artist
			singleSearchResult.ToFind = artist.FirstAlbum
			singleSearchResult.Label = "First album date"
			searchResult = append(searchResult, singleSearchResult)
			continue
		}

		if strings.Contains(lowerArtistCreationDate, lowerDataSearch) {
			singleSearchResult.Artist = artist
			singleSearchResult.ToFind = artist.CreationDate
			singleSearchResult.Label = "Creation date"
			searchResult = append(searchResult, singleSearchResult)
			continue
		}

		for _, member := range artist.Members {
			lowerMember := strings.ToLower(member)
			if strings.Contains(lowerMember, lowerDataSearch) {
				singleSearchResult.Artist = artist
				singleSearchResult.ToFind = member
				singleSearchResult.Label = "Member"
				searchResult = append(searchResult, singleSearchResult)
				continue
			}
		}
	}
	
	for _, index := range data.Location.Index {
		for _, location := range index.Locations {
			var singleSearchResult md.SearchResult
			lowerLocation := strings.ToLower(location)
			if strings.Contains(lowerLocation, lowerDataSearch) {
				id := index.Id
				for _, artist := range data.Artist {
					if id == artist.Id {
						singleSearchResult.Artist = artist
						singleSearchResult.ToFind = location
						singleSearchResult.Label = "Location"
						searchResult = append(searchResult, singleSearchResult)
						break
					}
				}
			}
		}
	}
	return searchResult
}
