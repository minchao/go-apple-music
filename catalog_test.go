package applemusic

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestTrack_Type(t *testing.T) {
	track := Track{[]byte(`{"type": "songs"}`)}

	if got, want := track.Type(), "songs"; got != want {
		t.Errorf("Track Type is %v, want %v", got, want)
	}
}

func TestTrack_Parse(t *testing.T) {
	raw := []byte(`{"type": "songs"}`)
	var track Track
	if err := json.Unmarshal(raw, &track); err != nil {
		t.Fatalf("Unmarshal Track returned error: %v", err)
	}

	want := &Song{Type: "songs"}
	got, err := track.Parse()
	if err != nil {
		t.Fatalf("Track.Parse returned unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Track.Parse returned %+v, want %+v", got, want)
	}
}
