package applemusic

import "context"

// LibraryMusicVideoAttributes represents the attributes of library music video object.
type LibraryMusicVideoAttributes struct {
	AlbumName        string         `json:"albumName"`
	ArtistName       string         `json:"artistName"`
	Artwork          Artwork        `json:"artwork"`
	ContentRating    string         `json:"contentRating,omitempty"`
	DurationInMillis int64          `json:"durationInMillis,omitempty"`
	Name             string         `json:"name"`
	PlayParams       PlayParameters `json:"playParams,omitempty"`
	TrackNumber      int            `json:"trackNumber,omitempty"`
}

// LibraryMusicVideo represents a Resource object that represents a library music video.
type LibraryMusicVideo struct {
	Id         string                      `json:"id"`
	Type       string                      `json:"type"`
	Href       string                      `json:"href,omitempty"`
	Attributes LibraryMusicVideoAttributes `json:"attributes,omitempty"`
}

// LibraryMusicVideos represents a list of library music video.
type LibraryMusicVideos struct {
	Data []LibraryMusicVideo `json:"data"`
	Href string              `json:"href,omitempty"`
	Next string              `json:"next,omitempty"`
}

func (s *MeService) getLibraryMusicVideos(ctx context.Context, u string, opt interface{}) (*LibraryMusicVideos, *Response, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	libraryMusicVideos := &LibraryMusicVideos{}
	resp, err := s.client.Do(ctx, req, libraryMusicVideos)
	if err != nil {
		return nil, resp, err
	}

	return libraryMusicVideos, resp, nil
}

// GetAllLibraryMusicVideos fetches all the library music videos in alphabetical order.
func (s *MeService) GetAllLibraryMusicVideos(ctx context.Context, opt *PageOptions) (*LibraryMusicVideos, *Response, error) {
	u := "v1/me/library/music-videos"

	return s.getLibraryMusicVideos(ctx, u, opt)
}
