package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestMeService_GetAllLibrarySongs(t *testing.T) {
	setup()
	defer teardown()

	librarySongsJSON := []byte(`{
    "data": [
        {
            "attributes": {
                "albumName": "War (Deluxe Edition) [Remastered]",
                "artistName": "U2",
                "artwork": {
                    "height": 1200,
                    "url": "https://example.mzstatic.com/image/thumb/Features/2e/72/92/dj.wuyxwvik.jpg/{w}x{h}bb.jpg",
                    "width": 1200
                },
                "name": "\"40\"",
                "playParams": {
                    "id": "i.vMXdDeVhKQWRAd",
                    "isLibrary": true,
                    "kind": "song"
                },
                "trackNumber": 10
            },
            "href": "/v1/me/library/songs/i.vMXdDeVhKQWRAd",
            "id": "i.vMXdDeVhKQWRAd",
            "type": "library-songs"
        },
        {
            "attributes": {
                "albumName": "War",
                "artistName": "U2",
                "artwork": {
                    "height": 1200,
                    "url": "https://example.mzstatic.com/image/thumb/Music/4b/ca/43/mzi.bxlrvukd.jpg/{w}x{h}bb.jpg",
                    "width": 1200
                },
                "name": "\"40\"",
                "playParams": {
                    "id": "i.dlvVYxxTPLaemG",
                    "isLibrary": true,
                    "kind": "song"
                },
                "trackNumber": 10
            },
            "href": "/v1/me/library/songs/i.dlvVYxxTPLaemG",
            "id": "i.dlvVYxxTPLaemG",
            "type": "library-songs"
        }
    ],
    "href": "/v1/me/library/songs",
    "next": "/v1/me/library/songs?offset=100"
}`)

	mux.HandleFunc("/v1/me/library/songs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(librarySongsJSON)
	})

	got, _, err := client.Me.GetAllLibrarySongs(context.Background(), nil)
	if err != nil {
		t.Errorf("Me.GetAllLibrarySongs returned error: %v", err)
	}

	want := librarySongs
	want.Next = "/v1/me/library/songs?offset=100"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Me.TestMeService_GetAllLibrarySongs = %+v, want %+v", got, want)
	}
}

var librarySongs = &LibrarySongs{
	Data: []LibrarySong{
		{
			Attributes: LibrarySongAttributes{
				AlbumName:  "War (Deluxe Edition) [Remastered]",
				ArtistName: "U2",
				Artwork: Artwork{
					Height: 1200,
					URL:    "https://example.mzstatic.com/image/thumb/Features/2e/72/92/dj.wuyxwvik.jpg/{w}x{h}bb.jpg",
					Width:  1200,
				},
				Name: "\"40\"",
				PlayParams: PlayParameters{
					Id:        "i.vMXdDeVhKQWRAd",
					IsLibrary: true,
					Kind:      "song",
				},
				TrackNumber: 10,
			},
			Href: "/v1/me/library/songs/i.vMXdDeVhKQWRAd",
			Id:   "i.vMXdDeVhKQWRAd",
			Type: "library-songs",
		},
		{
			Attributes: LibrarySongAttributes{
				AlbumName:  "War",
				ArtistName: "U2",
				Artwork: Artwork{
					Height: 1200,
					URL:    "https://example.mzstatic.com/image/thumb/Music/4b/ca/43/mzi.bxlrvukd.jpg/{w}x{h}bb.jpg",
					Width:  1200,
				},
				Name: "\"40\"",
				PlayParams: PlayParameters{
					Id:        "i.dlvVYxxTPLaemG",
					IsLibrary: true,
					Kind:      "song",
				},
				TrackNumber: 10,
			},
			Href: "/v1/me/library/songs/i.dlvVYxxTPLaemG",
			Id:   "i.dlvVYxxTPLaemG",
			Type: "library-songs",
		},
	},
	Href: "/v1/me/library/songs",
	Next: "/v1/me/library/songs?offset=100",
}
