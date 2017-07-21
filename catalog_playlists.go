package applemusic

type PlaylistType string

const (
	PlaylistTypeUserShared  = PlaylistType("user-shared")
	PlaylistTypeEditorial   = PlaylistType("editorial")
	PlaylistTypeExternal    = PlaylistType("external")
	PlaylistTypePersonalMix = PlaylistType("personal-mix")
)

type PlaylistAttributes struct {
	Artwork          *Artwork        `json:"artwork,omitempty"`
	CuratorName      string          `json:"curatorName,omitempty"`
	Description      *EditorialNotes `json:"description,omitempty"`
	LastModifiedDate string          `json:"lastModifiedDate"`
	Name             string          `json:"name"`
	PlaylistType     PlaylistType    `json:"playlistType"`
	PlayParams       *PlayParameters `json:"playParams,omitempty"`
	URL              string          `json:"url"`
}

type PlaylistRelationships struct {
}

// Playlist represents a playlist.
type Playlist struct {
	Id            string                `json:"id"`
	Type          string                `json:"type"`
	Href          string                `json:"href"`
	Attributes    PlaylistAttributes    `json:"attributes"`
	Relationships PlaylistRelationships `json:"relationships"`
}

type Playlists struct {
	Data []Playlist `json:"data"`
}
