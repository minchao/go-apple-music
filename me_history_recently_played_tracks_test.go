package applemusic

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestMeService_GetHistoryRecentlyPlayedTracks(t *testing.T) {
	setup()
	defer teardown()

	historyRecentlyPlayedTracksJSON := []byte(`{
	  "next": "/v1/me/recent/played/tracks?offset=30",
	  "data": [
		{
		  "id": "7018167",
		  "type": "songs",
		  "href": "/v1/catalog/de/songs/7018167",
		  "attributes": {
			"albumName": "Soulhack",
			"genreNames": [
			  "Electronic",
			  "Musik"
			],
			"trackNumber": 1,
			"releaseDate": "2003-02-06",
			"durationInMillis": 148720,
			"isrc": "DEP960300040",
			"artwork": {
			  "width": 672,
			  "height": 600,
			  "url": "https://is1-ssl.mzstatic.com/image/thumb/Music/y2004/m04/d27/h00/s05.vqgpmhur.jpg/{w}x{h}bb.jpg",
			  "bgColor": "91c4ec",
			  "textColor1": "090a0b",
			  "textColor2": "070c0f",
			  "textColor3": "243038",
			  "textColor4": "23313b"
			},
			"playParams": {
			  "id": "7018167",
			  "kind": "song"
			},
			"url": "https://music.apple.com/de/album/city-ports/7018192?i=7018167",
			"discNumber": 1,
			"hasLyrics": false,
			"isAppleDigitalMaster": false,
			"name": "City Ports",
			"previews": [
			  {
				"url": "https://audio-ssl.itunes.apple.com/itunes-assets/AudioPreview115/v4/51/a3/67/51a36707-d65a-81f3-fe31-59804fb2bdfa/mzaf_13431577798219308851.plus.aac.p.m4a"
			  }
			],
			"artistName": "Forss"
		  }
		},
		{
		  "id": "712862708",
		  "type": "songs",
		  "href": "/v1/catalog/de/songs/712862708",
		  "attributes": {
			"albumName": "Walking On a Dream (Special Edition)",
			"genreNames": [
			  "Electronic",
			  "Musik",
			  "Rock",
			  "Adult Alternative",
			  "Alternative"
			],
			"trackNumber": 2,
			"releaseDate": "2008-08-30",
			"durationInMillis": 198440,
			"isrc": "AUEI10800039",
			"artwork": {
			  "width": 1400,
			  "height": 1400,
			  "url": "https://is1-ssl.mzstatic.com/image/thumb/Music115/v4/7f/63/2c/7f632c08-3960-de3d-d34c-ee8140038b69/13UADIM60773.rgb.jpg/{w}x{h}bb.jpg",
			  "bgColor": "0c123e",
			  "textColor1": "fdfdf5",
			  "textColor2": "becbed",
			  "textColor3": "cdced0",
			  "textColor4": "9aa6c9"
			},
			"composerName": "Luke Steele, Jonathan Sloan & Nick Littlemore",
			"url": "https://music.apple.com/de/album/walking-on-a-dream/712862605?i=712862708",
			"playParams": {
			  "id": "712862708",
			  "kind": "song"
			},
			"discNumber": 1,
			"hasLyrics": true,
			"isAppleDigitalMaster": false,
			"name": "Walking On a Dream",
			"previews": [
			  {
				"url": "https://audio-ssl.itunes.apple.com/itunes-assets/AudioPreview115/v4/b1/fd/91/b1fd910d-50d6-16bc-1dde-39a1175769f0/mzaf_10336910136317283104.plus.aac.p.m4a"
			  }
			],
			"artistName": "Empire of the Sun"
		  }
		}
	  ]
}`)

	mux.HandleFunc("/v1/me/recent/played/tracks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(historyRecentlyPlayedTracksJSON)
	})

	got, _, err := client.Me.GetHistoryRecentlyPlayedTracks(context.Background(), nil)
	if err != nil {
		t.Errorf("Me.GetAllLibrarySongs returned error: %v", err)
	}

	want := historyRecentlyPlayedTracks
	want.Next = "/v1/me/recent/played/tracks?offset=30"

	if !cmp.Equal(got, want) {
		t.Errorf("Me.TestMeService_GetHistoryRecentlyPlayedTracks = %+v, want %+v", got, want)
	}
}

