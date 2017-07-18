package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestStorefrontsService_Get(t *testing.T) {
	setup()
	defer teardown()

	storefrontJSON := []byte(`{
  "data": [
    {
      "attributes": {
        "defaultLanguageTag": "ja-jp",
        "name": "Japan",
        "supportedLanguageTags": [
          "ja-jp",
          "en-us"
        ]
      },
      "href": "/v1/storefronts/jp",
      "id": "jp",
      "type": "storefronts"
    }
  ]
}`)

	mux.HandleFunc("/v1/storefronts/jp", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(storefrontJSON)
	})

	got, _, err := client.Storefront.Get(context.Background(), "jp", nil)
	if err != nil {
		t.Errorf("Storefront.Get returned error: %v", err)
	}

	want := &Storefronts{
		Data: []Storefront{
			{
				Attributes: Attributes{
					DefaultLanguageTag: "ja-jp",
					Name:               "Japan",
					SupportedLanguageTags: []string{
						"ja-jp",
						"en-us",
					},
				},
				Href: "/v1/storefronts/jp",
				Id:   "jp",
				Type: "storefronts",
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Storefront.Get = %+v, want %+v", got, want)
	}
}

func TestStorefrontsService_GetByIds(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/storefronts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"ids": "jp,tw",
		})

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "data": [
    {
      "attributes": {
        "defaultLanguageTag": "ja-jp",
        "name": "Japan",
        "supportedLanguageTags": [
          "ja-jp",
          "en-us"
        ]
      },
      "href": "/v1/storefronts/jp",
      "id": "jp",
      "type": "storefronts"
    },
    {
      "attributes": {
        "defaultLanguageTag": "en-gb",
        "name": "Taiwan",
        "supportedLanguageTags": [
          "en-us",
          "zh-tw"
        ]
      },
      "href": "/v1/storefronts/tw",
      "id": "tw",
      "type": "storefronts"
    }
  ]
}`))
	})

	got, _, err := client.Storefront.GetByIds(context.Background(), []string{"jp", "tw"}, nil)
	if err != nil {
		t.Errorf("Storefront.GetByIds returned error: %v", err)
	}

	want := &Storefronts{
		Data: []Storefront{
			{
				Attributes: Attributes{
					DefaultLanguageTag: "ja-jp",
					Name:               "Japan",
					SupportedLanguageTags: []string{
						"ja-jp",
						"en-us",
					},
				},
				Href: "/v1/storefronts/jp",
				Id:   "jp",
				Type: "storefronts",
			},
			{
				Attributes: Attributes{
					DefaultLanguageTag: "en-gb",
					Name:               "Taiwan",
					SupportedLanguageTags: []string{
						"en-us",
						"zh-tw",
					},
				},
				Href: "/v1/storefronts/tw",
				Id:   "tw",
				Type: "storefronts",
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Storefront.GetByIds = %+v, want %+v", got, want)
	}
}

func TestStorefrontsService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/storefronts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"offset": "5",
			"limit":  "2",
			"l":      "en-gb",
		})

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
  "data": [
    {
      "attributes": {
        "defaultLanguageTag": "en-gb",
        "name": "Taiwan",
        "supportedLanguageTags": [
          "en-us",
          "zh-tw"
        ]
      },
      "href": "/v1/storefronts/tw",
      "id": "tw",
      "type": "storefronts"
    },
    {
      "attributes": {
        "defaultLanguageTag": "en-gb",
        "name": "Micronesia, Federated States of",
        "supportedLanguageTags": [
          "en-gb"
        ]
      },
      "href": "/v1/storefronts/fm",
      "id": "fm",
      "type": "storefronts"
    }
  ],
  "next": "/v1/storefronts?offset=7&limit=2&l=en-gb"
}`))
	})

	got, _, err := client.Storefront.GetAll(
		context.Background(),
		&PageOptions{Limit: 2, Offset: 5, Options: Options{Language: "en-gb"}},
	)
	if err != nil {
		t.Errorf("Storefront.GetAll returned error: %v", err)
	}

	want := &Storefronts{
		Data: []Storefront{
			{
				Attributes: Attributes{
					DefaultLanguageTag: "en-gb",
					Name:               "Taiwan",
					SupportedLanguageTags: []string{
						"en-us",
						"zh-tw",
					},
				},
				Href: "/v1/storefronts/tw",
				Id:   "tw",
				Type: "storefronts",
			},
			{
				Attributes: Attributes{
					DefaultLanguageTag: "en-gb",
					Name:               "Micronesia, Federated States of",
					SupportedLanguageTags: []string{
						"en-gb",
					},
				},
				Href: "/v1/storefronts/fm",
				Id:   "fm",
				Type: "storefronts",
			},
		},
		Next: "/v1/storefronts?offset=7&limit=2&l=en-gb",
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Storefront.GetAll = %+v, want %+v", got, want)
	}
}
