package applemusic

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

func TestCatalogService_GetAllCharts(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/charts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"types": "songs,albums",
			"genre": "20",
			"limit": "1",
		})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(chartsJSON)
	})

	opt := &ChartsOptions{
		Types: "songs,albums",
		Genre: "20",
		Limit: 1,
	}

	got, _, err := client.Catalog.GetAllCharts(context.Background(), "us", opt)
	if err != nil {
		t.Errorf("Catalog.GetAllCharts returned error: %v", err)
	}
	if want := charts; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.GetAllCharts = %+v, want %+v", got, want)

		b, _ := json.Marshal(want)

		t.Log("json:", string(b))
	}
}

var chartsJSON = []byte(`{
    "results": {
        "albums": [
            {
                "name": "Most Played Albums on Apple Music",
                "chart": "most-played",
                "href": "/v1/catalog/us/charts?types=albums&chart=most-played&genre=20&limit=1",
                "next": "/v1/catalog/us/charts?types=albums&chart=most-played&genre=20&offset=1",
                "data": [
                    {
                        "id": "1256684768",
                        "type": "albums",
                        "href": "/v1/catalog/us/albums/1256684768",
                        "attributes": {
                            "artwork": {
                                "width": 1800,
                                "height": 1800,
                                "url": "https://is5-ssl.mzstatic.com/image/thumb/Music128/v4/fd/4b/c6/fd4bc65d-e552-a1b0-1a62-49d6fd4e77ae/source/{w}x{h}bb.jpg",
                                "bgColor": "eae7de",
                                "textColor1": "050c13",
                                "textColor2": "1d3336",
                                "textColor3": "33383b",
                                "textColor4": "465758"
                            },
                            "artistName": "Lana Del Rey",
                            "isSingle": false,
                            "url": "https://itunes.apple.com/us/album/lust-for-life/id1256684768",
                            "isComplete": true,
                            "genreNames": [
                                "Alternative",
                                "Music",
                                "Rock"
                            ],
                            "trackCount": 16,
                            "releaseDate": "2017-07-21",
                            "name": "Lust for Life",
                            "copyright": "℗ 2017 Lana Del Rey under exclusive licence to Polydor Ltd. (UK). Under exclusive licence to Interscope Records in the USA",
                            "playParams": {
                                "id": "1256684768",
                                "kind": "album"
                            },
                            "editorialNotes": {
                                "standard": "For the most part, Lana Del Rey’s fifth album is quintessentially her: gloomy, glamorous, and smitten with California. But a newfound lightness might surprise longtime fans. Each song on <i>Lust</i> feels like a postcard from a dream: She fantasizes about 1969 (“Coachella - Woodstock In My Mind”), outruns paparazzi on the Pacific Coast Highway (“13 Beaches”), and dances on the H of the Hollywood sign (“Lust for Life” feat. The Weeknd). She even duets with Stevie Nicks, the queen of bittersweet rock. On “Get Free,” she makes a vow to shift her mindset: \"Now I do, I want to move/Out of the black, into the blue.”",
                                "short": "Back in all her gloomy glamour, with Stevie Nicks and A$AP Rocky."
                            },
                            "contentRating": "explicit"
                        }
                    }
                ]
            }
        ],
        "songs": [
            {
                "name": "Most Played Songs on Apple Music",
                "chart": "most-played",
                "href": "/v1/catalog/us/charts?types=songs&chart=most-played&genre=20&limit=1",
                "next": "/v1/catalog/us/charts?types=songs&chart=most-played&genre=20&offset=1",
                "data": [
                    {
                        "id": "1233502267",
                        "type": "songs",
                        "href": "/v1/catalog/us/songs/1233502267",
                        "attributes": {
                            "artwork": {
                                "width": 1800,
                                "height": 1800,
                                "url": "https://is1-ssl.mzstatic.com/image/thumb/Music117/v4/2c/6b/cc/2c6bcc08-8a7a-dd89-d344-4e649f2a1bf8/source/{w}x{h}bb.jpg",
                                "bgColor": "0d1718",
                                "textColor1": "f6f9fe",
                                "textColor2": "f6b32b",
                                "textColor3": "c7cbd0",
                                "textColor4": "c89327"
                            },
                            "artistName": "Imagine Dragons",
                            "url": "https://itunes.apple.com/us/album/believer/id1233502263?i=1233502267",
                            "discNumber": 1,
                            "genreNames": [
                                "Alternative",
                                "Music"
                            ],
                            "durationInMillis": 204345,
                            "releaseDate": "2017-02-01",
                            "name": "Believer",
                            "playParams": {
                                "id": "1233502267",
                                "kind": "song"
                            },
                            "trackNumber": 3,
                            "composerName": "Dan Reynolds, Wayne Sermon, Ben McKee, Daniel Platzman, Robin Fredriksson, Mattias Larsson & Justin Drew Tranter"
                        }
                    }
                ]
            }
        ]
    }
}`)