var historyRecentlyPlayedTracks = &HistoryRecentlyPlayedTracks{
	Data: []HistoryRecentlyPlayedTrack{
		{
			Id:   "7018167",
			Type: "songs",
			Href: "/v1/catalog/de/songs/7018167",
			Attributes: HistoryRecentlyPlayedTrackAttributes{
				AlbumName: "Soulhack",
				GenreNames: []string{
					"Electronic",
					"Musik",
				},
				TrackNumber:      1,
				ReleaseDate:      "2003-02-06",
				DurationInMillis: 148720,
				ISRC:             "DEP960300040",
				Artwork: Artwork{
					Width:      672,
					Height:     600,
					URL:        "https://is1-ssl.mzstatic.com/image/thumb/Music/y2004/m04/d27/h00/s05.vqgpmhur.jpg/{w}x{h}bb.jpg",
					BgColor:    "91c4ec",
					TextColor1: "090a0b",
					TextColor2: "070c0f",
					TextColor3: "243038",
					TextColor4: "23313b",
					IsMosaic:   false,
				},
				PlayParams: PlayParameters{
					Id:   "7018167",
					Kind: "song",
				},
				URL:                  "https://music.apple.com/de/album/city-ports/7018192?i=7018167",
				DiscNumber:           1,
				HasLyrics:            false,
				IsAppleDigitalMaster: false,
				Name:                 "City Ports",
				Previews: []Preview{
					{
						Url: "https://audio-ssl.itunes.apple.com/itunes-assets/AudioPreview115/v4/51/a3/67/51a36707-d65a-81f3-fe31-59804fb2bdfa/mzaf_13431577798219308851.plus.aac.p.m4a",
					},
				},
				ArtistName: "Forss",
			},
		},
		{
			Id:   "712862708",
			Type: "songs",
			Href: "/v1/catalog/de/songs/712862708",
			Attributes: HistoryRecentlyPlayedTrackAttributes{
				AlbumName: "Walking On a Dream (Special Edition)",
				GenreNames: []string{
					"Electronic",
					"Musik",
					"Rock",
					"Adult Alternative",
					"Alternative",
				},
				TrackNumber:      2,
				ReleaseDate:      "2008-08-30",
				DurationInMillis: 198440,
				ISRC:             "AUEI10800039",
				Artwork: Artwork{
					Width:      1400,
					Height:     1400,
					URL:        "https://is1-ssl.mzstatic.com/image/thumb/Music115/v4/7f/63/2c/7f632c08-3960-de3d-d34c-ee8140038b69/13UADIM60773.rgb.jpg/{w}x{h}bb.jpg",
					BgColor:    "0c123e",
					TextColor1: "fdfdf5",
					TextColor2: "becbed",
					TextColor3: "cdced0",
					TextColor4: "9aa6c9",
					IsMosaic:   false,
				},
				PlayParams: PlayParameters{
					Id:   "712862708",
					Kind: "song",
				},
				DiscNumber:           1,
				URL:                  "https://music.apple.com/de/album/walking-on-a-dream/712862605?i=712862708",
				HasLyrics:            true,
				IsAppleDigitalMaster: false,
				Name:                 "Walking On a Dream",
				Previews: []Preview{
					{
						Url: "https://audio-ssl.itunes.apple.com/itunes-assets/AudioPreview115/v4/b1/fd/91/b1fd910d-50d6-16bc-1dde-39a1175769f0/mzaf_10336910136317283104.plus.aac.p.m4a",
					},
				},
				ArtistName: "Empire of the Sun",
			},
		},
	},
	Next: "/v1/me/recent/played/tracks?offset=30",
}
