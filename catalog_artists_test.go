package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestCatalogService_GetArtist(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/artists/178834", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(artistsJSON)
	})

	got, _, err := client.Catalog.GetArtist(context.Background(), "us", "178834", nil)
	if err != nil {
		t.Errorf("Catalog.GetArtist returned error: %v", err)
	}
	if want := artists; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.GetArtist = %+v, want %+v", got, want)
	}
}

func TestCatalogService_GetArtistsByIds(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/artists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"ids": "178834,462006",
		})

		w.WriteHeader(http.StatusOK)
	})

	_, _, err := client.Catalog.GetArtistsByIds(context.Background(), "us", []string{"178834", "462006"}, nil)
	if err != nil {
		t.Errorf("Catalog.GetArtistsByIds returned error: %v", err)
	}
}

var artistsJSON = []byte(`{
  "data": [
    {
      "id": "178834",
      "type": "artists",
      "href": "/v1/catalog/us/artists/178834",
      "attributes": {
        "url": "https://itunes.apple.com/us/artist/bruce-springsteen/id178834",
        "name": "Bruce Springsteen",
        "genreNames": [
          "Rock"
        ]
      },
      "relationships": {
        "albums": {
          "data": [
            {
              "id": "1185902474",
              "type": "albums",
              "href": "/v1/catalog/us/albums/1185902474"
            }
          ],
          "href": "/v1/catalog/us/artists/178834/albums",
          "next": "/v1/catalog/us/artists/178834/albums?offset=25"
        }
      }
    }
  ]
}`)

var artists = &Artists{
	Data: []Artist{
		{
			Id:   "178834",
			Type: "artists",
			Href: "/v1/catalog/us/artists/178834",
			Attributes: ArtistAttributes{
				URL:  "https://itunes.apple.com/us/artist/bruce-springsteen/id178834",
				Name: "Bruce Springsteen",
				GenreNames: []string{
					"Rock",
				},
			},
			Relationships: ArtistRelationships{
				Albums: Albums{
					Data: []Album{
						{
							Id:   "1185902474",
							Type: "albums",
							Href: "/v1/catalog/us/albums/1185902474",
						},
					},
					Href: "/v1/catalog/us/artists/178834/albums",
					Next: "/v1/catalog/us/artists/178834/albums?offset=25",
				},
			},
		},
	},
}
