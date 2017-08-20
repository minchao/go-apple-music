package applemusic

import "encoding/json"

// Resource represents a resource—such as an album, song, or playlist—in the Apple Music catalog or iCloud Music Library.
type Resource struct {
	json.RawMessage `json:",inline"`
}

// Type returns the type of resource.
func (r Resource) Type() string {
	var resource struct {
		Type string `json:"type"`
	}
	err := json.Unmarshal(r.RawMessage, &resource)
	if err != nil {
		return ""
	}
	return resource.Type
}

// Parse parses the Resource.
// For recognized Resource types, a value of the corresponding struct type will be returned.
func (t Resource) Parse() (resource interface{}, err error) {
	switch t.Type() {
	case "albums":
		resource = &Album{}
	case "music-videos":
		resource = &MusicVideo{}
	case "playlists":
		resource = &Playlist{}
	case "songs":
		resource = &Song{}
	case "stations":
		resource = &Station{}
	}
	err = json.Unmarshal(t.RawMessage, &resource)
	return resource, err
}

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
	IsMosaic   bool   `json:"isMosaic,omitempty"` // Undocumented, Used in Playlists.
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
