package applemusic

type MusicVideoAttributes struct {
	URL              string          `json:"url"`
	Name             string          `json:"name"`
	GenreNames       []string        `json:"genreNames"`
	ArtistName       string          `json:"artistName"`
	ReleaseDate      string          `json:"releaseDate"`
	Artwork          Artwork         `json:"artwork"`
	PlayParams       *PlayParameters `json:"playParams,omitempty"`
	DurationInMillis int64           `json:"durationInMillis,omitempty"`
	ContentRating    string          `json:"contentRating,omitempty"`
	EditorialNotes   *EditorialNotes `json:"editorialNotes,omitempty"`
	TrackNumber      int             `json:"trackNumber,omitempty"`
	VideoSubType     string          `json:"videoSubType,omitempty"`
}

type MusicVideoRelationships struct {
	Albums  Albums  `json:"albums"`           // Default inclusion: Identifiers only
	Artists Artists `json:"artists"`          // Default inclusion: Identifiers only
	Genres  *Genres `json:"genres,omitempty"` // Default inclusion: None
}

// MusicVideo represents a music video.
type MusicVideo struct {
	Id            string                  `json:"id"`
	Type          string                  `json:"type"`
	Href          string                  `json:"href"`
	Attributes    MusicVideoAttributes    `json:"attributes"`
	Relationships MusicVideoRelationships `json:"relationships"`
}

type MusicVideos struct {
	Data []MusicVideo `json:"data"`
}
