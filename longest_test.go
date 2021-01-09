package longest_test

import (
	"reflect"
	"testing"

	"github.com/cpustejovsky/longest"
)

func TestFindLongest(t *testing.T) {
	given := []string{"raccoon", "wildebeest", "newt", "termite", "tiger", "rat", "badger"}
	want1 := []string{"badger", "raccoon", "newt", "tiger", "rat", "termite"}
	want2 := []string{"badger", "rat", "tiger", "raccoon", "newt", "termite"}

	got := longest.FindLongest(given)

	if !reflect.DeepEqual(got, want1) && !reflect.DeepEqual(got, want2) {
		t.Errorf("\ngot:\n%v\n", got)
	}
}
