package wxcommon

import (
	"asense-serviceweave/commonutil/characterutil"
	"asense-serviceweave/commonutil/netutil"
	"asense-serviceweave/wechathub/wxcommon/core"
	"asense-serviceweave/wechathub/wxcommon/model"
	"encoding/json"
	"errors"
	"github.com/valyala/fasthttp"
	"strconv"
)

// GetAccessToken 获取access_token
// @param appid
// @param secret
// @return *model.AccessTokenResponse
// @return error
func GetAccessToken(appid, secret string) (*model.AccessTokenResponse, error) {
	request := map[string]string{
		"grant_type": "client_credential",
		"appid":      appid,
		"secret":     secret,
	}
	apiUrl, err := netutil.UrlEncode(core.TokenUrl, request)
	if err != nil {
		return nil, err
	}
	status, response, err := fasthttp.Get(nil, apiUrl)
	if err != nil {
		return nil, err
	}
	if status != fasthttp.StatusOK {
		return nil, errors.New(characterutil.StitchingBuilderStr("http status code: ", strconv.Itoa(status)))
	}
	r := new(model.AccessTokenResponse)
	err = json.Unmarshal(response, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetStableAccessToken 获取稳定的access_token
// @param appid
// @param secret
// @param forceRefresh
// @return *model.AccessTokenResponse
// @return error
func GetStableAccessToken(appid, secret string, forceRefresh bool) (*model.AccessTokenResponse, error) {
	requestParam := model.AccessTokenRequest{
		GrantType:    "client_credential",
		AppID:        appid,
		AppSecret:    secret,
		ForceRefresh: forceRefresh,
	}
	requestData, _ := json.Marshal(requestParam)

	request := &fasthttp.Request{}
	request.SetRequestURI(core.StableTokenUrl)
	request.SetBodyRaw(requestData)
	request.Header.SetMethod("POST")
	request.Header.SetContentType("application/json")

	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)
	if err := fasthttp.Do(request, response); err != nil {
		return nil, err
	}
	if response.StatusCode() != fasthttp.StatusOK {
		return nil, errors.New(characterutil.StitchingBuilderStr("http status code: ", strconv.Itoa(response.StatusCode())))
	}

	r := new(model.AccessTokenResponse)
	err := json.Unmarshal(response.Body(), r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
