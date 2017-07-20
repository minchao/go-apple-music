package applemusic

type ArtistAttributes struct {
	GenreNames     []string       `json:"genreNames"`
	EditorialNotes EditorialNotes `json:"editorialNotes,omitempty"`
	Name           string         `json:"name"`
	URL            string         `json:"url"`
}

type ArtistRelationships struct {
	Albums      Albums      `json:"albums"`
	Genres      Genres      `json:"genres"`
	MusicVideos MusicVideos `json:"music-videos"`
	Playlists   Playlists   `json:"playlists"`
}

// Artist represents an artist of an album.
type Artist struct {
	Id            string              `json:"id"`
	Type          string              `json:"type"`
	Href          string              `json:"href"`
	Attributes    ArtistAttributes    `json:"attributes,omitempty"`
	Relationships ArtistRelationships `json:"relationships,omitempty"`
}

type Artists struct {
	Data []Artist `json:"data"`
	Href string   `json:"href,omitempty"`
}
