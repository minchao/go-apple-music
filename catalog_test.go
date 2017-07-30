package applemusic

import (
	"encoding/json"
	"testing"
)

// See json.RawMessage
func TestTrack(t *testing.T) {
	// TODO(rsc): Should not need the * in *RawMessage
	var data struct {
		X  float64
		Id *Track
		Y  float32
	}
	const raw = `["\u0056",null]`
	const msg = `{"X":0.1,"Id":["\u0056",null],"Y":0.2}`
	err := json.Unmarshal([]byte(msg), &data)
	if err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if string([]byte(*data.Id)) != raw {
		t.Fatalf("Raw mismatch: have %#q want %#q", []byte(*data.Id), raw)
	}
	b, err := json.Marshal(&data)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(b) != msg {
		t.Fatalf("Marshal: have %#q want %#q", b, msg)
	}
}

// See json.RawMessage
func TestNullTrack(t *testing.T) {
	// TODO(rsc): Should not need the * in *RawMessage
	var data struct {
		X  float64
		Id *Track
		Y  float32
	}
	data.Id = new(Track)
	const msg = `{"X":0.1,"Id":null,"Y":0.2}`
	err := json.Unmarshal([]byte(msg), &data)
	if err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if data.Id != nil {
		t.Fatalf("Raw mismatch: have non-nil, want nil")
	}
	b, err := json.Marshal(&data)
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}
	if string(b) != msg {
		t.Fatalf("Marshal: have %#q want %#q", b, msg)
	}
}

func TestTrack_Type(t *testing.T) {
	track := Track(`{"type": "songs"}`)

	if got, want := track.Type(), "songs"; got != want {
		t.Errorf("Track Type is %v, want %v", got, want)
	}
}
