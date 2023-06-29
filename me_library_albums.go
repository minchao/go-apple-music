package applemusic

import "context"

// LibraryAlbumAttributes represents the attributes of library album object.
type LibraryAlbumAttributes struct {
	Name       string         `json:"name"`
	ArtistName string         `json:"artistName"`
	TrackCount int            `json:"trackCount"`
	GenreNames []string       `json:"genreNames"`
	Artwork    Artwork        `json:"artwork"`
	PlayParams PlayParameters `json:"playParams,omitempty"`
	DateAdded  string         `json:"dateAdded"`
}

type LibraryAlbumRelationships struct {
	Catalog Albums  `json:"catalog"`
	Artists Artists `json:"artists"`
}

// LibraryAlbum represents a Resource object that represents a library album.
type LibraryAlbum struct {
	Id            string                    `json:"id"`
	Type          string                    `json:"type"`
	Href          string                    `json:"href,omitempty"`
	Attributes    LibraryAlbumAttributes    `json:"attributes,omitempty"`
	Relationships LibraryAlbumRelationships `json:"relationships,omitempty"`
}

// LibraryAlbums represents a list of library albums.
type LibraryAlbums struct {
	Data []LibraryAlbum `json:"data"`
	Href string         `json:"href,omitempty"`
	Next string         `json:"next,omitempty"`
}

func (s *MeService) getLibraryAlbums(ctx context.Context, u string, opt interface{}) (*LibraryAlbums, *Response, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	libraryAlbums := &LibraryAlbums{}
	resp, err := s.client.Do(ctx, req, libraryAlbums)
	if err != nil {
		return nil, resp, err
	}

	return libraryAlbums, resp, nil
}

// GetAllLibraryAlbums fetches all the library albums in alphabetical order.
func (s *MeService) GetAllLibraryAlbums(ctx context.Context, opt *PageOptions) (*LibraryAlbums, *Response, error) {
	u := "v1/me/library/albums"

	return s.getLibraryAlbums(ctx, u, opt)
}
