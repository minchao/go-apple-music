package applemusic

import (
	"context"
	"fmt"
	"strings"
)

// MusicVideoAttributes represents the attributes of the resource.
type MusicVideoAttributes struct {
	URL              string          `json:"url"`
	Name             string          `json:"name"`
	GenreNames       []string        `json:"genreNames"`
	ArtistName       string          `json:"artistName"`
	ReleaseDate      string          `json:"releaseDate"`
	Artwork          Artwork         `json:"artwork"`
	PlayParams       *PlayParameters `json:"playParams,omitempty"`
	DurationInMillis int64           `json:"durationInMillis,omitempty"`
	ContentRating    string          `json:"contentRating,omitempty"`
	EditorialNotes   *EditorialNotes `json:"editorialNotes,omitempty"`
	TrackNumber      int             `json:"trackNumber,omitempty"`
	VideoSubType     string          `json:"videoSubType,omitempty"`
}

// MusicVideoRelationships represents a to-one or to-many relationship from one resource object to others.
type MusicVideoRelationships struct {
	Albums  Albums  `json:"albums"`           // Default inclusion: Identifiers only
	Artists Artists `json:"artists"`          // Default inclusion: Identifiers only
	Genres  *Genres `json:"genres,omitempty"` // Default inclusion: None
}

// MusicVideo represents a music video.
type MusicVideo struct {
	Id            string                  `json:"id"`
	Type          string                  `json:"type"`
	Href          string                  `json:"href"`
	Attributes    MusicVideoAttributes    `json:"attributes"`
	Relationships MusicVideoRelationships `json:"relationships"`
}

// MusicVideos represents a list of music videos.
type MusicVideos struct {
	Data []MusicVideo `json:"data"`
	Href string       `json:"href,omitempty"`
	Next string       `json:"next,omitempty"`
}

func (s *CatalogService) getMusicVideos(ctx context.Context, u string) (*MusicVideos, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	musicVideos := &MusicVideos{}
	resp, err := s.client.Do(ctx, req, musicVideos)
	if err != nil {
		return nil, resp, err
	}

	return musicVideos, resp, nil
}

// GetMusicVideo fetches a music video using its identifier.
func (s *CatalogService) GetMusicVideo(ctx context.Context, storefront, id string, opt *Options) (*MusicVideos, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/music-videos/%s", storefront, id)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getMusicVideos(ctx, u)
}

// GetMusicVideosByIds fetches one or more music videos using their identifiers.
func (s *CatalogService) GetMusicVideosByIds(ctx context.Context, storefront string, ids []string, opt *Options) (*MusicVideos, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/music-videos?ids=%s", storefront, strings.Join(ids, ","))
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getMusicVideos(ctx, u)
}
