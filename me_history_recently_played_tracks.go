package applemusic

import "context"

// HistoryRecentlyPlayedTrackAttributes represents the attributes of history recently played song.
type HistoryRecentlyPlayedTrackAttributes struct {
	AlbumName            string         `json:"albumName"`
	ArtistName           string         `json:"artistName"`
	ArtistURL            string         `json:"artistUrl,omitempty"`
	Artwork              Artwork        `json:"artwork"`
	DiscNumber           int            `json:"discNumber"`
	DurationInMillis     int64          `json:"durationInMillis,omitempty"`
	Name                 string         `json:"name"`
	PlayParams           PlayParameters `json:"playParams,omitempty"`
	TrackNumber          int            `json:"trackNumber"`
	GenreNames           []string       `json:"genreNames"`
	ReleaseDate          string         `json:"releaseDate"`
	ISRC                 string         `json:"isrc"`
	URL                  string         `json:"url"`
	HasLyrics            bool           `json:"hasLyrics"`
	IsAppleDigitalMaster bool           `json:"isAppleDigitalMaster"`
	Previews             []Preview      `json:"previews,omitempty"`
}

// HistoryRecentlyPlayedTrack represents a Resource object that represents a history recently played song.
type HistoryRecentlyPlayedTrack struct {
	Id         string                               `json:"id"`
	Type       string                               `json:"type"`
	Href       string                               `json:"href,omitempty"`
	Attributes HistoryRecentlyPlayedTrackAttributes `json:"attributes,omitempty"`
}

// HistoryRecentlyPlayedTracks represents a list of history recently played songs.
type HistoryRecentlyPlayedTracks struct {
	Data []HistoryRecentlyPlayedTrack `json:"data"`
	Href string                       `json:"href,omitempty"`
	Next string                       `json:"next,omitempty"`
}

func (s *MeService) getHistoryRecentlyPlayedTracks(ctx context.Context, u string, opt interface{}) (*HistoryRecentlyPlayedTracks, *Response, error) {
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	historyRecentlyPlayedTracks := &HistoryRecentlyPlayedTracks{}
	resp, err := s.client.Do(ctx, req, historyRecentlyPlayedTracks)
	if err != nil {
		return nil, resp, err
	}

	return historyRecentlyPlayedTracks, resp, nil
}

// GetHistoryRecentlyPlayedTracks fetches all history recently played songs.
func (s *MeService) GetHistoryRecentlyPlayedTracks(ctx context.Context, opt *PageOptions) (*HistoryRecentlyPlayedTracks, *Response, error) {
	u := "v1/me/recent/played/tracks"

	return s.getHistoryRecentlyPlayedTracks(ctx, u, opt)
}
