package applemusic

type GenreAttributes struct {
	Name string `json:"name"`
}

// Genre represents a genre for resources.
type Genre struct {
	Id         string          `json:"id"`
	Type       string          `json:"type"`
	Href       string          `json:"href"`
	Attributes GenreAttributes `json:"attributes"`
}

type Genres struct {
	Data []Genre `json:"data"`
	Href string  `json:"href,omitempty"`
}
