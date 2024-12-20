package stringsplit

import (
	"reflect"
	"testing"
)

func TestStringSplit(t *testing.T) {
	got := StringSplit("ab")
	want := []string{"ab"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestStringSplit_3Characters(t *testing.T) {
	got := StringSplit("abc")
	want := []string{"ab", "c_"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestStringSplit_4Characters(t *testing.T) {
	got := StringSplit("abcd")
	want := []string{"ab", "cd"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q, want %q", got, want)
	}
}