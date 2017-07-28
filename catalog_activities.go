package applemusic

import (
	"context"
	"fmt"
	"strings"
)

type ActivityAttributes struct {
	URL            string          `json:"url"`
	Name           string          `json:"name"`
	Artwork        Artwork         `json:"artwork"`
	EditorialNotes *EditorialNotes `json:"editorialNotes,omitempty"`
}

type ActivityRelationships struct {
	Playlists Playlists `json:"playlists"` // Default inclusion: Identifiers only
}

// Activity represents an activity.
type Activity struct {
	Id            string                `json:"id"`
	Type          string                `json:"type"`
	Href          string                `json:"href"`
	Attributes    ActivityAttributes    `json:"attributes"`
	Relationships ActivityRelationships `json:"relationships"`
}

// Activities represents a list of activities.
type Activities struct {
	Data []Activity `json:"data"`
	Href string     `json:"href,omitempty"`
	Next string     `json:"next,omitempty"`
}

func (s *CatalogService) getActivities(ctx context.Context, u string) (*Activities, *Response, error) {
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	activities := &Activities{}
	resp, err := s.client.Do(ctx, req, activities)
	if err != nil {
		return nil, resp, err
	}

	return activities, resp, nil
}

// GetActivity fetches an activity using its identifier.
func (s *CatalogService) GetActivity(ctx context.Context, storefront, id string, opt *Options) (*Activities, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/activities/%s", storefront, id)
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getActivities(ctx, u)
}

// GetActivitiesByIds fetches one or more activities using their identifiers.
func (s *CatalogService) GetActivitiesByIds(ctx context.Context, storefront string, ids []string, opt *Options) (*Activities, *Response, error) {
	u := fmt.Sprintf("v1/catalog/%s/activities?ids=%s", storefront, strings.Join(ids, ","))
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

	return s.getActivities(ctx, u)
}
