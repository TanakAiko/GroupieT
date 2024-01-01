package data

type Api_Info struct {
	Artist   []Artist_Struct
	Location Location_Struct
	Dates    Date_Struct
	Relation Relation_Struct
}

type OneArtist struct {
	Artist Artist_Struct
	Locations []string
	Date []string
	Datelocations map[string][]string
}

type SearchResult struct{
	Label string
	ToFind interface{}
	Artist Artist_Struct
}

type Info_OneArtist struct{
	Info Api_Info
	ArtistSelected OneArtist
	Search []SearchResult
}