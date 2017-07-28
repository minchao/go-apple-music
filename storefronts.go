package applemusic

import (
	"context"
	"fmt"
	"strings"
)

// StorefrontsService handles communication with the storefront related methods of the Apple Music API.
type StorefrontsService service

// StorefrontAttributes represents a to-one or to-many relationship from one resource object to others.
type StorefrontAttributes struct {
	DefaultLanguageTag    string   `json:"defaultLanguageTag"`
	Name                  string   `json:"name"`
	SupportedLanguageTags []string `json:"supportedLanguageTags"`
}

// Storefront represents a storefront, an iTunes Store territory that the content is available in.
type Storefront struct {
	Id         string               `json:"id"`
	Type       string               `json:"type"`
	Href       string               `json:"href"`
	Attributes StorefrontAttributes `json:"attributes"`
}

// Storefronts represents a list of storefronts.
type Storefronts struct {
	Data []Storefront `json:"data"`
	Next string       `json:"next,omitempty"`
}

func (s *StorefrontsService) get(ctx context.Context, u string) (*Storefronts, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	storefronts := &Storefronts{}
	resp, err := s.client.Do(ctx, req, storefronts)
	if err != nil {
		return nil, resp, err
	}

	return storefronts, resp, nil
}

// Get fetches a single storefront using its identifier.
func (s *StorefrontsService) Get(ctx context.Context, id string, opt *Options) (*Storefronts, *Response, error) {
	u := fmt.Sprintf("v1/storefronts/%s", id)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.get(ctx, u)
}

// GetByIds fetches multiple storefronts by ids.
func (s *StorefrontsService) GetByIds(ctx context.Context, ids []string, opt *Options) (*Storefronts, *Response, error) {
	u := fmt.Sprintf("v1/storefronts?ids=%s", strings.Join(ids, ","))
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.get(ctx, u)
}

// GetAll fetches all the storefronts in alphabetical order.
func (s *StorefrontsService) GetAll(ctx context.Context, opt *PageOptions) (*Storefronts, *Response, error) {
	u := "v1/storefronts"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.get(ctx, u)
}
