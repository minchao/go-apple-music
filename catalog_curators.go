package applemusic

import (
	"context"
	"fmt"
	"strings"
)

// CuratorAttributes represents the attributes of the resource.
type CuratorAttributes struct {
	URL            string          `json:"url"`
	Name           string          `json:"name"`
	Artwork        Artwork         `json:"artwork"`
	EditorialNotes *EditorialNotes `json:"editorialNotes,omitempty"`
}

// CuratorRelationships represents a to-one or to-many relationship from one resource object to others.
type CuratorRelationships struct {
	Playlists Playlists `json:"playlists"` // Default inclusion: Identifiers only
}

// Curator represents a curator of resources.
type Curator struct {
	Id            string               `json:"id"`
	Type          string               `json:"type"`
	Href          string               `json:"href"`
	Attributes    CuratorAttributes    `json:"attributes"`
	Relationships CuratorRelationships `json:"relationships"`
}

// Curators represents a list of curators.
type Curators struct {
	Data []Curator `json:"data"`
	Href string    `json:"href,omitempty"`
	Next string    `json:"next,omitempty"`
}

func (s *CatalogService) getCurators(ctx context.Context, u string) (*Curators, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	curators := &Curators{}
	resp, err := s.client.Do(ctx, req, curators)
	if err != nil {
		return nil, resp, err
	}

	return curators, resp, nil
}

// GetCurator fetches an curator using its identifier.
func (s *CatalogService) GetCurator(ctx context.Context, storefront, id string, opt *Options) (*Curators, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/curators/%s", storefront, id)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getCurators(ctx, u)
}

// GetCuratorsByIds fetches one or more curators using their identifiers.
func (s *CatalogService) GetCuratorsByIds(ctx context.Context, storefront string, ids []string, opt *Options) (*Curators, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/curators?ids=%s", storefront, strings.Join(ids, ","))
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getCurators(ctx, u)
}
