package applemusic

import "context"

// LibrarySongAttributes represents the attributes of library song object.
type LibrarySongAttributes struct {
	AlbumName        string         `json:"albumName"`
	ArtistName       string         `json:"artistName"`
	Artwork          Artwork        `json:"artwork"`
	ContentRating    string         `json:"contentRating,omitempty"`
	DiscNumber       int            `json:"discNumber"`
	DurationInMillis int64          `json:"durationInMillis,omitempty"`
	Name             string         `json:"name"`
	PlayParams       PlayParameters `json:"playParams,omitempty"`
	TrackNumber      int            `json:"trackNumber"`
}

// LibrarySong represents a Resource object that represents a library song.
type LibrarySong struct {
	Id         string                `json:"id"`
	Type       string                `json:"type"`
	Href       string                `json:"href,omitempty"`
	Attributes LibrarySongAttributes `json:"attributes,omitempty"`
}

// LibrarySongs represents a list of library songs.
type LibrarySongs struct {
	Data []LibrarySong `json:"data"`
	Href string        `json:"href,omitempty"`
	Next string        `json:"next,omitempty"`
}

func (s *MeService) getLibrarySongs(ctx context.Context, u string, opt interface{}) (*LibrarySongs, *Response, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	librarySongs := &LibrarySongs{}
	resp, err := s.client.Do(ctx, req, librarySongs)
	if err != nil {
		return nil, resp, err
	}

	return librarySongs, resp, nil
}

// GetAllLibrarySongs fetches all the library songs in alphabetical order.
func (s *MeService) GetAllLibrarySongs(ctx context.Context, opt *PageOptions) (*LibrarySongs, *Response, error) {
	u := "v1/me/library/songs"

	return s.getLibrarySongs(ctx, u, opt)
}
