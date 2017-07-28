package applemusic

// AppleCuratorAttributes represents the attributes of the resource.
type AppleCuratorAttributes struct {
	URL            string          `json:"url"`
	Name           string          `json:"name"`
	Artwork        Artwork         `json:"artwork"`
	EditorialNotes *EditorialNotes `json:"editorialNotes,omitempty"`
}

// AppleCuratorRelationships represents a to-one or to-many relationship from one resource object to others.
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

// AppleCurators represents a list of apple curators.
type AppleCurators struct {
	Data []AppleCurator `json:"data"`
	Href string         `json:"href,omitempty"`
	Next string         `json:"next,omitempty"`
}
