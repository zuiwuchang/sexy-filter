package configure

import (
	"path/filepath"
	"testing"
)

func TestConfigure(t *testing.T) {
	appPath, e := filepath.Abs("../configure")
	if e != nil {
		t.Fatal(e)
	}

	Init(appPath)
}
