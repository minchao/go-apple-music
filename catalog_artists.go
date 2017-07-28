package applemusic

import (
	"context"
	"fmt"
	"strings"
)

type ArtistAttributes struct {
	GenreNames     []string        `json:"genreNames"`
	EditorialNotes *EditorialNotes `json:"editorialNotes,omitempty"`
	Name           string          `json:"name"`
	URL            string          `json:"url"`
}

type ArtistRelationships struct {
	Albums      Albums       `json:"albums"`
	Genres      *Genres      `json:"genres,omitempty"`
	MusicVideos *MusicVideos `json:"music-videos,omitempty"`
	Playlists   *Playlists   `json:"playlists,omitempty"`
}

// Artist represents an artist of an album.
type Artist struct {
	Id            string              `json:"id"`
	Type          string              `json:"type"`
	Href          string              `json:"href"`
	Attributes    ArtistAttributes    `json:"attributes"`
	Relationships ArtistRelationships `json:"relationships"`
}

// Artists represents a list of artists.
type Artists struct {
	Data []Artist `json:"data"`
	Href string   `json:"href,omitempty"`
	Next string   `json:"next,omitempty"`
}

func (s *CatalogService) getArtists(ctx context.Context, u string) (*Artists, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	artists := &Artists{}
	resp, err := s.client.Do(ctx, req, artists)
	if err != nil {
		return nil, resp, err
	}

	return artists, resp, nil
}

// GetArtist fetches a artist using its identifier.
func (s *CatalogService) GetArtist(ctx context.Context, storefront, id string, opt *Options) (*Artists, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/artists/%s", storefront, id)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getArtists(ctx, u)
}

// GetArtistsByIds fetches one or more artists using their identifiers.
func (s *CatalogService) GetArtistsByIds(ctx context.Context, storefront string, ids []string, opt *Options) (*Artists, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/artists?ids=%s", storefront, strings.Join(ids, ","))
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getArtists(ctx, u)
}
