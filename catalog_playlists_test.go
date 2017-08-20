package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestCatalogService_GetPlaylist(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/playlists/pl.acc464c750b94302b8806e5fcbe56e17", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(playlistsJSON)
	})

	got, _, err := client.Catalog.GetPlaylist(context.Background(), "us", "pl.acc464c750b94302b8806e5fcbe56e17", nil)
	if err != nil {
		t.Errorf("Catalog.GetPlaylist returned error: %v", err)
	}
	if want := playlists; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.GetPlaylist = %+v, want %+v", got, want)
	}
}

func TestCatalogService_GetPlaylistsByIds(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/playlists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"ids": "pl.acc464c750b94302b8806e5fcbe56e17,pl.97c6f95b0b884bedbcce117f9ea5d54b",
		})

		w.WriteHeader(http.StatusOK)
	})

	_, _, err := client.Catalog.GetPlaylistsByIds(context.Background(), "us", []string{"pl.acc464c750b94302b8806e5fcbe56e17", "pl.97c6f95b0b884bedbcce117f9ea5d54b"}, nil)
	if err != nil {
		t.Errorf("Catalog.GetPlaylistsByIds returned error: %v", err)
	}
}

var playlistsJSON = []byte(`{
  "data": [
    {
      "id": "pl.acc464c750b94302b8806e5fcbe56e17",
      "type": "playlists",
      "href": "/v1/catalog/us/playlists/pl.acc464c750b94302b8806e5fcbe56e17",
      "attributes": {
        "url": "https://itunes.apple.com/us/playlist/janet-jackson-no-1-songs/idpl.acc464c750b94302b8806e5fcbe56e17",
        "name": "Janet Jackson: No.1 Songs",
        "playlistType": "editorial",
        "lastModifiedDate": "2015-04-11T16:15:51Z",
        "curatorName": "Apple Music R&B",
        "artwork": {
          "width": 4320,
          "height": 1080,
          "url": "https://is4-ssl.mzstatic.com/image/thumb/Features71/v4/49/f0/f6/49f0f636-cefe-0fba-a6a1-01321374e768/source/{w}x{h}cc.jpg",
          "bgColor": "161d16",
          "textColor1": "ffffff",
          "textColor2": "e3aa71",
          "textColor3": "d0d1d0",
          "textColor4": "ba8e5f",
          "isMosaic": true
        },
        "playParams": {
          "id": "pl.acc464c750b94302b8806e5fcbe56e17",
          "kind": "playlist"
        }
      },
      "relationships": {
        "tracks": {
          "data": [
            {
              "id": "1043322098",
              "type": "songs",
              "href": "/v1/catalog/us/songs/1043322098",
              "attributes": {
                "artwork": {
                  "width": 1500,
                  "height": 1500,
                  "url": "https://is3-ssl.mzstatic.com/image/thumb/Music6/v4/30/bc/4b/30bc4bfa-dfb8-73eb-95ad-446077c716fb/source/{w}x{h}bb.jpg",
                  "bgColor": "9c1c25",
                  "textColor1": "f5e8ea",
                  "textColor2": "e8dbdd",
                  "textColor3": "e3bfc3",
                  "textColor4": "d9b5b8"
                },
                "artistName": "Janet Jackson",
                "url": "https://itunes.apple.com/us/album/when-i-think-of-you/id1043321686?i=1043322098",
                "discNumber": 1,
                "genreNames": [
                  "Pop",
                  "Music"
                ],
                "durationInMillis": 237485,
                "releaseDate": "1986-02-04",
                "name": "When I Think of You",
                "playParams": {
                  "id": "1043322098",
                  "kind": "song"
                },
                "trackNumber": 6,
                "composerName": "Janet Jackson, James Harris III & Terry Lewis"
              }
            }
          ],
          "href": "/v1/catalog/us/playlists/pl.acc464c750b94302b8806e5fcbe56e17/tracks"
        },
        "curator": {
          "data": [
            {
              "id": "976439551",
              "type": "apple-curators",
              "href": "/v1/catalog/us/apple-curators/976439551"
            }
          ],
          "href": "/v1/catalog/us/playlists/pl.acc464c750b94302b8806e5fcbe56e17/curator"
        }
      }
    }
  ]
}`)

var playlists = &Playlists{
	Data: []Playlist{
		{
			Id:   "pl.acc464c750b94302b8806e5fcbe56e17",
			Type: "playlists",
			Href: "/v1/catalog/us/playlists/pl.acc464c750b94302b8806e5fcbe56e17",
			Attributes: PlaylistAttributes{
				URL:              "https://itunes.apple.com/us/playlist/janet-jackson-no-1-songs/idpl.acc464c750b94302b8806e5fcbe56e17",
				Name:             "Janet Jackson: No.1 Songs",
				PlaylistType:     PlaylistTypeEditorial,
				LastModifiedDate: "2015-04-11T16:15:51Z",
				CuratorName:      "Apple Music R&B",
				Artwork: &Artwork{
					Width:      4320,
					Height:     1080,
					URL:        "https://is4-ssl.mzstatic.com/image/thumb/Features71/v4/49/f0/f6/49f0f636-cefe-0fba-a6a1-01321374e768/source/{w}x{h}cc.jpg",
					BgColor:    "161d16",
					TextColor1: "ffffff",
					TextColor2: "e3aa71",
					TextColor3: "d0d1d0",
					TextColor4: "ba8e5f",
					IsMosaic:   true,
				},
				PlayParams: &PlayParameters{
					Id:   "pl.acc464c750b94302b8806e5fcbe56e17",
					Kind: "playlist",
				},
			},
			Relationships: PlaylistRelationships{
				Tracks: Tracks{
					Data: []Track{
						{
							[]byte(`{
              "id": "1043322098",
              "type": "songs",
              "href": "/v1/catalog/us/songs/1043322098",
              "attributes": {
                "artwork": {
                  "width": 1500,
                  "height": 1500,
                  "url": "https://is3-ssl.mzstatic.com/image/thumb/Music6/v4/30/bc/4b/30bc4bfa-dfb8-73eb-95ad-446077c716fb/source/{w}x{h}bb.jpg",
                  "bgColor": "9c1c25",
                  "textColor1": "f5e8ea",
                  "textColor2": "e8dbdd",
                  "textColor3": "e3bfc3",
                  "textColor4": "d9b5b8"
                },
                "artistName": "Janet Jackson",
                "url": "https://itunes.apple.com/us/album/when-i-think-of-you/id1043321686?i=1043322098",
                "discNumber": 1,
                "genreNames": [
                  "Pop",
                  "Music"
                ],
                "durationInMillis": 237485,
                "releaseDate": "1986-02-04",
                "name": "When I Think of You",
                "playParams": {
                  "id": "1043322098",
                  "kind": "song"
                },
                "trackNumber": 6,
                "composerName": "Janet Jackson, James Harris III & Terry Lewis"
              }
            }`)},
					},
					Href: "/v1/catalog/us/playlists/pl.acc464c750b94302b8806e5fcbe56e17/tracks",
				},
				Curator: Curators{
					Data: []Curator{
						{
							Id:   "976439551",
							Type: "apple-curators",
							Href: "/v1/catalog/us/apple-curators/976439551",
						},
					},
					Href: "/v1/catalog/us/playlists/pl.acc464c750b94302b8806e5fcbe56e17/curator",
				},
			},
		},
	},
}
