package js

import (
	"fmt"
	//"sexy-filter/net"
	"path/filepath"
	"testing"
)

func TestFile0(t *testing.T) {
	/**
	_, e := testFile(
		"../plugins-js/t66y-4.js",
		"../test-js/a.html",
	)
	if e != nil {
		t.Fatal(e)
	}
	/**/
}
func TestUrl0(t *testing.T) {
	/**
	msg, e := testUrl(
		net.ProxySocks5, "127.0.0.1:1080",
		"", "",
		"../plugins-js/t66y-5.js",
	)
	if e != nil {
		t.Fatal(e)
	}
	fmt.Println(msg)
	/**/
}

func TestSingle(t *testing.T) {
	str, _ := filepath.Abs("../")
	WorkDir = str
	fmt.Println(WorkDir)
	InitSingle()

	single := GetSingle()
	fmt.Println(single.GetPlugins())
}
