package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestCatalogService_GetCurator(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/curators/1107687517", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(curatorsJSON)
	})

	got, _, err := client.Catalog.GetCurator(context.Background(), "us", "1107687517", nil)
	if err != nil {
		t.Errorf("Catalog.GetCurator returned error: %v", err)
	}
	if want := curators; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.GetCurator = %+v, want %+v", got, want)
	}
}

func TestCatalogService_GetCuratorsByIds(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/curators", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"ids": "976439448,1107687517",
		})

		w.WriteHeader(http.StatusOK)
	})

	_, _, err := client.Catalog.GetCuratorsByIds(context.Background(), "us", []string{"976439448", "1107687517"}, nil)
	if err != nil {
		t.Errorf("Catalog.GetCuratorsByIds returned error: %v", err)
	}
}

var curatorsJSON = []byte(`{
    "data": [
        {
            "id": "1107687517",
            "type": "curators",
            "href": "/v1/catalog/us/curators/1107687517",
            "attributes": {
                "url": "https://itunes.apple.com/us/curator/largeup/id1107687517",
                "name": "LargeUp",
                "artwork": {
                    "width": 1080,
                    "height": 1080,
                    "url": "https://is2-ssl.mzstatic.com/image/thumb/Features60/v4/86/14/c5/8614c510-11a5-ef0e-4299-2edbd3987a73/source/{w}x{h}bb.jpg",
                    "bgColor": "161113",
                    "textColor1": "ffffff",
                    "textColor2": "f74b2f",
                    "textColor3": "d0cfcf",
                    "textColor4": "ca3f29"
                },
                "editorialNotes": {
                    "standard": "LargeUp is the global platform for Caribbean music and culture. Since 2009, LargeUp.com has captured the vibrant sounds, styles and flavors of the islands, spotlighting the best in reggae, dancehall, soca + beyond.",
                    "short": "Island music and culture."
                }
            },
            "relationships": {
                "playlists": {
                    "data": [
                        {
                            "id": "pl.b80bbeb133bb42aa8adf17287b0bc31b",
                            "type": "playlists",
                            "href": "/v1/catalog/us/playlists/pl.b80bbeb133bb42aa8adf17287b0bc31b"
                        }
                    ],
                    "href": "/v1/catalog/us/curators/1107687517/playlists",
                    "next": "/v1/catalog/us/curators/1107687517/playlists?offset=10"
                }
            }
        }
    ]
}`)

var curators = &Curators{
	Data: []Curator{
		{
			Id:   "1107687517",
			Type: "curators",
			Href: "/v1/catalog/us/curators/1107687517",
			Attributes: CuratorAttributes{
				URL:  "https://itunes.apple.com/us/curator/largeup/id1107687517",
				Name: "LargeUp",
				Artwork: Artwork{
					Width:      1080,
					Height:     1080,
					URL:        "https://is2-ssl.mzstatic.com/image/thumb/Features60/v4/86/14/c5/8614c510-11a5-ef0e-4299-2edbd3987a73/source/{w}x{h}bb.jpg",
					BgColor:    "161113",
					TextColor1: "ffffff",
					TextColor2: "f74b2f",
					TextColor3: "d0cfcf",
					TextColor4: "ca3f29",
				},
				EditorialNotes: &EditorialNotes{
					Standard: "LargeUp is the global platform for Caribbean music and culture. Since 2009, LargeUp.com has captured the vibrant sounds, styles and flavors of the islands, spotlighting the best in reggae, dancehall, soca + beyond.",
					Short:    "Island music and culture.",
				},
			},
			Relationships: CuratorRelationships{
				Playlists: Playlists{
					Data: []Playlist{
						{
							Id:   "pl.b80bbeb133bb42aa8adf17287b0bc31b",
							Type: "playlists",
							Href: "/v1/catalog/us/playlists/pl.b80bbeb133bb42aa8adf17287b0bc31b",
						},
					},
					Href: "/v1/catalog/us/curators/1107687517/playlists",
					Next: "/v1/catalog/us/curators/1107687517/playlists?offset=10",
				},
			},
		},
	},
}
