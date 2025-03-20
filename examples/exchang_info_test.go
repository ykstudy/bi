package main

import (
	"crypto/tls"
	"example/config"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestHttpApi(t *testing.T) {
	res, err := http.Get("https://api.binance.com/api/v3/exchangeInfo?symbol=BTCUSDT")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func TestHttpProxyApi(t *testing.T) {
	conf := config.NewConfig()
	proxy, err := url.Parse(conf.B.ProxyUrl)
	if err != nil {
		t.Fatal(err)
	}
	tr := &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{
		Transport: tr,
	}
	res, err := httpClient.Get("https://api.binance.com/api/v3/exchangeInfo?symbol=BTCUSDT")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("unexpected status code: %d", res.StatusCode)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}
