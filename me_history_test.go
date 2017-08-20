package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestMeService_GetHistoryHeavyRotation(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/me/history/heavy-rotation", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(meHistoryHeavyRotationJSON)
	})

	got, _, err := client.Me.GetHistoryHeavyRotation(context.Background(), nil)
	if err != nil {
		t.Errorf("Me.GetHistoryHeavyRotation returned error: %v", err)
	}
	if want := meHistoryHeavyRotation; !reflect.DeepEqual(got, want) {
		t.Errorf("Me.GetHistoryHeavyRotation = %+v, want %+v", got, want)
	}
}

var meHistoryHeavyRotationJSON = []byte(`{
  "data": [
    {
      "attributes": {
        "artistName": "C\u00e9line Dion",
        "artwork": {
          "bgColor": "121413",
          "height": 600,
          "textColor1": "fefefe",
          "textColor2": "d5d5d5",
          "textColor3": "cfcfcf",
          "textColor4": "aeaeae",
          "url": "https://example.mzstatic.com/image/thumb/Music/v4/88/23/5b/88235b62-fc2f-e457-c55e-0fe689fb467b/source/{w}x{h}bb.jpg",
          "width": 604
        },
        "copyright": "\u2117 Compilation (P) 2008 Sony Music Entertainment Canada Inc.",
        "genreNames": [
          "Pop",
          "Music",
          "Adult Contemporary",
          "Soft Rock",
          "Vocal",
          "Rock"
        ],
        "isComplete": true,
        "isSingle": false,
        "name": "The Essential Celine Dion",
        "playParams": {
          "id": "464056948",
          "kind": "album"
        },
        "releaseDate": "2008",
        "trackCount": 27,
        "url": "https://itunes.apple.com/us/album/the-essential-celine-dion/id464056948"
      },
      "href": "/v1/catalog/us/albums/464056948",
      "id": "464056948",
      "type": "albums"
    }
  ],
  "href": "/v1/me/history/heavy-rotation?limit=1",
  "next": "/v1/me/history/heavy-rotation?offset=1"
}`)

var meHistoryHeavyRotation = &HistoryHeavyRotation{
	Data: []Resource{
		{
			[]byte(`{
      "attributes": {
        "artistName": "C\u00e9line Dion",
        "artwork": {
          "bgColor": "121413",
          "height": 600,
          "textColor1": "fefefe",
          "textColor2": "d5d5d5",
          "textColor3": "cfcfcf",
          "textColor4": "aeaeae",
          "url": "https://example.mzstatic.com/image/thumb/Music/v4/88/23/5b/88235b62-fc2f-e457-c55e-0fe689fb467b/source/{w}x{h}bb.jpg",
          "width": 604
        },
        "copyright": "\u2117 Compilation (P) 2008 Sony Music Entertainment Canada Inc.",
        "genreNames": [
          "Pop",
          "Music",
          "Adult Contemporary",
          "Soft Rock",
          "Vocal",
          "Rock"
        ],
        "isComplete": true,
        "isSingle": false,
        "name": "The Essential Celine Dion",
        "playParams": {
          "id": "464056948",
          "kind": "album"
        },
        "releaseDate": "2008",
        "trackCount": 27,
        "url": "https://itunes.apple.com/us/album/the-essential-celine-dion/id464056948"
      },
      "href": "/v1/catalog/us/albums/464056948",
      "id": "464056948",
      "type": "albums"
    }`),
		},
	},
	Href: "/v1/me/history/heavy-rotation?limit=1",
	Next: "/v1/me/history/heavy-rotation?offset=1",
}
