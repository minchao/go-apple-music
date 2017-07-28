package applemusic

// Curator represents a curator of resources.
type Curator struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Href string `json:"href"`
}

// Curators represents a list of curators.
type Curators struct {
	Data []Curator `json:"data"`
	Href string    `json:"href,omitempty"`
	Next string    `json:"next,omitempty"`
}
