package phpCommons

import (
	"testing"

	"fmt"
)

func Test_array_values(t *testing.T) {
	b := []string{"aaaaa", "vvvv", "cc", "aaaa", "aaaaa", "aaaaa", "cc", "vvvv"}
	fmt.Printf("%#v\n", array_values(b))
	//	c := []string{"aaaa", "aaaaa", "cc", "vvvv"}
}
