package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestCatalogService_GetGenre(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/genres/14", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{
    "data": [
        {
            "id": "14",
            "type": "genres",
            "href": "/v1/catalog/us/genres/14",
            "attributes": {
                "name": "Pop"
            }
        }
    ]
}`))
	})

	genre := &Genres{
		Data: []Genre{
			{
				Id:   "14",
				Type: "genres",
				Href: "/v1/catalog/us/genres/14",
				Attributes: GenreAttributes{
					Name: "Pop",
				},
			},
		},
	}

	got, _, err := client.Catalog.GetGenre(context.Background(), "us", "14", nil)
	if err != nil {
		t.Errorf("Catalog.GetGenre returned error: %v", err)
	}
	if want := genre; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.GetGenre = %+v, want %+v", got, want)
	}
}

func TestCatalogService_GetGenresByIds(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/genres", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"ids": "14,21",
		})

		w.WriteHeader(http.StatusOK)
	})

	_, _, err := client.Catalog.GetGenresByIds(context.Background(), "us", []string{"14", "21"}, nil)
	if err != nil {
		t.Errorf("Catalog.GetGenresByIds returned error: %v", err)
	}
}

func TestCatalogService_GetAllGenres(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/genres", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"offset": "5",
			"limit":  "2",
			"l":      "en-us",
		})

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(genresJSON)
	})

	opt := &PageOptions{
		Offset: 5,
		Limit:  2,
		Options: Options{
			Language: "en-us",
		},
	}

	got, _, err := client.Catalog.GetAllGenres(context.Background(), "us", opt)
	if err != nil {
		t.Errorf("Catalog.GetAllGenres returned error: %v", err)
	}
	if want := genres; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.GetAllGenres = %+v, want %+v", got, want)
	}
}

var genresJSON = []byte(`{
    "data": [
        {
            "id": "5",
            "type": "genres",
            "href": "/v1/catalog/us/genres/5",
            "attributes": {
                "name": "Classical"
            }
        },
        {
            "id": "3",
            "type": "genres",
            "href": "/v1/catalog/us/genres/3",
            "attributes": {
                "name": "Comedy"
            }
        }
    ],
    "next": "/v1/catalog/us/genres?offset=7"
}`)

var genres = &Genres{
	Data: []Genre{
		{
			Id:   "5",
			Type: "genres",
			Href: "/v1/catalog/us/genres/5",
			Attributes: GenreAttributes{
				Name: "Classical",
			},
		},
		{
			Id:   "3",
			Type: "genres",
			Href: "/v1/catalog/us/genres/3",
			Attributes: GenreAttributes{
				Name: "Comedy",
			},
		},
	},
	Next: "/v1/catalog/us/genres?offset=7",
}
