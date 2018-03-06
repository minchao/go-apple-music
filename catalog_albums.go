package applemusic

import (
	"context"
	"fmt"
	"strings"
)

// AlbumAttributes represents the attributes of the resource.
type AlbumAttributes struct {
	ArtistName     string          `json:"artistName"`
	Artwork        Artwork         `json:"artwork"`
	ContentRating  string          `json:"contentRating,omitempty"`
	Copyright      string          `json:"copyright"`
	EditorialNotes *EditorialNotes `json:"editorialNotes,omitempty"`
	GenreNames     []string        `json:"genreNames"`
	IsComplete     bool            `json:"isComplete"`
	IsSingle       bool            `json:"isSingle"`
	Name           string          `json:"name"`
	RecordLabel    string          `json:"recordLabel"`
	ReleaseDate    string          `json:"releaseDate"`
	PlayParams     *PlayParameters `json:"playParams,omitempty"`
	TrackCount     int64           `json:"trackCount"`
	URL            string          `json:"url"`
}

// AlbumRelationships represents a to-one or to-many relationship from one resource object to others.
type AlbumRelationships struct {
	Artists Artists `json:"artists"`          // Default inclusion: Identifiers only
	Genres  *Genres `json:"genres,omitempty"` // Default inclusion: None
	Tracks  Tracks  `json:"tracks"`           // The songs and music videos on the album. Default inclusion: Objects
}

// Album represents an album.
type Album struct {
	Id            string             `json:"id"`
	Type          string             `json:"type"`
	Href          string             `json:"href"`
	Attributes    AlbumAttributes    `json:"attributes"`
	Relationships AlbumRelationships `json:"relationships"`
}

// Albums represents a list of albums.
type Albums struct {
	Data []Album `json:"data"`
	Href string  `json:"href,omitempty"`
	Next string  `json:"next,omitempty"`
}

func (s *CatalogService) getAlbums(ctx context.Context, u string) (*Albums, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	albums := &Albums{}
	resp, err := s.client.Do(ctx, req, albums)
	if err != nil {
		return nil, resp, err
	}

	return albums, resp, nil
}

// GetAlbum fetches an album using its identifier.
func (s *CatalogService) GetAlbum(ctx context.Context, storefront, id string, opt *Options) (*Albums, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/albums/%s", storefront, id)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getAlbums(ctx, u)
}

// GetAlbumsByIds fetches one or more albums using their identifiers.
func (s *CatalogService) GetAlbumsByIds(ctx context.Context, storefront string, ids []string, opt *Options) (*Albums, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/albums?ids=%s", storefront, strings.Join(ids, ","))
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getAlbums(ctx, u)
}
