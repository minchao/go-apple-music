package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestMeService_GetAllLibraryAlbums(t *testing.T) {
	setup()
	defer teardown()

	librarySongsJSON := []byte(`{
	  "next": "/v1/me/library/albums?offset=25",
	  "data": [
		{
		  "id": "l.DDD6dwm",
		  "type": "library-albums",
		  "href": "/v1/me/library/albums/l.DDD6dwm",
		  "relationships": {
			"artists": {
			  "href": "/v1/me/library/albums/l.tsKXU8/artists",
			  "data": [
				{
				  "id": "r.u7AyKin",
				  "type": "library-artists",
				  "href": "/v1/me/library/artists/r.u7AyKin"
				}
			  ]
			},
			"catalog": {
			  "href": "/v1/me/library/albums/l.tsKXU8/catalog",
			  "data": [
				{
				  "id": "568399080",
				  "type": "albums",
				  "href": "/v1/catalog/de/albums/568399080"
				}
			  ]
			}
		  },
		  "attributes": {
			"trackCount": 11,
			"genreNames": [
			  "Jazz"
			],
			"name": "A Curious Tale of Trials + Persons",
			"artistName": "Little Simz",
			"artwork": {
			  "width": 1200,
			  "height": 1200,
			  "url": "https://is1-ssl.mzstatic.com/image/thumb/Music7/v4/26/4f/e0/264fe07d-9ae8-7178-4265-4d07a887f162/5060186929491_1.jpg/{w}x{h}bb.jpg"
			},
			"playParams": {
			  "id": "l.DDD6dwm",
			  "kind": "album",
			  "isLibrary": true
			},
			"dateAdded": "2019-07-21T00:28:08Z"
		  }
		},
		{
		  "id": "l.U1ziHMp",
		  "type": "library-albums",
		  "href": "/v1/me/library/albums/l.U1ziHMp",
		  "attributes": {
			"trackCount": 14,
			"genreNames": [
			  "Funk"
			],
			"name": "A Funk Odyssey",
			"artistName": "Jamiroquai",
			"artwork": {
			  "width": 1200,
			  "height": 1200,
			  "url": "https://is1-ssl.mzstatic.com/image/thumb/Music115/v4/93/11/41/9311418f-b406-1dce-c2f4-d1175fe9ae45/696998595422.jpg/{w}x{h}bb.jpg"
			},
			"playParams": {
			  "id": "l.U1ziHMp",
			  "kind": "album",
			  "isLibrary": true
			},
			"dateAdded": "2019-07-21T00:28:08Z"
		  }
		}
	  ],
	  "meta": {
		"total": 716
	  }
	}`)

	mux.HandleFunc("/v1/me/library/albums", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(librarySongsJSON)
	})

	got, _, err := client.Me.GetAllLibraryAlbums(context.Background(), nil)
	if err != nil {
		t.Errorf("Me.GetAllLibraryAlbums returned error: %v", err)
	}

	want := libraryAlbums
	want.Next = "/v1/me/library/albums?offset=25"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Me.TestMeService_GetAllLibraryAlbums = %+v, want %+v", got, want)
	}
}

var libraryAlbums = &LibraryAlbums{
	Data: []LibraryAlbum{
		LibraryAlbum{
			Id:   "l.DDD6dwm",
			Type: "library-albums",
			Href: "/v1/me/library/albums/l.DDD6dwm",
			Attributes: LibraryAlbumAttributes{
				TrackCount: 11,
				GenreNames: []string{"Jazz"},
				Name:       "A Curious Tale of Trials + Persons",
				ArtistName: "Little Simz",
				Artwork: Artwork{
					Width:  1200,
					Height: 1200,
					URL:    "https://is1-ssl.mzstatic.com/image/thumb/Music7/v4/26/4f/e0/264fe07d-9ae8-7178-4265-4d07a887f162/5060186929491_1.jpg/{w}x{h}bb.jpg",
				},
				PlayParams: PlayParameters{
					Id:        "l.DDD6dwm",
					Kind:      "album",
					IsLibrary: true,
				},
				DateAdded: "2019-07-21T00:28:08Z",
			},
			Relationships: LibraryAlbumRelationships{
				Catalog: Albums{
					Data: []Album{
						Album{
							Id:   "568399080",
							Type: "albums",
							Href: "/v1/catalog/de/albums/568399080",
						},
					},
					Href: "/v1/me/library/albums/l.tsKXU8/catalog",
				},
				Artists: Artists{
					Href: "/v1/me/library/albums/l.tsKXU8/artists",
					Data: []Artist{
						Artist{
							Id:   "r.u7AyKin",
							Type: "library-artists",
							Href: "/v1/me/library/artists/r.u7AyKin",
						},
					},
				},
			},
		},
		LibraryAlbum{
			Id:   "l.U1ziHMp",
			Type: "library-albums",
			Href: "/v1/me/library/albums/l.U1ziHMp",
			Attributes: LibraryAlbumAttributes{
				TrackCount: 14,
				GenreNames: []string{"Funk"},
				Name:       "A Funk Odyssey",
				ArtistName: "Jamiroquai",
				Artwork: Artwork{
					Width:  1200,
					Height: 1200,
					URL:    "https://is1-ssl.mzstatic.com/image/thumb/Music115/v4/93/11/41/9311418f-b406-1dce-c2f4-d1175fe9ae45/696998595422.jpg/{w}x{h}bb.jpg",
				},
				PlayParams: PlayParameters{
					Id:        "l.U1ziHMp",
					Kind:      "album",
					IsLibrary: true,
				},
				DateAdded: "2019-07-21T00:28:08Z",
			},
		},
	},
	Next: "/v1/me/library/albums?offset=25",
}
