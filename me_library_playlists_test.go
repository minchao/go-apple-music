package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestMeService_GetLibraryPlaylist(t *testing.T) {
	setup()
	defer teardown()

	playlistJSON := []byte(`{
    "data": [
        {
            "attributes": {
                "canDelete": true,
                "canEdit": true,
                "description": {
                    "standard": "My description"
                },
                "name": "Some Playlist",
                "playParams": {
                    "id": "p.MoGJYM3CYXW09B",
                    "isLibrary": true,
                    "kind": "playlist"
                }
            },
            "href": "/v1/me/library/playlists/p.MoGJYM3CYXW09B",
            "id": "p.MoGJYM3CYXW09B",
            "type": "library-playlists"
        }
    ]
}`)

	mux.HandleFunc("/v1/me/library/playlists/p.MoGJYM3CYXW09B", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(playlistJSON)
	})

	got, _, err := client.Me.GetLibraryPlaylist(context.Background(), "p.MoGJYM3CYXW09B", nil)
	if err != nil {
		t.Errorf("Me.GetLibraryPlaylist returned error: %v", err)
	}

	want := &LibraryPlaylists{
		Data: []LibraryPlaylist{
			{
				Attributes: LibraryPlaylistAttributes{
					CanDelete: true,
					CanEdit:   true,
					Description: &EditorialNotes{
						Standard: "My description",
					},
					Name: "Some Playlist",
					PlayParams: &PlayParameters{
						Id:        "p.MoGJYM3CYXW09B",
						Kind:      "playlist",
						IsLibrary: true,
					},
				},
				Href: "/v1/me/library/playlists/p.MoGJYM3CYXW09B",
				Id:   "p.MoGJYM3CYXW09B",
				Type: "library-playlists",
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Me.GetLibraryPlaylist = %+v, want %+v", got, want)
	}
}

func TestMeService_GetLibraryPlaylistByIds(t *testing.T) {
	setup()
	defer teardown()

	libraryPlaylistsJSON := []byte(`{
    "data": [
        {
            "attributes": {
                "canDelete": true,
                "canEdit": true,
                "description": {
                    "standard": "My description"
                },
                "name": "Some Playlist",
                "playParams": {
                    "id": "p.MoGJYM3CYXW09B",
                    "isLibrary": true,
                    "kind": "playlist"
                }
            },
            "href": "/v1/me/library/playlists/p.MoGJYM3CYXW09B",
            "id": "p.MoGJYM3CYXW09B",
            "type": "library-playlists"
        },
        {
            "attributes": {
                "canDelete": true,
                "canEdit": true,
                "name": "Media API Playlist",
                "playParams": {
                    "id": "p.8Wx6vK6IQeP0N2",
                    "isLibrary": true,
                    "kind": "playlist"
                }
            },
            "href": "/v1/me/library/playlists/p.8Wx6vK6IQeP0N2",
            "id": "p.8Wx6vK6IQeP0N2",
            "type": "library-playlists"
        }
    ]
}`)

	mux.HandleFunc("/v1/me/library/playlists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"ids": "p.MoGJYM3CYXW09B,p.8Wx6vK6IQeP0N2",
			"l":   "en-gb",
		})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(libraryPlaylistsJSON)
	})

	got, _, err := client.Me.GetLibraryPlaylistsByIds(
		context.Background(),
		&LibraryPlaylistsByIdsOptions{
			Ids:     "p.MoGJYM3CYXW09B,p.8Wx6vK6IQeP0N2",
			Options: Options{Language: "en-gb"},
		},
	)
	if err != nil {
		t.Errorf("Me.GetLibraryPlaylistsByIds returned error: %v", err)
	}

	want := libraryPlaylists

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Me.GetLibraryPlaylistsByIds = %+v, want %+v", got, want)
	}
}

func TestMeService_GetAllLibraryPlaylist(t *testing.T) {
	setup()
	defer teardown()

	libraryPlaylistsJSON := []byte(`{
    "data": [
        {
            "attributes": {
                "canDelete": true,
                "canEdit": true,
                "description": {
                    "standard": "My description"
                },
                "name": "Some Playlist",
                "playParams": {
                    "id": "p.MoGJYM3CYXW09B",
                    "isLibrary": true,
                    "kind": "playlist"
                }
            },
            "href": "/v1/me/library/playlists/p.MoGJYM3CYXW09B",
            "id": "p.MoGJYM3CYXW09B",
            "type": "library-playlists"
        },
        {
            "attributes": {
                "canDelete": true,
                "canEdit": true,
                "name": "Media API Playlist",
                "playParams": {
                    "id": "p.8Wx6vK6IQeP0N2",
                    "isLibrary": true,
                    "kind": "playlist"
                }
            },
            "href": "/v1/me/library/playlists/p.8Wx6vK6IQeP0N2",
            "id": "p.8Wx6vK6IQeP0N2",
            "type": "library-playlists"
        }
    ],
	"next": "/v1/me/library/playlists?offset=7&limit=2&l=en-gb"
}`)

	mux.HandleFunc("/v1/me/library/playlists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"offset": "5",
			"limit":  "2",
			"l":      "en-gb",
		})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(libraryPlaylistsJSON)
	})

	got, _, err := client.Me.GetAllLibraryPlaylists(
		context.Background(),
		&PageOptions{Limit: 2, Offset: 5, Options: Options{Language: "en-gb"}},
	)
	if err != nil {
		t.Errorf("Me.TestMeService_GetAllLibraryPlaylist returned error: %v", err)
	}

	want := libraryPlaylists
	want.Next = "/v1/me/library/playlists?offset=7&limit=2&l=en-gb"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Me.TestMeService_GetAllLibraryPlaylist = %+v, want %+v", got, want)
	}
}

