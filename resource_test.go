package applemusic

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestResource_Type(t *testing.T) {
	resource := Resource{[]byte(`{"type": "songs"}`)}

	if got, want := resource.Type(), "songs"; got != want {
		t.Errorf("Resource Type is %v, want %v", got, want)
	}
}

func TestResource_Parse(t *testing.T) {
	raw := []byte(`{"type": "songs"}`)
	var resource Resource
	if err := json.Unmarshal(raw, &resource); err != nil {
		t.Fatalf("Unmarshal Resource returned error: %v", err)
	}

	want := &Song{Type: "songs"}
	got, err := resource.Parse()
	if err != nil {
		t.Fatalf("Resource.Parse returned unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Resource.Parse returned %+v, want %+v", got, want)
	}
}
