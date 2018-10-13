package applemusic

import (
	"context"
	"fmt"
)

// ChartAlbums represents a chart of albums.
type ChartAlbums struct {
	Name  string `json:"name"`
	Chart string `json:"chart"`
	Albums
}

// ChartSongs represents a chart of songs.
type ChartSongs struct {
	Name  string `json:"name"`
	Chart string `json:"chart"`
	Songs
}

// ChartMusicVideos represents a chart of music videos.
type ChartMusicVideos struct {
	Name  string `json:"name"`
	Chart string `json:"chart"`
	MusicVideos
}

// ChartPlaylists represent a chart of playlists.
type ChartPlaylists struct {
	Name  string `json:"name"`
	Chart string `json:"chart"`
	Playlists
}

// ChartResults represents a results, that contains chart collections.
type ChartResults struct {
	Albums      *[]ChartAlbums      `json:"albums,omitempty"`
	Songs       *[]ChartSongs       `json:"songs,omitempty"`
	MusicVideos *[]ChartMusicVideos `json:"music-videos,omitempty"`
	Playlists   *[]ChartPlaylists   `json:"playlists,omitempty"`
}

// Charts represents the result of one or more charts.
type Charts struct {
	Results ChartResults `json:"results"`
}

// ChartsOptions specifies the parameters to fetch charts.
type ChartsOptions struct {
	// A list of the types of charts to include in the results.
	// The possible values are albums, songs, and music-videos.
	Types string `url:"types"`

	// (Optional) The localization to use, specified by a language tag.
	// The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront.
	// Otherwise, the storefront’s defaultLanguageTag is used.
	Language string `url:"l,omitempty"`

	// (Optional) The chart to fetch for the specified types.
	// For possible values, get all the charts by sending this endpoint without the chart parameter.
	// The possible values for this parameter are the chart attributes of the Chart objects in the response.
	Chart string `url:"chart,omitempty"`

	// (Optional) The identifier for the genre to use in the chart results. To get the genre identifiers.
	Genre string `url:"genre,omitempty"`

	// (Optional) The number of resources to include per chart.
	// The default value is 20 and the maximum value is 50.
	Limit int `url:"limit,omitempty"`

	// (Optional; only with chart specified) The next page or group of objects to fetch.
	Offset int `url:"offset,omitempty"`
}

// GetAllCharts fetches one or more charts.
func (s *CatalogService) GetAllCharts(ctx context.Context, storefront string, opt *ChartsOptions) (*Charts, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/charts", storefront)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	charts := &Charts{}
	resp, err := s.client.Do(ctx, req, charts)
	if err != nil {
		return nil, resp, err
	}

	return charts, resp, nil
}
