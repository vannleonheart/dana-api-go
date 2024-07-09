package dana

import (
	"fmt"
	"github.com/vannleonheart/goutil"
	"net/http"
	"strings"
)

func (c *Client) GetAccessToken() (*AccessTokenResponse, error) {
	timestamp := c.getTimestamp()
	strToSign := fmt.Sprintf("%s|%s", c.Config.ClientId, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		return nil, err
	}

	requestHeaders := map[string]string{
		"Content-type": "application/json",
		"X-TIMESTAMP":  timestamp,
		"X-CLIENT-KEY": c.Config.ClientId,
		"X-SIGNATURE":  *signature,
	}

	requestBody := map[string]interface{}{
		"grantType":      "client_credentials",
		"additionalInfo": map[string]string{},
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.BaseUrl, URLAccessToken)

	var result AccessTokenResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) EnsureAccessToken() error {
	if c.accessToken == nil {
		accessToken, err := c.GetAccessToken()
		if err != nil {
			return err
		}
		c.SetAccessToken(accessToken.AccessToken)
	}

	return nil
}

func (c *Client) GetAuthCode(externalId, redirectUrl string, scopes *[]string) (*string, error) {
	rnd := goutil.NewRandomString("")

	currentScopes := []string{"PUBLIC_ID", "QUERY_BALANCE", "MINI_DANA"}
	if scopes != nil {
		currentScopes = append(currentScopes, *scopes...)
	}

	queryParams := map[string]interface{}{
		"partnerId":   c.Config.ClientId,
		"timestamp":   c.getTimestamp(),
		"externalId":  externalId,
		"channelId":   "DANAID",
		"merchantId":  c.Config.MerchantId,
		"scopes":      strings.Join(currentScopes, ","),
		"redirectUrl": redirectUrl,
		"state":       rnd.GenerateRange(5, 32),
	}

	qs, err := goutil.GenerateQueryString(queryParams)
	if err != nil {
		return nil, err
	}

	requestUrl := fmt.Sprintf("%s/%s?%s", c.Config.WebUrl, URLGetAuthCode, *qs)

	return &requestUrl, nil
}

func (c *Client) ApplyToken(token string, granType *string) (*ApplyTokenResponse, error) {
	timestamp := c.getTimestamp()
	strToSign := fmt.Sprintf("%s|%s", c.Config.ClientId, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		return nil, err
	}

	requestHeaders := map[string]string{
		"Content-type": "application/json",
		"X-TIMESTAMP":  timestamp,
		"X-CLIENT-KEY": c.Config.ClientId,
		"X-SIGNATURE":  *signature,
	}

	currentGrantType := "AUTHORIZATION_CODE"
	if granType != nil {
		currentGrantType = *granType
	}

	authCode := ""
	refreshToken := ""

	switch currentGrantType {
	case "AUTHORIZATION_CODE":
		authCode = token
	case "REFRESH_TOKEN":
		refreshToken = token
	}

	requestBody := map[string]interface{}{
		"grantType":    currentGrantType,
		"authCode":     authCode,
		"refreshToken": refreshToken,
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.BaseUrl, URLApplyToken)

	var result ApplyTokenResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) ApplyOTT(accessToken, externalId, originHost, ipAddress string) (*ApplyOTTResponse, error) {
	timestamp := c.getTimestamp()

	requestBody := map[string]interface{}{
		"userResources": []string{"OTT"},
		"additionalInfo": map[string]string{
			"accessToken": accessToken,
		},
	}
	encodeRequestBody := c.encodeRequestData(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLApplyOTT), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		return nil, err
	}

	requestHeaders := map[string]string{
		"Content-type":           "application/json",
		"Authorization-Customer": fmt.Sprintf("Bearer %s", accessToken),
		"X-TIMESTAMP":            timestamp,
		"X-SIGNATURE":            *signature,
		"ORIGIN":                 originHost,
		"X-PARTNER-ID":           c.Config.ClientId,
		"X-EXTERNAL-ID":          externalId,
		"X-IP-ADDRESS":           ipAddress,
		"X-DEVICE-ID":            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
		"X-LATITUDE":             DefaultLatitude,
		"X-LONGITUDE":            DefaultLongitude,
		"CHANNEL-ID":             "95221",
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.BaseUrl, URLApplyOTT)

	var result ApplyOTTResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UnbindAccount(accessToken, externalId, originHost, ipAddress string) (*GeneralResponse, error) {
	timestamp := c.getTimestamp()

	requestBody := map[string]interface{}{
		"merchantId": c.Config.MerchantId,
		"additionalInfo": map[string]string{
			"accessToken": accessToken,
		},
	}

	encodeRequestBody := c.encodeRequestData(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLUnbindToken), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		return nil, err
	}

	requestHeaders := map[string]string{
		"Content-type":           "application/json",
		"Authorization-Customer": fmt.Sprintf("Bearer %s", accessToken),
		"X-TIMESTAMP":            timestamp,
		"X-SIGNATURE":            *signature,
		"ORIGIN":                 originHost,
		"X-PARTNER-ID":           c.Config.ClientId,
		"X-EXTERNAL-ID":          externalId,
		"X-IP-ADDRESS":           ipAddress,
		"X-DEVICE-ID":            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
		"X-LATITUDE":             DefaultLatitude,
		"X-LONGITUDE":            DefaultLongitude,
		"CHANNEL-ID":             "95221",
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.BaseUrl, URLUnbindToken)

	var result GeneralResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
