package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestCatalogService_Search(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/search", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"term":  "james+brown",
			"limit": "1",
			"types": "artists,albums",
		})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(searchJSON)
	})

	opt := &SearchOptions{
		Term:  "james+brown",
		Limit: 1,
		Types: "artists,albums",
	}

	got, _, err := client.Catalog.Search(context.Background(), "us", opt)
	if err != nil {
		t.Errorf("Catalog.Search returned error: %v", err)
	}
	if want := search; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.Search = %+v, want %+v", got, want)
	}
}

func TestCatalogService_SearchHints(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/search/hints", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"term":  "james+brown",
			"limit": "1",
			"types": "artists,albums",
		})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
    "results": {
        "terms": [
            "love kendrick lamar",
            "love on the brain",
            "love galore",
            "love",
            "love songs",
            "love the way you lie",
            "loverboy",
            "lovely day",
            "love me like you do",
            "love me now"
        ]
    }
}`))
	})

	opt := &SearchHintsOptions{
		Term:  "james+brown",
		Limit: 1,
		Types: "artists,albums",
	}

	searchHints := &SearchHints{
		Results: SearchHintsResults{
			Terms: []string{
				"love kendrick lamar",
				"love on the brain",
				"love galore",
				"love",
				"love songs",
				"love the way you lie",
				"loverboy",
				"lovely day",
				"love me like you do",
				"love me now",
			},
		},
	}

	got, _, err := client.Catalog.SearchHints(context.Background(), "us", opt)
	if err != nil {
		t.Errorf("Catalog.SearchHints returned error: %v", err)
	}
	if want := searchHints; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.SearchHints = %+v, want %+v", got, want)
	}
}

var searchJSON = []byte(`{
    "results": {
        "albums": {
            "href": "/v1/catalog/us/search?types=albums&term=james+brown&limit=1",
            "next": "/v1/catalog/us/search?types=albums&term=james+brown&offset=1",
            "data": [
                {
                    "id": "900721190",
                    "type": "albums",
                    "href": "/v1/catalog/us/albums/900721190",
                    "attributes": {
                        "artwork": {
                            "width": 1400,
                            "height": 1400,
                            "url": "https://is1-ssl.mzstatic.com/image/thumb/Music4/v4/76/85/e5/7685e5c8-9346-88db-95ff-af87bf84151b/source/{w}x{h}bb.jpg",
                            "bgColor": "ffffff",
                            "textColor1": "0a0a09",
                            "textColor2": "2a240f",
                            "textColor3": "3b3b3a",
                            "textColor4": "544f3f"
                        },
                        "artistName": "James Brown",
                        "isSingle": false,
                        "url": "https://itunes.apple.com/us/album/20-all-time-greatest-hits/id900721190",
                        "isComplete": true,
                        "genreNames": [
                            "Soul",
                            "Music",
                            "R&B/Soul",
                            "Funk"
                        ],
                        "trackCount": 20,
                        "releaseDate": "1991-01-01",
                        "name": "20 All-Time Greatest Hits!",
                        "copyright": "℗ 1991 Universal Records, a Division of UMG Recordings, Inc.",
                        "playParams": {
                            "id": "900721190",
                            "kind": "album"
                        }
                    }
                }
            ]
        },
        "artists": {
            "href": "/v1/catalog/us/search?types=artists&term=james+brown&limit=1",
            "next": "/v1/catalog/us/search?types=artists&term=james+brown&offset=1",
            "data": [
                {
                    "id": "117118",
                    "type": "artists",
                    "href": "/v1/catalog/us/artists/117118",
                    "attributes": {
                        "url": "https://itunes.apple.com/us/artist/james-brown/id117118",
                        "name": "James Brown",
                        "genreNames": [
                            "R&B/Soul"
                        ],
                        "editorialNotes": {
                            "standard": "The Godfather of Soul, The Hardest Working Man in Show Business, Soul Brother Number One . . . Mountain-of-a-man nicknames, no doubt, but not one of them ever did James Brown the justice he deserved. JB’s influence is so large it’s plain impossible to imagine what music today would sound like without him. James didn’t just master soul and funk; he invented them. From the on-bended-knee plea called “Try Me” straight to ground zero of the hip-hop revolution, where his screams, grunts, and funky drummer backboned every cut worth mentioning, Brown was soul power itself, an inspiration to an entire nation. Now that Mr. Dynamite has taken his last sweat-drenched shuffle offstage, we’re all left like an awe-struck Apollo audience — still screaming for just one more encore while trying to take in the force of nature we were blessed to witness. James Brown’s face should be on money. His likeness carved on a mountain. And, most importantly, his music played forever."
                        }
                    }
                }
            ]
        }
    }
}`)

var search = &Search{
	Results: SearchResults{
		Albums: &Albums{
			Href: "/v1/catalog/us/search?types=albums&term=james+brown&limit=1",
			Next: "/v1/catalog/us/search?types=albums&term=james+brown&offset=1",
			Data: []Album{
				{
					Id:   "900721190",
					Type: "albums",
					Href: "/v1/catalog/us/albums/900721190",
					Attributes: AlbumAttributes{
						Artwork: Artwork{
							Width:      1400,
							Height:     1400,
							URL:        "https://is1-ssl.mzstatic.com/image/thumb/Music4/v4/76/85/e5/7685e5c8-9346-88db-95ff-af87bf84151b/source/{w}x{h}bb.jpg",
							BgColor:    "ffffff",
							TextColor1: "0a0a09",
							TextColor2: "2a240f",
							TextColor3: "3b3b3a",
							TextColor4: "544f3f",
						},
						ArtistName: "James Brown",
						IsSingle:   false,
						URL:        "https://itunes.apple.com/us/album/20-all-time-greatest-hits/id900721190",
						IsComplete: true,
						GenreNames: []string{
							"Soul",
							"Music",
							"R&B/Soul",
							"Funk",
						},
						TrackCount:  20,
						ReleaseDate: "1991-01-01",
						Name:        "20 All-Time Greatest Hits!",
						Copyright:   "℗ 1991 Universal Records, a Division of UMG Recordings, Inc.",
						PlayParams: &PlayParameters{
							Id:   "900721190",
							Kind: "album",
						},
					},
				},
			},
		},
		Artists: &Artists{
			Href: "/v1/catalog/us/search?types=artists&term=james+brown&limit=1",
			Next: "/v1/catalog/us/search?types=artists&term=james+brown&offset=1",
			Data: []Artist{
				{
					Id:   "117118",
					Type: "artists",
					Href: "/v1/catalog/us/artists/117118",
					Attributes: ArtistAttributes{
						URL:  "https://itunes.apple.com/us/artist/james-brown/id117118",
						Name: "James Brown",
						GenreNames: []string{
							"R&B/Soul",
						},
						EditorialNotes: &EditorialNotes{
							Standard: "The Godfather of Soul, The Hardest Working Man in Show Business, Soul Brother Number One . . . Mountain-of-a-man nicknames, no doubt, but not one of them ever did James Brown the justice he deserved. JB’s influence is so large it’s plain impossible to imagine what music today would sound like without him. James didn’t just master soul and funk; he invented them. From the on-bended-knee plea called “Try Me” straight to ground zero of the hip-hop revolution, where his screams, grunts, and funky drummer backboned every cut worth mentioning, Brown was soul power itself, an inspiration to an entire nation. Now that Mr. Dynamite has taken his last sweat-drenched shuffle offstage, we’re all left like an awe-struck Apollo audience — still screaming for just one more encore while trying to take in the force of nature we were blessed to witness. James Brown’s face should be on money. His likeness carved on a mountain. And, most importantly, his music played forever.",
						},
					},
				},
			},
		},
	},
}
