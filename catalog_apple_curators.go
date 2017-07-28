package applemusic

import (
	"context"
	"fmt"
	"strings"
)

// AppleCuratorAttributes represents the attributes of the resource.
type AppleCuratorAttributes struct {
	URL            string          `json:"url"`
	Name           string          `json:"name"`
	Artwork        Artwork         `json:"artwork"`
	EditorialNotes *EditorialNotes `json:"editorialNotes,omitempty"`
}

// AppleCuratorRelationships represents a to-one or to-many relationship from one resource object to others.
type AppleCuratorRelationships struct {
	Playlists Playlists `json:"playlists"` // Default inclusion: Identifiers only
}

// AppleCurator represents an Apple curator.
type AppleCurator struct {
	Id            string                    `json:"id"`
	Type          string                    `json:"type"`
	Href          string                    `json:"href"`
	Attributes    AppleCuratorAttributes    `json:"attributes"`
	Relationships AppleCuratorRelationships `json:"relationships"`
}

// AppleCurators represents a list of apple curators.
type AppleCurators struct {
	Data []AppleCurator `json:"data"`
	Href string         `json:"href,omitempty"`
	Next string         `json:"next,omitempty"`
}

func (s *CatalogService) getAppleCurators(ctx context.Context, u string) (*AppleCurators, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	appleCurators := &AppleCurators{}
	resp, err := s.client.Do(ctx, req, appleCurators)
	if err != nil {
		return nil, resp, err
	}

	return appleCurators, resp, nil
}

// GetAppleCurator fetches an apple curator using its identifier.
func (s *CatalogService) GetAppleCurator(ctx context.Context, storefront, id string, opt *Options) (*AppleCurators, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/apple-curators/%s", storefront, id)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getAppleCurators(ctx, u)
}

// GetAppleCuratorsByIds fetches one or more apple curators using their identifiers.
func (s *CatalogService) GetAppleCuratorsByIds(ctx context.Context, storefront string, ids []string, opt *Options) (*AppleCurators, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/apple-curators?ids=%s", storefront, strings.Join(ids, ","))
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getAppleCurators(ctx, u)
}
