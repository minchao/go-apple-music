package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestCatalogService_GetActivity(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/activities/976439514", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(activitiesJSON)
	})

	got, _, err := client.Catalog.GetActivity(context.Background(), "us", "976439514", nil)
	if err != nil {
		t.Errorf("Catalog.GetActivity returned error: %v", err)
	}
	if want := activities; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.GetActivity = %+v, want %+v", got, want)
	}
}

func TestCatalogService_GetActivitiesByIds(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/activities", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"ids": "976439514,976439503",
		})

		w.WriteHeader(http.StatusOK)
	})

	_, _, err := client.Catalog.GetActivitiesByIds(context.Background(), "us", []string{"976439514", "976439503"}, nil)
	if err != nil {
		t.Errorf("Catalog.GetActivitiesByIds returned error: %v", err)
	}
}

var activitiesJSON = []byte(`{
    "data": [
        {
            "id": "976439514",
            "type": "activities",
            "href": "/v1/catalog/us/activities/976439514",
            "attributes": {
                "url": "https://itunes.apple.com/us/activity/party/id976439514",
                "name": "Party",
                "artwork": {
                    "width": 1080,
                    "height": 1080,
                    "url": "https://is1-ssl.mzstatic.com/image/thumb/Features117/v4/99/ed/53/99ed53f5-34b7-6419-b167-cd5a8a3caed2/source/{w}x{h}bb.jpg",
                    "bgColor": "fff3ab",
                    "textColor1": "541621",
                    "textColor2": "3d2d09",
                    "textColor3": "76423c",
                    "textColor4": "645529"
                },
                "editorialNotes": {
                    "short": "Get it started."
                }
            },
            "relationships": {
                "playlists": {
                    "data": [
                        {
                            "id": "pl.2d4d74790f074233b82d07bfae5c219c",
                            "type": "playlists",
                            "href": "/v1/catalog/us/playlists/pl.2d4d74790f074233b82d07bfae5c219c"
                        }
                    ],
                    "href": "/v1/catalog/us/activities/976439514/playlists",
                    "next": "/v1/catalog/us/activities/976439514/playlists?offset=10"
                }
            }
        }
    ]
}`)

var activities = &Activities{
	Data: []Activity{
		{
			Id:   "976439514",
			Type: "activities",
			Href: "/v1/catalog/us/activities/976439514",
			Attributes: ActivityAttributes{
				URL:  "https://itunes.apple.com/us/activity/party/id976439514",
				Name: "Party",
				Artwork: Artwork{
					Width:      1080,
					Height:     1080,
					URL:        "https://is1-ssl.mzstatic.com/image/thumb/Features117/v4/99/ed/53/99ed53f5-34b7-6419-b167-cd5a8a3caed2/source/{w}x{h}bb.jpg",
					BgColor:    "fff3ab",
					TextColor1: "541621",
					TextColor2: "3d2d09",
					TextColor3: "76423c",
					TextColor4: "645529",
				},
				EditorialNotes: &EditorialNotes{
					Short: "Get it started.",
				},
			},
			Relationships: ActivityRelationships{
				Playlists: Playlists{
					Data: []Playlist{
						{
							Id:   "pl.2d4d74790f074233b82d07bfae5c219c",
							Type: "playlists",
							Href: "/v1/catalog/us/playlists/pl.2d4d74790f074233b82d07bfae5c219c",
						},
					},
					Href: "/v1/catalog/us/activities/976439514/playlists",
					Next: "/v1/catalog/us/activities/976439514/playlists?offset=10",
				},
			},
		},
	},
}
