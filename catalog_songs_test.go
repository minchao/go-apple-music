package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestCatalogService_GetSong(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/songs/900032829", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(songsJSON)
	})

	got, _, err := client.Catalog.GetSong(context.Background(), "us", "900032829", nil)
	if err != nil {
		t.Errorf("Catalog.GetSong returned error: %v", err)
	}
	if want := songs; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.GetSong = %+v, want %+v", got, want)
	}
}

func TestCatalogService_GetSongsByIds(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/songs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"ids": "203709340,201281527",
		})

		w.WriteHeader(http.StatusOK)
	})

	_, _, err := client.Catalog.GetSongsByIds(context.Background(), "us", []string{"203709340", "201281527"}, nil)
	if err != nil {
		t.Errorf("Catalog.GetSongsByIds returned error: %v", err)
	}
}

var songsJSON = []byte(`{
    "data": [
        {
            "id": "900032829",
            "type": "songs",
            "href": "/v1/catalog/us/songs/900032829",
            "attributes": {
                "artwork": {
                    "width": 2400,
                    "height": 2400,
                    "url": "https://is1-ssl.mzstatic.com/image/thumb/Music3/v4/8d/5f/4e/8d5f4e8f-d677-ba24-15f0-f8035047a4cc/source/{w}x{h}bb.jpg",
                    "bgColor": "346687",
                    "textColor1": "c9fcf0",
                    "textColor2": "b4fbf3",
                    "textColor3": "abdedb",
                    "textColor4": "9bdddd"
                },
                "artistName": "Michael de Jong",
                "url": "https://itunes.apple.com/us/album/something-for-the-pain/id900032785?i=900032829",
                "discNumber": 1,
                "genreNames": [
                    "Singer/Songwriter",
                    "Music",
                    "Blues"
                ],
                "durationInMillis": 327693,
                "releaseDate": "2014-07-11",
                "name": "Something For the Pain",
                "playParams": {
                    "id": "900032829",
                    "kind": "song"
                },
                "trackNumber": 7,
                "composerName": "Michael de Jong"
            },
            "relationships": {
                "albums": {
                    "data": [
                        {
                            "id": "900032785",
                            "type": "albums",
                            "href": "/v1/catalog/us/albums/900032785"
                        }
                    ],
                    "href": "/v1/catalog/us/songs/900032829/albums"
                },
                "artists": {
                    "data": [
                        {
                            "id": "6671250",
                            "type": "artists",
                            "href": "/v1/catalog/us/artists/6671250"
                        }
                    ],
                    "href": "/v1/catalog/us/songs/900032829/artists"
                }
            }
        }
    ]
}`)

var songs = &Songs{
	Data: []Song{
		{
			Id:   "900032829",
			Type: "songs",
			Href: "/v1/catalog/us/songs/900032829",
			Attributes: SongAttributes{
				Artwork: Artwork{
					Width:      2400,
					Height:     2400,
					URL:        "https://is1-ssl.mzstatic.com/image/thumb/Music3/v4/8d/5f/4e/8d5f4e8f-d677-ba24-15f0-f8035047a4cc/source/{w}x{h}bb.jpg",
					BgColor:    "346687",
					TextColor1: "c9fcf0",
					TextColor2: "b4fbf3",
					TextColor3: "abdedb",
					TextColor4: "9bdddd",
				},
				ArtistName: "Michael de Jong",
				URL:        "https://itunes.apple.com/us/album/something-for-the-pain/id900032785?i=900032829",
				DiscNumber: 1,
				GenreNames: []string{
					"Singer/Songwriter",
					"Music",
					"Blues",
				},
				DurationInMillis: 327693,
				ReleaseDate:      "2014-07-11",
				Name:             "Something For the Pain",
				PlayParams: &PlayParameters{
					Id:   "900032829",
					Kind: "song",
				},
				TrackNumber:  7,
				ComposerName: "Michael de Jong",
			},
			Relationships: SongRelationships{
				Albums: Albums{
					Data: []Album{
						{
							Id:   "900032785",
							Type: "albums",
							Href: "/v1/catalog/us/albums/900032785",
						},
					},
					Href: "/v1/catalog/us/songs/900032829/albums",
				},
				Artists: Artists{
					Data: []Artist{
						{
							Id:   "6671250",
							Type: "artists",
							Href: "/v1/catalog/us/artists/6671250",
						},
					},
					Href: "/v1/catalog/us/songs/900032829/artists",
				},
			},
		},
	},
}
