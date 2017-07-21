package applemusic

// Curator represents a curator of resources.
type Curator struct {
	Id   string `json:"id"`
	Type string `json:"type"`
	Href string `json:"href"`
}

type Curators struct {
	Data []Curator `json:"data"`
	Href string    `json:"href,omitempty"`
}
