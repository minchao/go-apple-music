package applemusic

import "context"

// GetStorefront fetches a user’s storefront.
func (s *MeService) GetStorefront(ctx context.Context, opt *PageOptions) (*Storefronts, *Response, error) {
	u := "v1/me/storefront"
	u, err := addOptions(u, opt)
	if err != nil {
		return nil, nil, err
	}

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
