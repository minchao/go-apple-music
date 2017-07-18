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
			"ids": "jp",
		})

		w.WriteHeader(http.StatusOK)
		w.Write(storefrontsJSON)
	})

	got, _, err := client.Storefront.GetByIds(context.Background(), []string{"jp"}, nil)
	if err != nil {
		t.Errorf("Storefront.GetByIds returned error: %v", err)
	}

	if want := wantStorefronts; !reflect.DeepEqual(got, want) {
		t.Errorf("Storefront.GetByIds = %+v, want %+v", got, want)
	}
}

func TestStorefrontsService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/storefronts", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(storefrontsJSON)
	})

	got, _, err := client.Storefront.GetAll(context.Background(), nil)
	if err != nil {
		t.Errorf("Storefront.GetAll returned error: %v", err)
	}

	if want := wantStorefronts; !reflect.DeepEqual(got, want) {
		t.Errorf("Storefront.GetAll = %+v, want %+v", got, want)
	}
}

var storefrontsJSON = []byte(`{
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

var wantStorefronts = &Storefronts{
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
