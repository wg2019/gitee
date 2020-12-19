package main

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type httpClient struct {
	http.Client
}

var client *httpClient

// GetHttpClient 获取http客户端
func GetHttpClient() (*httpClient, error) {
	if client == nil {
		client = new(httpClient)
		jar, err := cookiejar.New(nil)
		if err != nil {
			return nil, err
		}
		reqUrl, err := url.Parse(Domain)
		if err != nil {
			return nil, err
		}
		jar.SetCookies(reqUrl, []*http.Cookie{
			{
				Domain: "gitee.com",
				Name:   "gitee-session-n",
				Value:  Account.Cookie,
			},
		})
		client.Jar = jar
	}
	return client, nil
}

func (c *httpClient) Delete(url, contentType string, param url.Values) (resp *http.Response, err error) {
	req, err := http.NewRequest("DELETE", url, strings.NewReader(param.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-CSRF-Token", Account.Token)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", "*/*")
	return c.Do(req)
}
