package applemusic

import (
	"encoding/json"
	"errors"
)

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

// Track represents a songs or music-videos.
type Track []byte

// MarshalJSON returns m as the JSON encoding of m.
func (t Track) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}
	return t, nil
}

// UnmarshalJSON sets *t to a copy of data.
func (t *Track) UnmarshalJSON(data []byte) error {
	if t == nil {
		return errors.New("Track: UnmarshalJSON on nil pointer")
	}
	*t = append((*t)[0:0], data...)
	return nil
}

// Type returns the type of track resource.
func (t Track) Type() string {
	var track struct {
		Type string `json:"type"`
	}
	err := json.Unmarshal(t, &track)
	if err != nil {
		return ""
	}
	return track.Type
}

// Parse parses the Track.
// For recognized Track types, a value of the corresponding struct type will be returned.
func (t Track) Parse() (track interface{}, err error) {
	switch t.Type() {
	case "songs":
		track = &Song{}
	case "music-videos":
		track = &MusicVideo{}
	}
	err = json.Unmarshal(t, &track)
	return track, err
}

// Tracks represents a list of songs and music videos.
type Tracks struct {
	Data []Track `json:"data"`
	Href string  `json:"href,omitempty"`
}
