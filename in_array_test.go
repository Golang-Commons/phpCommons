package phpCommons

import (
	"testing"
)

func Test_in_array(t *testing.T) {
	a := "aaaaa"
	c := "aaaaac"
	b := []string{"aaaaa", "vvvv", "cc"}
	in, err := in_array(a, b)
	if err != nil || !in {
		t.Error(err)
	}
	in, err = in_array(c, b)
	if err != nil || in {
		t.Error(err)
	}
}
