package applemusic

import (
	"context"
	"net/http"
	"reflect"
	"testing"
)

func TestCatalogService_GetAppleCurator(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/apple-curators/976439526", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		w.WriteHeader(http.StatusOK)
		w.Write(appleCuratorsJSON)
	})

	got, _, err := client.Catalog.GetAppleCurator(context.Background(), "us", "976439526", nil)
	if err != nil {
		t.Errorf("Catalog.GetAppleCurator returned error: %v", err)
	}
	if want := appleCurators; !reflect.DeepEqual(got, want) {
		t.Errorf("Catalog.GetAppleCurator = %+v, want %+v", got, want)
	}
}

func TestCatalogService_GetAppleCuratorsByIds(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/catalog/us/apple-curators", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"ids": "976439526,1017168810",
		})

		w.WriteHeader(http.StatusOK)
	})

	_, _, err := client.Catalog.GetAppleCuratorsByIds(context.Background(), "us", []string{"976439526", "1017168810"}, nil)
	if err != nil {
		t.Errorf("Catalog.GetAppleCuratorsByIds returned error: %v", err)
	}
}

var appleCuratorsJSON = []byte(`{
    "data": [
        {
            "id": "976439526",
            "type": "apple-curators",
            "href": "/v1/catalog/us/apple-curators/976439526",
            "attributes": {
                "url": "https://itunes.apple.com/us/curator/apple-music-alternative/id976439526",
                "name": "Apple Music Alternative",
                "artwork": {
                    "width": 1080,
                    "height": 1080,
                    "url": "https://is3-ssl.mzstatic.com/image/thumb/Features5/v4/16/e3/e7/16e3e76f-c06e-8479-c8ea-b2df32661ef4/source/{w}x{h}bb.jpg",
                    "bgColor": "181f2e",
                    "textColor1": "ecc6bc",
                    "textColor2": "e1bdb5",
                    "textColor3": "c1a59f",
                    "textColor4": "b99e9a"
                },
                "editorialNotes": {
                    "standard": "In an age where alternative rock bands fill stadiums and ascend the pop charts, it begs the question: alternative to what? Early on, the alternative movement was a reaction to the commercial excesses of mainstream rock. Alt-rock instead brought quirky hooks, a do-it-yourself ethos, deeply personal songwriting, and genre-bending adventures to audiences hungry for something different. Although it truly exploded in the early ’90s, the roots of alternative rock started with the punk revolt of the late ’70s, when bands like the Ramones and the Sex Pistols proved that just about anyone could get up onstage or make a record. Throughout the ’80s, an international network of under-the-radar bands developed, nurtured by college radio DJs and small clubs. While hardcore kept the traditional loud-and-fast sound of punk alive, many newer bands had their own distinctive styles: R.E.M.'s jangling folk-influenced rock, Sonic Youth's dissonant noise, The Cure's epic gloom, Pixies' whisper-to-a-scream dynamics, New Order's electronic grooves. \n<br /><br />\nEventually, these bands were dubbed \"alternative rock,\" thanks to their left-of-center sounds and attitudes. By the early ’90s, though, grunge bands like Nirvana and Pearl Jam were combining punk’s raw energy with classic hard-rock hooks and invading the pop charts. Suddenly, other alternative heroes like Red Hot Chili Peppers and Soundgarden found massive audiences. Over the next decade, alternative bands of various subgenres introduced a whole generation of young rockers to punk (Green Day), hip-hop (Rage Against the Machine), industrial (Nine Inch Nails), art rock (Radiohead), power pop (Weezer), psychedelia (The Flaming Lips), metal (Tool), the British Invasion (Oasis), electronic music (The Prodigy), and much more. By the 21st century, alternative rock had grown popular enough to allow bands like Foo Fighters and Coldplay to sell out stadiums in minutes. At the same time, the anything-goes spirit of alternative rock remained alive and well, with newer bands embracing garage rock (The White Stripes), emo (Paramore), and New Wave (The Killers).",
                    "short": "In the '90s, hair metal and hip-hop dominated. Then everything changed . . . "
                }
            },
            "relationships": {
                "playlists": {
                    "data": [
                        {
                            "id": "pl.705e4024180a4b8ab0b4700b888c66ee",
                            "type": "playlists",
                            "href": "/v1/catalog/us/playlists/pl.705e4024180a4b8ab0b4700b888c66ee"
                        }
                    ],
                    "href": "/v1/catalog/us/apple-curators/976439526/playlists",
                    "next": "/v1/catalog/us/apple-curators/976439526/playlists?offset=10"
                }
            }
        }
    ]
}`)

