package js

import (
	"fmt"
	"testing"
)

func TestFile0(t *testing.T) {
	msg, e := testFile(
		"../plugins-js/t66y-4.js",
		"../test-js/a.html",
	)
	if e != nil {
		t.Fatal(e)
	}
	fmt.Println(msg)
}
