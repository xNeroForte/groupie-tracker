package gtracker

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Relation struct {
	Id             string       `json:"id"`
	DatesLocations []LocAndDate `json:"datesLocations"`
}

type LocAndDate struct {
	Lieu  string
	Dates []string
}

var (
	body        []byte
	newRelation Relation
)

func SetAllEvents(oldArtist ArtistRaw) Artist {
	resp, err2 := http.Get(oldArtist.Relations)
	var NewArtist Artist
	if err2 != nil {
		fmt.Println("Page introuvable")
		return NewArtist
	}
	var err3 error
	body, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println("Il n'y a pas de body")
		return NewArtist
	}

	// Get ID
	ExtractInfo(string('"')+"id"+string('"')+":", ",", 1)

	// Get First Location
	ExtractInfo(string('"')+"datesLocations"+string('"')+":{"+string('"'), string('"'), 2)
	//"datesLocations":{"

	// Get Other Location
	ExtractInfo(string('"')+"],"+string('"'), string('"'), 3)
	//"],"

	// Get Day
	ExtractInfo(string('"')+":["+string('"'), "]", 4)
	//":["

	NewArtist = RawToArtist(oldArtist)
	return NewArtist
}

func RawToArtist(oldArtist ArtistRaw) Artist {
	var tempArtist Artist
	tempArtist.CreationDate = oldArtist.CreationDate
	tempArtist.FirstAlbum = oldArtist.FirstAlbum
	tempArtist.Id = oldArtist.Id
	tempArtist.Image = oldArtist.Image
	tempArtist.Members = oldArtist.Members
	tempArtist.Name = oldArtist.Name
	tempArtist.DatesLocations = newRelation.DatesLocations
	var tempRelationNil Relation
	newRelation = tempRelationNil
	return tempArtist
}

func ExtractInfo(CheckPoint string, EndPoint string, InfoType int) {
	CheckPointByte := []byte(CheckPoint)
	for i := 0; i < len(body); i++ {
		if body[i] == CheckPointByte[0] {
			Check := true
			for j := 0; j < len(CheckPointByte) && j < len(body); j++ {
				if CheckPointByte[j] != body[i+j] {
					Check = false
					break
				}
				if j == len(CheckPointByte)-1 && Check {
					GetInfo(i+j+1, EndPoint, InfoType)
				}
			}
		}
	}
}

func GetInfo(nb int, EndPoint string, InfoType int) {
	var byby []byte
	for i := nb; i < len(body); i++ {
		if body[i] == ([]byte(EndPoint))[0] {
			break
		}
		byby = append(byby, body[i])
	}
	Title := string(byby)
	SetEventsInfo(Title, InfoType)
}

func SetEventsInfo(Info string, InfoType int) {
	switch InfoType {
	case 1:
		newRelation.Id = Info
	case 2:
		var locTemp LocAndDate
		locTemp.Lieu = Info
		newRelation.DatesLocations = append(newRelation.DatesLocations, locTemp)
	case 3:
		var locTemp LocAndDate
		locTemp.Lieu = Info
		newRelation.DatesLocations = append(newRelation.DatesLocations, locTemp)
	case 4:
		for i := 0; i < len(newRelation.DatesLocations); i++ {
			if i < len(newRelation.DatesLocations) && len(newRelation.DatesLocations[i].Dates) == 0 {
				newRelation.DatesLocations[i].Dates = DateGestor(Info)
				break
			}
		}
	}
}

func DateGestor(dateRaw string) []string {
	date := ""
	var strTab []string
	HaveToAdd := true
	for i := 0; i < len(dateRaw); i++ {
		if i+2 < len(dateRaw) && dateRaw[i] == '"' && dateRaw[i+1] == ',' && dateRaw[i+2] == '"' {
			strTab = append(strTab, date)
			date = ""
			HaveToAdd = false
		}
		if i == len(dateRaw)-1 {
			strTab = append(strTab, date)
		}
		if HaveToAdd {
			date += string(dateRaw[i])
		}
		if i+1 < len(dateRaw) && dateRaw[i] == '"' && dateRaw[i+1] != ',' {
			HaveToAdd = true
		}

	}
	//   07-12-2019","14-12-2019","31-12-2019","11-01-2020","18-01-2020"
	return strTab
}

func NbToString(n int) string {
	var arr []rune
	var result string
	nbr := n
	for {
		if nbr >= 1 || nbr <= -1 {
			var digit int
			if n > 0 {
				digit = nbr % 10
				nbr -= digit
			} else {
				digit = 0 - (nbr % 10)
				nbr += digit
			}
			arr = append(arr, rune(digit+48))
			nbr /= 10
		} else {
			break
		}
	}
	if n == 0 {
		result = "0"
	} else {
		for j := len(arr) - 1; j >= 0; j-- {
			result += string(arr[j])
		}
	}
	return result
}
