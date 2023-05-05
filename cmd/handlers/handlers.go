package handlers

import (
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var arr_cities [][]string

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Errorhandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errorhandler(w, http.StatusMethodNotAllowed)
		return
	}
	html, err := template.ParseFiles("./web/index.html")
	if err != nil {
		log.Println(err.Error())
		Errorhandler(w, http.StatusInternalServerError)
		return
	}

	band, err1 := JsonArtists()
	if err1 != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}

	var members []string
	var created []int
	for _, v := range band {
		for _, vv := range v.Members {
			if !Contains(members, v.Name) {
				members = append(members, vv)
			}
		}
		if !ContainsInt(created, v.CreationDate) {
			created = append(created, v.CreationDate)
		}
	}
	// creating array of cities without duplicates
	locations, err2 := JsonLocations()
	if err2 != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
	var new_locations, res_locations []string

	var arr [][]reflect.Value
	for _, v := range locations.Index {
		arr = append(arr, reflect.ValueOf(v.DatesLocations).MapKeys())
	}

	for _, v := range arr {
		var array []string
		for _, vv := range v {
			vv2 := strings.ReplaceAll(vv.String(), "_", " ")
			vv2 = strings.Title(vv2)
			replacer := strings.NewReplacer("-", ", ", "Uk", "UK", "Usa", "USA")
			res_vv := replacer.Replace(vv2)

			new_locations = append(new_locations, res_vv)

			array = append(array, res_vv)
		}
		arr_cities = append(arr_cities, array)

	}

	// adding cities without duplicates to res_locations - array
	for _, v := range new_locations {
		if !Contains(res_locations, v) {
			res_locations = append(res_locations, v)
		}
	}

	res := SearchInput{
		Group:   band,
		People:  members,
		Created: created,
		Places:  res_locations,
	}
	err = html.Execute(w, res)
	if err != nil {
		log.Println(err.Error())
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	link := r.URL.Path
	linkList := strings.Split(link, "/")
	id, err := strconv.Atoi(linkList[len(linkList)-1])
	if len(linkList) > 3 || linkList[1] != "artist" || (id <= 0 || id > 52) {
		Errorhandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errorhandler(w, http.StatusMethodNotAllowed)
		return
	}
	html, err := template.ParseFiles("./web/artist.html")
	if err != nil {
		log.Println(err.Error())
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
	// putting all necessary data for ArtistPage to res
	mainPage, err3 := JsonArtists()
	if err3 != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
	concerts, err4 := JsonConcerts(strconv.Itoa(id))
	if err4 != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
	MapData := map[string][]string{}

	for key, value := range concerts.DatesLocations {
		key = strings.ReplaceAll(key, "_", " ")
		key = strings.ReplaceAll(key, "-", ", ")
		MapData[key] = value
	}
	res := AllData{
		Main:     mainPage[id-1],
		Concerts: MapData,
	}
	err = html.Execute(w, res)

	if err != nil {
		log.Println(err.Error())
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	// getting value from search input
	text := r.FormValue("text")
	txt := strings.ToLower(text)
	if r.URL.Path != "/search/" || text == "" {
		Errorhandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errorhandler(w, http.StatusMethodNotAllowed)
		return
	}

	html, err := template.ParseFiles("./web/search-bar.html")
	band, err1 := JsonArtists()
	if err1 != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}

	locations, err2 := JsonLocations()
	if err2 != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}

	res := SearchBar{
		Group:  band,
		Places: locations,
	}

	// search name,members,FirstAlbum,CreationDate
	var output []Artist
	var members []string

	for _, v := range res.Group {
		flag := false

		date, _ := strconv.Atoi(txt)
		if strings.Contains(strings.ToLower(v.Name), txt) || v.FirstAlbum == txt || v.CreationDate == date {
			output = append(output, v)
			members = append(members, strings.ToLower(v.Name))
			flag = true
		}
		for _, vv := range v.Members {
			if strings.Contains(strings.ToLower(vv), txt) && !Contains(members, txt) && flag == false {
				output = append(output, v)
				break
			}
		}

	}
	// search place
	for i, v := range arr_cities {
		for _, vv := range v {
			if strings.Contains(vv, txt) {
				output = append(output, res.Group[i])
				break
			}
		}
	}

	if err != nil {
		log.Println(err.Error())
		Errorhandler(w, http.StatusInternalServerError)
		return
	}

	err = html.Execute(w, output)

	if err != nil {
		log.Println(err.Error())
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
}

func Filters(w http.ResponseWriter, r *http.Request) {
	// getting value from filter input
	r.ParseForm() // Parse url parameters passed, then parse the response packet for the POST body (request body)
	// attention: If you do not call ParseForm method, the following data can not be obtained form

	inputs := r.Form
	creation, album, location, numberOfMembers := []string{}, []string{}, []string{}, []string{}
	for k, v := range inputs {
		switch k {
		case "creation-date":
			creation = v
		case "first-album":
			album = v
		case "location":
			location = v
		case "members":
			numberOfMembers = v
		}
	}

	if r.URL.Path != "/filters/" {
		Errorhandler(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		Errorhandler(w, http.StatusMethodNotAllowed)
		return
	}

	html, err := template.ParseFiles("./web/search-bar.html")
	band, err1 := JsonArtists()
	if err1 != nil {
		Errorhandler(w, http.StatusInternalServerError)
		return
	}

	res := SearchBar{
		Group: band,
	}

	// search CreationDate, FirstAlbum, Location, members
	var output []Artist

	if len(creation) != 2 || len(album) != 2 {
		Errorhandler(w, http.StatusNotFound)
		return
	}

	a, err2 := strconv.Atoi(creation[0])
	b, err3 := strconv.Atoi(creation[1])

	c, err4 := strconv.Atoi(album[0])
	d, err5 := strconv.Atoi(album[1])
	if err2 != nil || err3 != nil || err4 != nil || err5 != nil || len(location) == 0 || (a < 1922 || a > 2022) || (b < 1922 || b > 2022) || (c < 1922 || c > 2022) || (d < 1922 || d > 2022) {
		Errorhandler(w, http.StatusNotFound)
		return
	}

	for _, v := range res.Group {

		firstAlbum, _ := strconv.Atoi(strings.Split(v.FirstAlbum, "-")[2])
		if (v.CreationDate >= a && v.CreationDate <= b) && (firstAlbum >= c && firstAlbum <= d) {
			output = append(output, v)
		}

	}
	var result []Artist

	if len(location[0]) != 0 && len(numberOfMembers) != 0 {

		// strings to int
		nm := []int{}
		for _, z := range numberOfMembers {
			val, _ := strconv.Atoi(z)
			nm = append(nm, val)
		}

		for _, v := range output {
			for i, z := range arr_cities {
				if i == v.Id-1 {
					for _, zz := range z {
						if zz == location[0] {
							for _, vv := range nm {
								if len(v.Members) == vv {
									result = append(result, v)
									break
								}
							}
						}
					}
				}
			}
		}
	} else if len(location[0]) != 0 && len(numberOfMembers) == 0 {
		for _, v := range output {
			for i, z := range arr_cities {
				if i == v.Id-1 {
					for _, zz := range z {
						if zz == location[0] {
							result = append(result, v)
						}
					}
				}
			}
		}
	} else if len(location[0]) == 0 && len(numberOfMembers) != 0 {

		// strings to int
		nm := []int{}
		for _, z := range numberOfMembers {
			val, _ := strconv.Atoi(z)
			nm = append(nm, val)
		}

		for _, v := range output {
			for _, vv := range nm {
				if len(v.Members) == vv {
					result = append(result, v)
					break
				}
			}
		}
	} else {
		result = output
	}

	if err != nil {
		log.Println(err.Error())
		Errorhandler(w, http.StatusInternalServerError)
		return
	}

	err = html.Execute(w, result)

	if err != nil {
		log.Println(err.Error())
		Errorhandler(w, http.StatusInternalServerError)
		return
	}
}

func Errorhandler(w http.ResponseWriter, status int) {
	html, err := template.ParseFiles("./web/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var Result Errors
	Result.Status = status
	Result.Message = http.StatusText(status)
	w.WriteHeader(status)
	err = html.Execute(w, Result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