var libraryPlaylists = &LibraryPlaylists{
	Data: []LibraryPlaylist{
		{
			Attributes: LibraryPlaylistAttributes{
				CanDelete: true,
				CanEdit:   true,
				Description: &EditorialNotes{
					Standard: "My description",
				},
				Name: "Some Playlist",
				PlayParams: &PlayParameters{
					Id:        "p.MoGJYM3CYXW09B",
					Kind:      "playlist",
					IsLibrary: true,
				},
			},
			Href: "/v1/me/library/playlists/p.MoGJYM3CYXW09B",
			Id:   "p.MoGJYM3CYXW09B",
			Type: "library-playlists",
		},
		{
			Attributes: LibraryPlaylistAttributes{
				CanDelete: true,
				CanEdit:   true,
				Name:      "Media API Playlist",
				PlayParams: &PlayParameters{
					Id:        "p.8Wx6vK6IQeP0N2",
					Kind:      "playlist",
					IsLibrary: true,
				},
			},
			Href: "/v1/me/library/playlists/p.8Wx6vK6IQeP0N2",
			Id:   "p.8Wx6vK6IQeP0N2",
			Type: "library-playlists",
		},
	},
}

func TestMeService_CreateLibraryPlaylist(t *testing.T) {
	setup()
	defer teardown()

	libraryPlaylistsJSON := []byte(`{
    "data":[
        {
            "id":"p.zp6KqKxsoQWAGN",
            "type":"library-playlists",
            "href":"/v1/me/library/playlists/p.zp6KqKxsoQWAGN",
            "attributes":{
                "name":"Some Playlist",
                "description":{
                    "standard":"My description"
                },
                "playParams":{
                    "id":"p.zp6KqKxsoQWAGN",
                    "kind":"playlist",
                    "isLibrary":true
                },
                "canEdit":true
            }
        }
    ]
}`)

	mux.HandleFunc("/v1/me/library/playlists", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testJsonBodyValues(t, r, []byte(`{"attributes":{"name":"Some Playlist","description":"My description"}}`))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(libraryPlaylistsJSON)
	})

	got, _, err := client.Me.CreateLibraryPlaylist(
		context.Background(),
		CreateLibraryPlaylist{CreateLibraryPlaylistAttributes{Name: "Some Playlist", Description: "My description"}, nil},
		nil,
	)
	if err != nil {
		t.Errorf("Me.TestMeService_CreateLibraryPlaylist returned error: %v", err)
	}

	want := newLibraryPlaylists

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Me.TestMeService_CreateLibraryPlaylist = %+v, want %+v", got, want)
	}
}

var newLibraryPlaylists = &LibraryPlaylists{
	Data: []LibraryPlaylist{
		{
			Attributes: LibraryPlaylistAttributes{
				CanDelete: false,
				CanEdit:   true,
				Description: &EditorialNotes{
					Standard: "My description",
				},
				Name: "Some Playlist",
				PlayParams: &PlayParameters{
					Id:        "p.zp6KqKxsoQWAGN",
					Kind:      "playlist",
					IsLibrary: true,
				},
			},
			Href: "/v1/me/library/playlists/p.zp6KqKxsoQWAGN",
			Id:   "p.zp6KqKxsoQWAGN",
			Type: "library-playlists",
		},
	},
}

func TestMeService_AddLibraryTracksToPlaylist(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/me/library/playlists/p.zp6KqKxsoQWAGN/tracks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testJsonBodyValues(t, r, []byte(`{"data":[{"id":"201281527","type":"music"}]}`))

		w.WriteHeader(http.StatusNoContent)
	})

	_, err := client.Me.AddLibraryTracksToPlaylist(
		context.Background(),
		"p.zp6KqKxsoQWAGN",
		CreateLibraryPlaylistTrackData{Data: []CreateLibraryPlaylistTrack{{Id: "201281527", Type: "music"}}},
	)
	if err != nil {
		t.Errorf("Me.TestMeService_CreateLibraryPlaylist returned error: %v", err)
	}
}
