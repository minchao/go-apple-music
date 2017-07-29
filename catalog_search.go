package applemusic

import (
	"context"
	"fmt"
)

// SearchResults represents a results, that contains a map of search results.
// The members of the results object are the types of resources and the value for each is a Response Root object.
type SearchResults struct {
	Activities    *Activities    `json:"activities,omitempty"`
	Albums        *Albums        `json:"albums,omitempty"`
	AppleCurators *AppleCurators `json:"apple-curators,omitempty"`
	Artists       *Artists       `json:"artists,omitempty"`
	Curators      *Curators      `json:"curators,omitempty"`
	MusicVideos   *MusicVideos   `json:"music-videos,omitempty"`
	Playlists     *Playlists     `json:"playlists,omitempty"`
	Stations      *Stations      `json:"stations,omitempty"`
	Songs         *Songs         `json:"songs,omitempty"`
}

// Search represents the result of search for resources.
type Search struct {
	Results SearchResults `json:"results"`
}

// SearchOptions specifies the parameters to search the catalog.
type SearchOptions struct {
	// The entered text for the search with ‘+’ characters between each word,
	// to replace spaces (for example term=james+br).
	Term string `url:"term"`

	// (Optional) The localization to use, specified by a language tag.
	// The possible values are in the supportedLanguageTags array belonging to the Storefront object specified by storefront.
	// Otherwise, the storefront’s defaultLanguageTag is used.
	Language string `url:"l,omitempty"`

	// (Optional) The limit on the number of objects, or number of objects in the specified relationship, that are returned.
	// The default value is 5 and the maximum value is 25.
	Limit int `url:"limit,omitempty"`

	// (Optional; valid only with types specified) The next page or group of objects to fetch.
	Offset int `url:"offset,omitempty"`

	// (Optional) The list of the types of resources to include in the results.
	Types string `url:"types,omitempty"`
}

// Search searches the catalog using a query.
func (s *CatalogService) Search(ctx context.Context, storefront string, opt *SearchOptions) (*Search, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/search", storefront)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	search := &Search{}
	resp, err := s.client.Do(ctx, req, search)
	if err != nil {
		return nil, resp, err
	}

	return search, resp, nil
}

// SearchHintsOptions specifies the parameters to search hints.
type SearchHintsOptions struct {
	Term     string `url:"term"`
	Language string `url:"l,omitempty"`
	Limit    int    `url:"limit,omitempty"` // (Optional) The number of search terms to be returned. The default value is 10.
	Types    string `url:"types,omitempty"`
}

// SearchHintsResults represents a results, that contains terms array.
type SearchHintsResults struct {
	Terms []string `json:"terms"`
}

// SearchHints represents the result of search hints.
type SearchHints struct {
	Results SearchHintsResults `json:"results"`
}

// SearchHints fetches the search term results for a hint.
func (s *CatalogService) SearchHints(ctx context.Context, storefront string, opt *SearchHintsOptions) (*SearchHints, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/search/hints", storefront)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	searchHints := &SearchHints{}
	resp, err := s.client.Do(ctx, req, searchHints)
	if err != nil {
		return nil, resp, err
	}

	return searchHints, resp, nil
}
