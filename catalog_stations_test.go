package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestCatalogService_GetStation(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/stations/ra.985484166", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(stationsJSON)
	})

	got, _, err := client.Catalog.GetStation(context.Background(), "us", "ra.985484166", nil)
	if err != nil {
		t.Errorf("Catalog.GetStation returned error: %v", err)
	}
	if want := stations; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.GetStation = %+v, want %+v", got, want)
	}
}

func TestCatalogService_GetStationsByIds(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/stations", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"ids": "ra.985484166,ra.1128062616",
		})

		w.WriteHeader(http.StatusOK)
	})

	_, _, err := client.Catalog.GetStationsByIds(context.Background(), "us", []string{"ra.985484166", "ra.1128062616"}, nil)
	if err != nil {
		t.Errorf("Catalog.GetStationsByIds returned error: %v", err)
	}
}

var stationsJSON = []byte(`{
  "data": [
    {
      "id": "ra.985484166",
      "type": "stations",
      "href": "/v1/catalog/us/stations/ra.985484166",
      "attributes": {
        "url": "https://itunes.apple.com/us/station/alternative/idra.985484166",
        "name": "Alternative",
        "artwork": {
          "width": 4320,
          "height": 1080,
          "url": "https://is3-ssl.mzstatic.com/image/thumb/Features111/v4/9a/e0/94/9ae0941a-721e-0b00-527c-e57723dad20f/source/{w}x{h}bb.jpg",
          "bgColor": "ede9d4",
          "textColor1": "140c04",
          "textColor2": "121210",
          "textColor3": "3f382d",
          "textColor4": "3d3d37"
        },
        "playParams": {
          "id": "ra.985484166",
          "kind": "radioStation"
        },
        "editorialNotes": {
          "name": "Alternative",
          "short": "The margins to mainstream."
        },
        "isLive": false
      }
    }
  ]
}`)

var stations = &Stations{
	Data: []Station{
		{
			Id:   "ra.985484166",
			Type: "stations",
			Href: "/v1/catalog/us/stations/ra.985484166",
			Attributes: StationAttributes{
				URL:  "https://itunes.apple.com/us/station/alternative/idra.985484166",
				Name: "Alternative",
				Artwork: Artwork{
					Width:      4320,
					Height:     1080,
					URL:        "https://is3-ssl.mzstatic.com/image/thumb/Features111/v4/9a/e0/94/9ae0941a-721e-0b00-527c-e57723dad20f/source/{w}x{h}bb.jpg",
					BgColor:    "ede9d4",
					TextColor1: "140c04",
					TextColor2: "121210",
					TextColor3: "3f382d",
					TextColor4: "3d3d37",
				},
				PlayParams: PlayParameters{
					Id:   "ra.985484166",
					Kind: "radioStation",
				},
				EditorialNotes: &EditorialNotes{
					Name:  "Alternative",
					Short: "The margins to mainstream.",
				},
				IsLive: false,
			},
		},
	},
}
