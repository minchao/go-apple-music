package applemusic

import (
	"context"
)

// HistoryHeavyRotation represents a list of heavy rotation content.
type HistoryHeavyRotation struct {
	Data []Resource `json:"data"`
	Href string     `json:"href,omitempty"`
	Next string     `json:"next,omitempty"`
}

// GetHistoryHeavyRotation fetches the resources in heavy rotation for the user.
func (s *MeService) GetHistoryHeavyRotation(ctx context.Context, opt *PageOptions) (*HistoryHeavyRotation, *Response, error) {
	u := "v1/me/history/heavy-rotation"
	u, err := addOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	heavyRotation := &HistoryHeavyRotation{}
	resp, err := s.client.Do(ctx, req, heavyRotation)
	if err != nil {
		return nil, resp, err
	}

	return heavyRotation, resp, nil
}
