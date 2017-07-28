package applemusic

import (
	"context"
	"fmt"
	"strings"
)

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
	Curator Curators `json:"curator"` // Default inclusion: Identifiers only
	Tracks  Tracks   `json:"tracks"`  // The songs and music videos included in the playlist. Default inclusion: Objects
}

// Playlist represents a playlist.
type Playlist struct {
	Id            string                `json:"id"`
	Type          string                `json:"type"`
	Href          string                `json:"href"`
	Attributes    PlaylistAttributes    `json:"attributes"`
	Relationships PlaylistRelationships `json:"relationships"`
}

// Playlists represents a list of playlists.
type Playlists struct {
	Data []Playlist `json:"data"`
	Href string     `json:"href,omitempty"`
	Next string     `json:"next,omitempty"`
}

func (s *CatalogService) getPlaylists(ctx context.Context, u string) (*Playlists, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	playlists := &Playlists{}
	resp, err := s.client.Do(ctx, req, playlists)
	if err != nil {
		return nil, resp, err
	}

	return playlists, resp, nil
}

// GetPlaylist fetches a playlist using its identifier.
func (s *CatalogService) GetPlaylist(ctx context.Context, storefront, id string, opt *Options) (*Playlists, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/playlists/%s", storefront, id)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getPlaylists(ctx, u)
}

// GetPlaylistsByIds fetches one or more playlists using their identifiers.
func (s *CatalogService) GetPlaylistsByIds(ctx context.Context, storefront string, ids []string, opt *Options) (*Playlists, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/playlists?ids=%s", storefront, strings.Join(ids, ","))
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getPlaylists(ctx, u)
}
