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
