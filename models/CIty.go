package models

type City struct {
	Id         int
	Name       string
	Population int32
}

type Country struct {
	Name struct {
		Common     string `json:"common"`
		Official   string `json:"official"`
		NativeName struct {
			Ukr struct {
				Official string `json:"official"`
				Common   string `json:"common"`
			} `json:"ukr"`
		} `json:"nativeName"`
	} `json:"name"`
	Cca2        string `json:"cca2"`
	Independent bool   `json:"independent"`
	Status      string `json:"status"`
	UnMember    bool   `json:"unMember"`
	Idd         struct {
		Root     string   `json:"root"`
		Suffixes []string `json:"suffixes"`
	} `json:"idd"`
	Capital   []string `json:"capital"`
	Region    string   `json:"region"`
	Subregion string   `json:"subregion"`
	Languages struct {
		Ukr string `json:"ukr"`
	} `json:"languages"`
	Landlocked bool     `json:"landlocked"`
	Borders    []string `json:"borders"`
	Area       int      `json:"area"`
	Demonyms   struct {
		Eng struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"eng"`
		Fra struct {
			F string `json:"f"`
			M string `json:"m"`
		} `json:"fra"`
	} `json:"demonyms"`
	Flag string `json:"flag"`
	Maps struct {
		GoogleMaps     string `json:"googleMaps"`
		OpenStreetMaps string `json:"openStreetMaps"`
	} `json:"maps"`
	Population int    `json:"population"`
	Fifa       string `json:"fifa"`
	Car        struct {
		Signs []string `json:"signs"`
		Side  string   `json:"side"`
	} `json:"car"`
	Timezones  []string `json:"timezones"`
	Continents []string `json:"continents"`
	Flags      struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags"`
	CoatOfArms struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"coatOfArms"`
	StartOfWeek string `json:"startOfWeek"`
	CapitalInfo struct {
		Latlng []float64 `json:"latlng"`
	} `json:"capitalInfo"`
	PostalCode struct {
		Format string `json:"format"`
		Regex  string `json:"regex"`
	} `json:"postalCode"`
}
