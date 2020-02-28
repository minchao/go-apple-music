package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestMeService_GetAllLibraryMusicVideos(t *testing.T) {
	setup()
	defer teardown()

	libraryMusicVideosJSON := []byte(`{
    "data": [
        {
            "attributes": {
                "albumName": "No Line On the Horizon (Deluxe Edition)",
                "artistName": "U2",
                "artwork": {
                    "height": 1200,
                    "url": "https://example.mzstatic.com/image/thumb/Video/dd/3b/87/mzi.nitgjlfh.jpeg/{w}x{h}bb.jpeg",
                    "width": 1200
                },
                "name": "Anton Corbijn's  Exclusive Film \"Linear\"",
                "trackNumber": 13
            },
            "href": "/v1/me/library/music-videos/i.xrX5kEvtv0VrNA",
            "id": "i.xrX5kEvtv0VrNA",
            "type": "library-music-videos"
        }
    ]
}`)

	mux.HandleFunc("/v1/me/library/music-videos", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(libraryMusicVideosJSON)
	})

	got, _, err := client.Me.GetAllLibraryMusicVideos(context.Background(), nil)
	if err != nil {
		t.Errorf("Me.GetAllLibraryMusicVideos returned error: %v", err)
	}

	want := libraryMusicVideos

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Me.TestMeService_GetAllLibraryMusicVideos = %+v, want %+v", got, want)
	}
}

var libraryMusicVideos = &LibraryMusicVideos{
	Data: []LibraryMusicVideo{
		{
			Attributes: LibraryMusicVideoAttributes{
				AlbumName:  "No Line On the Horizon (Deluxe Edition)",
				ArtistName: "U2",
				Artwork: Artwork{
					Height: 1200,
					URL:    "https://example.mzstatic.com/image/thumb/Video/dd/3b/87/mzi.nitgjlfh.jpeg/{w}x{h}bb.jpeg",
					Width:  1200,
				},
				Name:        "Anton Corbijn's  Exclusive Film \"Linear\"",
				TrackNumber: 13,
			},
			Href: "/v1/me/library/music-videos/i.xrX5kEvtv0VrNA",
			Id:   "i.xrX5kEvtv0VrNA",
			Type: "library-music-videos",
		},
	},
}
