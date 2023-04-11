package applemusic

import (
	"context"
	"fmt"
)

// LibraryPlaylists represents a list of library playlist.
type LibraryPlaylists struct {
	Data []LibraryPlaylist `json:"data"`
	Next string            `json:"next,omitempty"`
}

type LibraryPlaylistTracks struct {
	Data []Song `json:"data"`
	Next string `json:"next,omitempty"`
}

func (s *MeService) getLibraryPlaylists(ctx context.Context, u string, opt interface{}) (*LibraryPlaylists, *Response, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	libraryPlaylists := &LibraryPlaylists{}
	resp, err := s.client.Do(ctx, req, libraryPlaylists)
	if err != nil {
		return nil, resp, err
	}

	return libraryPlaylists, resp, nil
}

// GetLibraryPlaylist fetches a library playlist using its identifier.
func (s *MeService) GetLibraryPlaylist(ctx context.Context, id string, opt *Options) (*LibraryPlaylists, *Response, error) {
	u := fmt.Sprintf("v1/me/library/playlists/%s", id)

	return s.getLibraryPlaylists(ctx, u, opt)
}

// LibraryPlaylistsByIdsOptions specifies the optional parameters to the
// MeService.GetLibraryPlaylistsByIds method.
type LibraryPlaylistsByIdsOptions struct {
	Ids string `url:"ids"`

	Options
}

// GetLibraryPlaylistsByIds fetches one or more library playlists using their identifiers.
func (s *MeService) GetLibraryPlaylistsByIds(ctx context.Context, opt *LibraryPlaylistsByIdsOptions) (*LibraryPlaylists, *Response, error) {
	u := "v1/me/library/playlists"

	return s.getLibraryPlaylists(ctx, u, opt)
}

// GetAllLibraryPlaylists fetches all the library playlists in alphabetical order.
func (s *MeService) GetAllLibraryPlaylists(ctx context.Context, opt *PageOptions) (*LibraryPlaylists, *Response, error) {
	u := "v1/me/library/playlists"

	return s.getLibraryPlaylists(ctx, u, opt)
}

func (s *MeService) getLibraryPlaylistsTracks(ctx context.Context, u string, opt interface{}) (*LibraryPlaylistTracks, *Response, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	libraryPlaylistTracks := &LibraryPlaylistTracks{}
	resp, err := s.client.Do(ctx, req, libraryPlaylistTracks)
	if err != nil {
		return nil, resp, err
	}

	return libraryPlaylistTracks, resp, nil
}

// GetLibraryPlaylist fetches a library playlist using its identifier.
func (s *MeService) GetLibraryPlaylistTracks(ctx context.Context, id string, opt *Options) ([]Song, error) {
	u := fmt.Sprintf("v1/me/library/playlists/%s/tracks", id)

	var tracks []Song
	for len(u) > 0 {
		lpt, _, err := s.getLibraryPlaylistsTracks(ctx, u, opt)
		if err != nil {
			return tracks, err
		}

		tracks = append(tracks, lpt.Data...)
		u = lpt.Next
	}

	return tracks, nil
}

type libraryPlaylistCatalogRelationships struct {
	Tracks LibraryPlaylistTracks `json:"tracks"`
}

type libraryPlaylistCatalog struct {
	Id            string                              `json:"id"`
	Type          string                              `json:"type"`
	Href          string                              `json:"href"`
	Attributes    LibraryPlaylistAttributes           `json:"attributes"`
	Relationships libraryPlaylistCatalogRelationships `json:"relationships"`
}

type libraryPlaylistCatalogResponse struct {
	Data []libraryPlaylistCatalog `json:"data"`
	Next string                   `json:"next,omitempty"`
}

func (s *MeService) getLibraryPlaylistCatalogTracks(ctx context.Context, u string, opt interface{}) (*LibraryPlaylistTracks, *Response, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	libraryPlaylistResponse := &libraryPlaylistCatalogResponse{}
	resp, err := s.client.Do(ctx, req, &libraryPlaylistResponse)
	if err != nil {
		return nil, resp, err
	}
	if libraryPlaylistResponse != nil && len(libraryPlaylistResponse.Data) > 0 {
		return &libraryPlaylistResponse.Data[0].Relationships.Tracks, resp, nil
	}
	return nil, resp, nil

}

// GetLibraryPlaylistCatalogTracks fetches a library playlist using its identifier to get the catalog tracks of the playlist.
func (s *MeService) GetLibraryPlaylistCatalogTracks(ctx context.Context, id string, opt *Options) ([]Song, error) {
	u := fmt.Sprintf("v1/me/library/playlists/%s/catalog", id)

	var tracks []Song
	for len(u) > 0 {
		opt.Include = "tracks"
		lpt, _, err := s.getLibraryPlaylistCatalogTracks(ctx, u, opt)
		if err != nil {
			return tracks, err
		}

		tracks = append(tracks, lpt.Data...)
		u = lpt.Next
	}

	return tracks, nil
}

type CreateLibraryPlaylistAttributes struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateLibraryPlaylistTrack struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type CreateLibraryPlaylistTrackData struct {
	Data []CreateLibraryPlaylistTrack `json:"data"`
}

type CreateLibraryPlaylistRelationships struct {
	Tracks CreateLibraryPlaylistTrackData `json:"tracks"`
}

type CreateLibraryPlaylist struct {
	Attributes    CreateLibraryPlaylistAttributes     `json:"attributes"`
	Relationships *CreateLibraryPlaylistRelationships `json:"relationships,omitempty"`
}

func (s *MeService) CreateLibraryPlaylist(ctx context.Context, body CreateLibraryPlaylist, opt *Options) (*LibraryPlaylists, *Response, error) {
	u := "v1/me/library/playlists"

	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("POST", u, body)
	if err != nil {
		return nil, nil, err
	}

	libraryPlaylists := &LibraryPlaylists{}
	resp, err := s.client.Do(ctx, req, libraryPlaylists)
	if err != nil {
		return nil, resp, err
	}

	return libraryPlaylists, resp, nil
}

func (s *MeService) AddLibraryTracksToPlaylist(ctx context.Context, playlistId string, body CreateLibraryPlaylistTrackData) (*Response, error) {
	u := fmt.Sprintf("/v1/me/library/playlists/%s/tracks", playlistId)

	req, err := s.client.NewRequest("POST", u, body)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
