package applemusic

type ActivityAttributes struct {
	URL            string          `json:"url"`
	Name           string          `json:"name"`
	Artwork        Artwork         `json:"artwork"`
	EditorialNotes *EditorialNotes `json:"editorialNotes,omitempty"`
}

type ActivityRelationships struct {
	Playlists Playlists `json:"playlists"` // Default inclusion: Identifiers only
}

// Activity represents an activity.
type Activity struct {
	Id            string                `json:"id"`
	Type          string                `json:"type"`
	Href          string                `json:"href"`
	Attributes    ActivityAttributes    `json:"attributes"`
	Relationships ActivityRelationships `json:"relationships"`
}

// Activities represents a list of activities.
type Activities struct {
	Data []Activity `json:"data"`
	Href string     `json:"href,omitempty"`
	Next string     `json:"next,omitempty"`
}
