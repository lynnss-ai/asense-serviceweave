package netutil

import (
	"errors"
	"net/url"
)

// UrlEncode URL编码
// apiUrl: URL地址
// params: 参数
// return: 编码后的URL地址
func UrlEncode(apiUrl string, params map[string]string) (string, error) {
	_url, err := url.Parse(apiUrl)
	if err != nil {
		return "", errors.New("URL parse error")
	}
	q := _url.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	_url.RawQuery = q.Encode()
	r, _ := url.QueryUnescape(_url.String())
	return r, nil
}