var charts = &Charts{
	Results: ChartResults{
		Albums: &[]ChartAlbums{
			{
				Name:  "Most Played Albums on Apple Music",
				Chart: "most-played",
				Albums: Albums{
					Href: "/v1/catalog/us/charts?types=albums&chart=most-played&genre=20&limit=1",
					Next: "/v1/catalog/us/charts?types=albums&chart=most-played&genre=20&offset=1",
					Data: []Album{
						{
							Id:   "1256684768",
							Type: "albums",
							Href: "/v1/catalog/us/albums/1256684768",
							Attributes: AlbumAttributes{
								Artwork: Artwork{
									Width:      1800,
									Height:     1800,
									URL:        "https://is5-ssl.mzstatic.com/image/thumb/Music128/v4/fd/4b/c6/fd4bc65d-e552-a1b0-1a62-49d6fd4e77ae/source/{w}x{h}bb.jpg",
									BgColor:    "eae7de",
									TextColor1: "050c13",
									TextColor2: "1d3336",
									TextColor3: "33383b",
									TextColor4: "465758",
								},
								ArtistName: "Lana Del Rey",
								IsSingle:   false,
								URL:        "https://itunes.apple.com/us/album/lust-for-life/id1256684768",
								IsComplete: true,
								GenreNames: []string{
									"Alternative",
									"Music",
									"Rock",
								},
								TrackCount:  16,
								ReleaseDate: "2017-07-21",
								Name:        "Lust for Life",
								Copyright:   "℗ 2017 Lana Del Rey under exclusive licence to Polydor Ltd. (UK). Under exclusive licence to Interscope Records in the USA",
								PlayParams: &PlayParameters{
									Id:   "1256684768",
									Kind: "album",
								},
								EditorialNotes: &EditorialNotes{
									Standard: "For the most part, Lana Del Rey’s fifth album is quintessentially her: gloomy, glamorous, and smitten with California. But a newfound lightness might surprise longtime fans. Each song on <i>Lust</i> feels like a postcard from a dream: She fantasizes about 1969 (“Coachella - Woodstock In My Mind”), outruns paparazzi on the Pacific Coast Highway (“13 Beaches”), and dances on the H of the Hollywood sign (“Lust for Life” feat. The Weeknd). She even duets with Stevie Nicks, the queen of bittersweet rock. On “Get Free,” she makes a vow to shift her mindset: \"Now I do, I want to move/Out of the black, into the blue.”",
									Short:    "Back in all her gloomy glamour, with Stevie Nicks and A$AP Rocky.",
								},
								ContentRating: "explicit",
							},
						},
					},
				},
			},
		},
		Songs: &[]ChartSongs{
			{
				Name:  "Most Played Songs on Apple Music",
				Chart: "most-played",
				Songs: Songs{
					Href: "/v1/catalog/us/charts?types=songs&chart=most-played&genre=20&limit=1",
					Next: "/v1/catalog/us/charts?types=songs&chart=most-played&genre=20&offset=1",
					Data: []Song{
						{
							Id:   "1233502267",
							Type: "songs",
							Href: "/v1/catalog/us/songs/1233502267",
							Attributes: SongAttributes{
								Artwork: Artwork{
									Width:      1800,
									Height:     1800,
									URL:        "https://is1-ssl.mzstatic.com/image/thumb/Music117/v4/2c/6b/cc/2c6bcc08-8a7a-dd89-d344-4e649f2a1bf8/source/{w}x{h}bb.jpg",
									BgColor:    "0d1718",
									TextColor1: "f6f9fe",
									TextColor2: "f6b32b",
									TextColor3: "c7cbd0",
									TextColor4: "c89327",
								},
								ArtistName: "Imagine Dragons",
								URL:        "https://itunes.apple.com/us/album/believer/id1233502263?i=1233502267",
								DiscNumber: 1,
								GenreNames: []string{
									"Alternative",
									"Music",
								},
								DurationInMillis: 204345,
								ReleaseDate:      "2017-02-01",
								Name:             "Believer",
								PlayParams: &PlayParameters{
									Id:   "1233502267",
									Kind: "song",
								},
								TrackNumber:  3,
								ComposerName: "Dan Reynolds, Wayne Sermon, Ben McKee, Daniel Platzman, Robin Fredriksson, Mattias Larsson & Justin Drew Tranter",
							},
						},
					},
				},
			},
		},
	},
}
