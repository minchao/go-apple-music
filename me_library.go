package applemusic

// LibraryPlaylistAttributes represents the attributes of library playlist.
type LibraryPlaylistAttributes struct {
	Artwork     *Artwork        `json:"artwork"`
	Description *EditorialNotes `json:"description"`
	Name        string          `json:"name"`
	PlayParams  *PlayParameters `json:"playParams"`
	CanDelete   bool            `json:"canDelete"`
	CanEdit     bool            `json:"canEdit"`
	HasCatalog  bool            `json:"hasCatalog"`
}

// LibraryPlaylistRelationships represents a to-one or to-many relationship from one resource object to others.
type LibraryPlaylistRelationships struct {
	Tracks Tracks `json:"tracks"` // The library songs and library music videos included in the playlist. Only available when fetching single library playlist resource by ID.
}

// LibraryPlaylist represents a library playlist.
type LibraryPlaylist struct {
	Id            string                       `json:"id"`
	Type          string                       `json:"type"`
	Href          string                       `json:"href"`
	Attributes    LibraryPlaylistAttributes    `json:"attributes"`
	Relationships LibraryPlaylistRelationships `json:"relationships"`
}
