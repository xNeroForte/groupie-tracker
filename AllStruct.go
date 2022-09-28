package gtracker

// ArtistsStruct = []Artist <----- Artist

type ArtistsStruct struct {
	Tab []Artist
}

type ArtistRaw struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Artist struct {
	Id             int          `json:"id"`
	Image          string       `json:"image"`
	Name           string       `json:"name"`
	Members        []string     `json:"members"`
	CreationDate   int          `json:"creationDate"`
	FirstAlbum     string       `json:"firstAlbum"`
	DatesLocations []LocAndDate `json:"relations"`
}

// LocationStruct = index <---- []location <---------- location

type LocationTabApi struct {
	LocationTab []Location `json:"index"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Date      string   `json:"dates"`
}

type RelationTabApi struct {
	Relations []Relation `json:"index"`
}
