package gtracker

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// `` = alt + 96

var (
	tmpl         *template.Template
	ArtistTab    []Artist
	ArtistTabRaw []ArtistRaw
	LocationTab  LocationTabApi
	RelationTab  Relation
)

func TryFunc() {
	fmt.Println("CA MARRRRCHEEEEEEEE")
}

func APIRequet() {
	req, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")

	d, _ := ioutil.ReadAll(req.Body)

	json.Unmarshal(d, &ArtistTabRaw)

	reqRelation, _ := http.Get("https://groupietrackers.herokuapp.com/api/relation/2")
	rel, _ := ioutil.ReadAll(reqRelation.Body)
	json.Unmarshal(rel, &RelationTab)
	var ArtistTabVide []Artist
	ArtistTab = ArtistTabVide
	var newArtist Artist
	for i := 0; i < len(ArtistTabRaw); i++ {
		//fmt.Println("La relation est " + ArtistTab[i].Relations)
		newArtist = SetAllEvents(ArtistTabRaw[i])
		ArtistTab = append(ArtistTab, newArtist)
	}
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	var FilteredArtistTab []Artist
	//Filter

	r.ParseForm()
	AllMyMot := r.Form["Filter"]
	if len(AllMyMot) > 0 {

		FilterToUse := AllMyMot[0]
		var FilterNb int = int(FilterToUse[len(FilterToUse)-1]) - 48
		fmt.Println(FilterNb)

		for _, art := range ArtistTab {
			if len(art.Members) == FilterNb && !Contains(FilteredArtistTab, art.Name) {
				FilteredArtistTab = append(FilteredArtistTab, art)
			}
		}

	}

	AllMyMot = r.Form["ReadMore"]
	if len(AllMyMot) > 0 {
		for i := 0; i < len(AllMyMot); i++ {
			fmt.Println(AllMyMot[i])
		}
		FromHtml := AllMyMot[0]

		if FromHtml == "ReadMore" {
			fmt.Println("READMOOOOOOOOOORRRRRRRRRREEEEEEE")
			fmt.Println("READMOOOOOOOOOORRRRRRRRRREEEEEEE")
			fmt.Println("READMOOOOOOOOOORRRRRRRRRREEEEEEE")
			fmt.Println("READMOOOOOOOOOORRRRRRRRRREEEEEEE")
			fmt.Println("READMOOOOOOOOOORRRRRRRRRREEEEEEE")
			fmt.Println("READMOOOOOOOOOORRRRRRRRRREEEEEEE")
		}

	}

	tmpl, _ = template.ParseGlob("./static/*.html")
	APIRequet()
	var AllArtists ArtistsStruct
	if len(AllMyMot) < 1 {
		AllArtists = ArtistsStruct{
			Tab: ArtistTab,
		}
	} else {
		AllArtists = ArtistsStruct{
			Tab: FilteredArtistTab,
		}
	}

	tmpl.ExecuteTemplate(w, "index", AllArtists)
}

func Contains(slice []Artist, elems string) bool {
	for _, v := range slice {
		if v.Name == elems {
			return true
		}
	}
	return false
}
