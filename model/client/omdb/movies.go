package omdb

type OmdbGetByTitle struct {
	Search       []OmdbMovie `json:"Search"`
	TotalResults string      `json:"totalResults"`
	Response     string      `json:"Response"`
}

type OmdbMovie struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbId string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type OmdbGetByImdbIdResponse struct {
	Title        string                   `json:"Title"`
	Year         string                   `json:"Year"`
	Rated        string                   `json:"Rated"`
	Released     string                   `json:"Released"`
	Runtime      string                   `json:"Runtime"`
	Genre        string                   `json:"Genre"`
	Director     string                   `json:"Director"`
	Writer       string                   `json:"Writer"`
	Actors       string                   `json:"Actors"`
	Plot         string                   `json:"Plot"`
	Language     string                   `json:"Language"`
	Country      string                   `json:"Country"`
	Awards       string                   `json:"Awards"`
	Poster       string                   `json:"Poster"`
	Ratings      []OmdbGetByImdbIdRatings `json:"Ratings"`
	Metascore    string                   `json:"Metascore"`
	ImdbRating   string                   `json:"imdbRating"`
	ImdbVotes    string                   `json:"imdbVotes"`
	ImdbId       string                   `json:"imdbID"`
	Type         string                   `json:"Type"`
	TotalSeasons string                   `json:"totalSeasons"`
	Response     string                   `json:"Response"`
}

type OmdbGetByImdbIdRatings struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}
