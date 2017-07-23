package applemusic

import (
	"context"
	"fmt"
	"strings"
)

type GenreAttributes struct {
	Name string `json:"name"`
}

// Genre represents a genre for resources.
type Genre struct {
	Id         string          `json:"id"`
	Type       string          `json:"type"`
	Href       string          `json:"href"`
	Attributes GenreAttributes `json:"attributes"`
}

// Genres represents a list of genres.
type Genres struct {
	Data []Genre `json:"data"`
	Href string  `json:"href,omitempty"`
	Next string  `json:"next,omitempty"`
}

func (s *CatalogService) getGenres(ctx context.Context, u string) (*Genres, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	genres := &Genres{}
	resp, err := s.client.Do(ctx, req, genres)
	if err != nil {
		return nil, resp, err
	}

	return genres, resp, nil
}

// GetGenre fetches a genre using its identifier.
func (s *CatalogService) GetGenre(ctx context.Context, storefront, id string, opt *Options) (*Genres, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/genres/%s", storefront, id)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getGenres(ctx, u)
}

// GetGenresByIds fetches one or more genres.
func (s *CatalogService) GetGenresByIds(ctx context.Context, storefront string, ids []string, opt *Options) (*Genres, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/genres?ids=%s", storefront, strings.Join(ids, ","))
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getGenres(ctx, u)
}

// GetAllGenres fetches all genres for the current top charts.
func (s *CatalogService) GetAllGenres(ctx context.Context, storefront string, opt *PageOptions) (*Genres, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/genres", storefront)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getGenres(ctx, u)
}
