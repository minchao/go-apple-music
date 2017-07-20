package applemusic

type PlaylistType string

const (
	PlaylistTypeUserShared  = PlaylistType("user-shared")
	PlaylistTypeEditorial   = PlaylistType("editorial")
	PlaylistTypeExternal    = PlaylistType("external")
	PlaylistTypePersonalMix = PlaylistType("personal-mix")
)

// Playlist represents a playlist.
type Playlist struct {
	Artwork          Artwork        `json:"artwork"`
	CuratorName      string         `json:"curatorName"`
	Description      EditorialNotes `json:"description"`
	LastModifiedDate string         `json:"lastModifiedDate"`
	Name             string         `json:"name"`
	PlaylistType     PlaylistType   `json:"playlistType"`
	PlayParams       PlayParameters `json:"playParams"`
	URL              string         `json:"url"`
}

type Playlists struct {
	Data []Playlist `json:"data"`
}
