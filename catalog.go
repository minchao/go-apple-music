package applemusic

// CatalogService handles communication with the catalog related methods of the Apple Music API.
type CatalogService service

// Tracks represents a list of songs and music videos.
type Tracks struct {
	Data []Resource `json:"data"`
	Href string     `json:"href,omitempty"`
}