var appleCurators = &AppleCurators{
	Data: []AppleCurator{
		{
			Id:   "976439526",
			Type: "apple-curators",
			Href: "/v1/catalog/us/apple-curators/976439526",
			Attributes: AppleCuratorAttributes{
				URL:  "https://itunes.apple.com/us/curator/apple-music-alternative/id976439526",
				Name: "Apple Music Alternative",
				Artwork: Artwork{
					Width:      1080,
					Height:     1080,
					URL:        "https://is3-ssl.mzstatic.com/image/thumb/Features5/v4/16/e3/e7/16e3e76f-c06e-8479-c8ea-b2df32661ef4/source/{w}x{h}bb.jpg",
					BgColor:    "181f2e",
					TextColor1: "ecc6bc",
					TextColor2: "e1bdb5",
					TextColor3: "c1a59f",
					TextColor4: "b99e9a",
				},
				EditorialNotes: &EditorialNotes{
					Standard: "In an age where alternative rock bands fill stadiums and ascend the pop charts, it begs the question: alternative to what? Early on, the alternative movement was a reaction to the commercial excesses of mainstream rock. Alt-rock instead brought quirky hooks, a do-it-yourself ethos, deeply personal songwriting, and genre-bending adventures to audiences hungry for something different. Although it truly exploded in the early ’90s, the roots of alternative rock started with the punk revolt of the late ’70s, when bands like the Ramones and the Sex Pistols proved that just about anyone could get up onstage or make a record. Throughout the ’80s, an international network of under-the-radar bands developed, nurtured by college radio DJs and small clubs. While hardcore kept the traditional loud-and-fast sound of punk alive, many newer bands had their own distinctive styles: R.E.M.'s jangling folk-influenced rock, Sonic Youth's dissonant noise, The Cure's epic gloom, Pixies' whisper-to-a-scream dynamics, New Order's electronic grooves. \n<br /><br />\nEventually, these bands were dubbed \"alternative rock,\" thanks to their left-of-center sounds and attitudes. By the early ’90s, though, grunge bands like Nirvana and Pearl Jam were combining punk’s raw energy with classic hard-rock hooks and invading the pop charts. Suddenly, other alternative heroes like Red Hot Chili Peppers and Soundgarden found massive audiences. Over the next decade, alternative bands of various subgenres introduced a whole generation of young rockers to punk (Green Day), hip-hop (Rage Against the Machine), industrial (Nine Inch Nails), art rock (Radiohead), power pop (Weezer), psychedelia (The Flaming Lips), metal (Tool), the British Invasion (Oasis), electronic music (The Prodigy), and much more. By the 21st century, alternative rock had grown popular enough to allow bands like Foo Fighters and Coldplay to sell out stadiums in minutes. At the same time, the anything-goes spirit of alternative rock remained alive and well, with newer bands embracing garage rock (The White Stripes), emo (Paramore), and New Wave (The Killers).",
					Short:    "In the '90s, hair metal and hip-hop dominated. Then everything changed . . . ",
				},
			},
			Relationships: AppleCuratorRelationships{
				Playlists: Playlists{
					Data: []Playlist{
						{
							Id:   "pl.705e4024180a4b8ab0b4700b888c66ee",
							Type: "playlists",
							Href: "/v1/catalog/us/playlists/pl.705e4024180a4b8ab0b4700b888c66ee",
						},
					},
					Href: "/v1/catalog/us/apple-curators/976439526/playlists",
					Next: "/v1/catalog/us/apple-curators/976439526/playlists?offset=10",
				},
			},
		},
	},
}
