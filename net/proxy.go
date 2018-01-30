package net

import (
	"errors"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"net/http"
	"net/url"
	"sexy-filter/log"
)

const (
	ProxyNone = iota
	ProxyHttp
	ProxyHttps
	ProxySocks5
)

func TestProxy(style int, addr, user, pwd, testUrl string) error {
	if style == ProxySocks5 {
		//socks5
		return testSocks5(addr, user, pwd, testUrl)
	} else if style == ProxyHttp {
		//http
		addr = "http://" + addr
		return testHttp(addr, testUrl)
	} else if style == ProxyHttps {
		//https
		addr = "https://" + addr
		return testHttp(addr, testUrl)
	}
	return errors.New("unkonw proxy type")
}
func testHttp(addr, testUrl string) error {
	proxyUrl, e := url.Parse(addr)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return e
	}

	c := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	r, e := c.Get(testUrl)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return e
	}

	_, e = ioutil.ReadAll(r.Body)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return e
	}
	return nil
}
func testSocks5(addr, user, pwd, testUrl string) error {
	var auth *proxy.Auth
	if user != "" {
		auth = &proxy.Auth{User: user, Password: pwd}
	}
	dialer, e := proxy.SOCKS5("tcp", addr, auth, proxy.Direct)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return e
	}
	httpTransport := &http.Transport{}
	httpTransport.Dial = dialer.Dial

	c := &http.Client{
		Transport: httpTransport,
	}
	r, e := c.Get(testUrl)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return e
	}

	_, e = ioutil.ReadAll(r.Body)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return e
	}
	return nil
}
