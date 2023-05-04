package handlers

type Errors struct {
	Status  int
	Message string
}

type Artist struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Locations    string
	ConcertDates string
	Relations    string
}

type Concert struct {
	Id             int
	DatesLocations map[string][]string
}

type AllData struct {
	Main     Artist
	Concerts map[string][]string
}

type Place struct {
	Index []struct {
		Id             int
		DatesLocations map[string][]string
	}
}

type SearchBar struct {
	Group  []Artist
	Places Place
}

type SearchInput struct {
	Group   []Artist
	People  []string
	Created []int
	Places  []string
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func ContainsInt(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}

	return false
}
