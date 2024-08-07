package models

// Country struct to represent the JSON object
type Country struct {
	Name           string         `json:"name"`
	TopLevelDomain []string       `json:"topLevelDomain"`
	Alpha2Code     string         `json:"alpha2Code"`
	Alpha3Code     string         `json:"alpha3Code"`
	CallingCodes   []string       `json:"callingCodes"`
	Capital        string         `json:"capital"`
	AltSpellings   []string       `json:"altSpellings"`
	Subregion      string         `json:"subregion"`
	Region         string         `json:"region"`
	Population     int            `json:"population"`
	Latlng         []float64      `json:"latlng"`
	Demonym        string         `json:"demonym"`
	Area           float64        `json:"area"`
	Timezones      []string       `json:"timezones"`
	Borders        []string       `json:"borders"`
	NativeName     string         `json:"nativeName"`
	NumericCode    string         `json:"numericCode"`
	Flags          Flags          `json:"flags"`
	Currencies     []Currency     `json:"currencies"`
	Languages      []Language     `json:"languages"`
	Translations   Translations   `json:"translations"`
	Flag           string         `json:"flag"`
	RegionalBlocs  []RegionalBloc `json:"regionalBlocs"`
	Cioc           string         `json:"cioc"`
	Independent    bool           `json:"independent"`
}

type Flags struct {
	Svg string `json:"svg"`
	Png string `json:"png"`
}

type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type Language struct {
	Iso639_1   string `json:"iso639_1"`
	Iso639_2   string `json:"iso639_2"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

type Translations struct {
	Br string `json:"br"`
	Pt string `json:"pt"`
	Nl string `json:"nl"`
	Hr string `json:"hr"`
	Fa string `json:"fa"`
	De string `json:"de"`
	Es string `json:"es"`
	Fr string `json:"fr"`
	Ja string `json:"ja"`
	It string `json:"it"`
	Hu string `json:"hu"`
}

type RegionalBloc struct {
	Acronym string `json:"acronym"`
	Name    string `json:"name"`
}
