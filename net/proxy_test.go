package net

import (
	"testing"
)

func TestProxy0(t *testing.T) {
	//socks5
	e := TestProxy(ProxySocks5, "127.0.0.1:1080", "", "", "https://www.google.com")
	if e != nil {
		t.Fatal(e)
	}
}
