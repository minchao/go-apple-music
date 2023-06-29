package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestMeService_GetStorefront(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/me/storefront", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(meStorefrontJSON)
	})

	got, _, err := client.Me.GetStorefront(context.Background(), nil)
	if err != nil {
		t.Errorf("Me.GetStorefront returned error: %v", err)
	}
	if want := meStorefront; !reflect.DeepEqual(got, want) {
		t.Errorf("Me.GetStorefront = %+v, want %+v", got, want)
	}
}

var meStorefrontJSON = []byte(`{
    "data": [
        {
            "id": "us",
            "type": "storefronts",
            "href": "/v1/storefronts/us",
            "attributes": {
                "name": "United States",
                "supportedLanguageTags": [
                    "en-us",
                    "es-mx"
                ],
                "defaultLanguageTag": "en-us"
            }
        }
    ]
}`)

var meStorefront = &Storefronts{
	Data: []Storefront{
		{
			Id:   "us",
			Type: "storefronts",
			Href: "/v1/storefronts/us",
			Attributes: StorefrontAttributes{
				Name: "United States",
				SupportedLanguageTags: []string{
					"en-us",
					"es-mx",
				},
				DefaultLanguageTag: "en-us",
			},
		},
	},
}
