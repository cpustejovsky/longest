package longest

import (
	"reflect"
	"testing"

	"github.com/cpustejovsky/gotest/longest"
)

func testFindLongest(t *testing.T) {
	given := []string{"raccoon", "wildebeest", "newt", "termite", "tiger", "rat", "badger"}
	want1 := []string{"badger", "raccoon", "newt", "tiger", "rat", "termite"}
	want2 := []string{"badger", "rat", "tiger", "raccoon", "newt", "termite"}

	got := longest.FindLongest(given)

	if !(reflect.DeepEqual(got, want1) || reflect.DeepEqual(got, want2)) {
		t.Errorf("\ngot:\n%v\n", got)
	}
}
