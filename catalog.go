package applemusic

// CatalogService handles communication with the catalog related methods of the Apple Music API.
type CatalogService service

// Artwork represents an artwork.
type Artwork struct {
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	URL        string `json:"url"`
	BgColor    string `json:"bgColor"`
	TextColor1 string `json:"textColor1"`
	TextColor2 string `json:"textColor2"`
	TextColor3 string `json:"textColor3"`
	TextColor4 string `json:"textColor4"`
	IsMosaic   bool   `json:"isMosaic,"` // Undocumented, Used in Playlists.
}

// EditorialNotes represents notes.
type EditorialNotes struct {
	Standard string `json:"standard"`
	Name     string `json:"name,omitempty"` // Undocumented
	Short    string `json:"short"`
}

// PlayParameters represents play parameters for resources.
type PlayParameters struct {
	Id   string `json:"id"`
	Kind string `json:"kind"`
}
