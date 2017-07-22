package applemusic

type AppleCuratorAttributes struct {
	URL            string          `json:"url"`
	Name           string          `json:"name"`
	Artwork        Artwork         `json:"artwork"`
	EditorialNotes *EditorialNotes `json:"editorialNotes,omitempty"`
}

type AppleCuratorRelationships struct {
	Playlists Playlists `json:"playlists"` // Default inclusion: Identifiers only
}

// AppleCurator represents an Apple curator.
type AppleCurator struct {
	Id            string                    `json:"id"`
	Type          string                    `json:"type"`
	Href          string                    `json:"href"`
	Attributes    AppleCuratorAttributes    `json:"attributes"`
	Relationships AppleCuratorRelationships `json:"relationships"`
}

type AppleCurators struct {
	Data []AppleCurator `json:"data"`
	Href string         `json:"href,omitempty"`
	Next string         `json:"next,omitempty"`
}
