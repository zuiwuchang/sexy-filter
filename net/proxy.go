package net

import (
	"errors"
	kStrings "github.com/zuiwuchang/king-go/strings"
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
	c, e := createSocks5Client(addr, user, pwd)
	if e != nil {
		return e
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
func createSocks5Client(addr, user, pwd string) (c *http.Client, e error) {
	var auth *proxy.Auth
	if user != "" {
		auth = &proxy.Auth{User: user, Password: pwd}
	}
	var dialer proxy.Dialer
	dialer, e = proxy.SOCKS5("tcp", addr, auth, proxy.Direct)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}
	httpTransport := &http.Transport{}
	httpTransport.Dial = dialer.Dial

	c = &http.Client{
		Transport: httpTransport,
	}
	return
}
func createHttpClient(addr string) (c *http.Client, e error) {
	var proxyUrl *url.URL
	proxyUrl, e = url.Parse(addr)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}
	c = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	return
}
func GetUrl(style int, addr, user, pwd, url string) (msg string, e error) {
	var c *http.Client
	switch style {
	case 0:
		c = &http.Client{}
	case ProxySocks5:
		c, e = createSocks5Client(addr, user, pwd)
	case ProxyHttp:
		addr = "http://" + addr
		c, e = createHttpClient(addr)
	case ProxyHttps:
		addr = "https://" + addr
		c, e = createHttpClient(addr)
	default:
		e = errors.New("unkonw proxy type")
		if log.Error != nil {
			log.Error.Println(e)
		}
	}
	if e != nil {
		return
	}
	var r *http.Response
	r, e = c.Get(url)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}

	var b []byte
	b, e = ioutil.ReadAll(r.Body)
	if e != nil {
		if log.Warn != nil {
			log.Warn.Println(e)
		}
		return
	}
	msg = kStrings.BytesToString(b)
	return
}
