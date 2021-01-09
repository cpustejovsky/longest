package longest_test

import (
	"reflect"
	"testing"

	"github.com/cpustejovsky/longest"
)

func TestCompare(t *testing.T) {
	tests := []struct {
		name  string
		first string
		last  string
		want  bool
	}{
		{"should return true", "world", "dog", true},
		{"should return false", "hello", "world", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longest.Compare(tt.first, tt.last)
			if got != tt.want {
				t.Errorf("\ngot:\n%v\nwanted:\n%v\n", got, tt.want)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name string
		list []string
		word string
		want bool
	}{
		{
			"should return true",
			[]string{"badger", "raccoon", "newt", "tiger", "rat", "termite"},
			"badger",
			true,
		},
		{
			"should return false",
			[]string{"badger", "raccoon", "newt", "tiger", "rat", "termite"},
			"world",
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longest.Contains(tt.list, tt.word)
			if got != tt.want {
				t.Errorf("\ngot:\n%v\nwanted:\n%v\n", got, tt.want)
			}
		})
	}

}

func TestFindLongest(t *testing.T) {
	given := []string{"raccoon", "wildebeest", "newt", "termite", "tiger", "rat", "badger"}
	want1 := []string{"badger", "raccoon", "newt", "tiger", "rat", "termite"}
	want2 := []string{"badger", "rat", "tiger", "raccoon", "newt", "termite"}

	got := longest.FindLongest(given)

	if !reflect.DeepEqual(got, want1) && !reflect.DeepEqual(got, want2) {
		t.Errorf("\nwith list:\n%v\ngot:\n%v\nwanted:\n%v\nor\n%v", given, got, want1, want2)
	}
}
