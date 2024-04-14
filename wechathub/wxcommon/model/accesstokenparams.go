package model

type (
	AccessTokenRequest struct {
		GrantType    string `json:"grant_type"`    // client_credential
		AppID        string `json:"appid"`         // appid
		AppSecret    string `json:"secret"`        // secret
		ForceRefresh bool   `json:"force_refresh"` // 是否强制刷新
	}

	AccessTokenResponse struct {
		CommonError
		AccessToken string `json:"access_token,omitempty"` // 获取到的凭证
		ExpiresIn   int    `json:"expires_in,omitempty"`   // 凭证有效时间，单位：秒 <目前是7200秒之内的值>
	}
)
