package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestCatalogService_GetMusicVideo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/music-videos/639032181", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(musicVideosJSON)
	})

	got, _, err := client.Catalog.GetMusicVideo(context.Background(), "us", "639032181", nil)
	if err != nil {
		t.Errorf("Catalog.GetMusicVideo returned error: %v", err)
	}
	if want := musicVideos; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.GetMusicVideo = %+v, want %+v", got, want)
	}
}

func TestCatalogService_GetMusicVideosByIds(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/music-videos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"ids": "639032181,870852283",
		})

		w.WriteHeader(http.StatusOK)
	})

	_, _, err := client.Catalog.GetMusicVideosByIds(context.Background(), "us", []string{"639032181", "870852283"}, nil)
	if err != nil {
		t.Errorf("Catalog.GetMusicVideosByIds returned error: %v", err)
	}
}

var musicVideosJSON = []byte(`{
  "data": [
    {
      "id": "639032181",
      "type": "music-videos",
      "href": "/v1/catalog/us/music-videos/639032181",
      "attributes": {
        "url": "https://itunes.apple.com/us/music-video/thatpower-feat-justin-bieber/id639032181",
        "name": "#thatPOWER (feat. Justin Bieber)",
        "genreNames": [
          "Pop"
        ],
        "isrc": "USUV71300701",
        "artistName": "will.i.am",
        "releaseDate": "2013-01-01",
        "artwork": {
          "width": 640,
          "height": 288,
          "url": "https://is1-ssl.mzstatic.com/image/thumb/Video/v4/ef/81/a8/ef81a86e-7144-ba9e-2de0-7b5fe6d2acfd/source/{w}x{h}bb.jpg"
        },
        "playParams": {
          "id": "639032181",
          "kind": "musicVideo"
        },
        "durationInMillis": 292833
      },
      "relationships": {
        "albums": {
          "data": [],
          "href": "/v1/catalog/us/music-videos/639032181/albums"
        },
        "artists": {
          "data": [
            {
              "id": "3495273",
              "type": "artists",
              "href": "/v1/catalog/us/artists/3495273"
            }
          ],
          "href": "/v1/catalog/us/music-videos/639032181/artists"
        }
      }
    }
  ]
}`)

var musicVideos = &MusicVideos{
	Data: []MusicVideo{
		{
			Id:   "639032181",
			Type: "music-videos",
			Href: "/v1/catalog/us/music-videos/639032181",
			Attributes: MusicVideoAttributes{
				URL:  "https://itunes.apple.com/us/music-video/thatpower-feat-justin-bieber/id639032181",
				Name: "#thatPOWER (feat. Justin Bieber)",
				GenreNames: []string{
					"Pop",
				},
				ISRC:        "USUV71300701",
				ArtistName:  "will.i.am",
				ReleaseDate: "2013-01-01",
				Artwork: Artwork{
					Width:  640,
					Height: 288,
					URL:    "https://is1-ssl.mzstatic.com/image/thumb/Video/v4/ef/81/a8/ef81a86e-7144-ba9e-2de0-7b5fe6d2acfd/source/{w}x{h}bb.jpg",
				},
				PlayParams: &PlayParameters{
					Id:   "639032181",
					Kind: "musicVideo",
				},
				DurationInMillis: 292833,
			},
			Relationships: MusicVideoRelationships{
				Albums: Albums{
					Data: []Album{},
					Href: "/v1/catalog/us/music-videos/639032181/albums",
				},
				Artists: Artists{
					Data: []Artist{
						{
							Id:   "3495273",
							Type: "artists",
							Href: "/v1/catalog/us/artists/3495273",
						},
					},
					Href: "/v1/catalog/us/music-videos/639032181/artists",
				},
			},
		},
	},
}
