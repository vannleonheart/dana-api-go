package dana

import (
	"fmt"
	"github.com/vannleonheart/goutil"
	"net/http"
	"strings"
)

func (c *Client) GetB2BAccessToken() (*GetB2BAccessTokenResponse, error) {
	timestamp := c.getTimestamp()
	strToSign := fmt.Sprintf("%s|%s", c.Config.ClientId, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "GetB2BAccessToken",
			"message":      "error when sign request",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

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

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLAccessToken)

	var result GetB2BAccessTokenResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "GetB2BAccessToken",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "GetB2BAccessToken",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &result, nil
}

func (c *Client) EnsureB2BAccessToken() error {
	if c.b2bAccessToken == nil {
		accessTokenResponse, err := c.GetB2BAccessToken()
		if err != nil {
			return err
		}
		c.SetB2BAccessToken(accessTokenResponse.AccessToken)
	}

	return nil
}

func (c *Client) GetCustomerAuthCode(scopes *[]string, redirectUrl string) (*string, *string, error) {
	externalId := c.getRequestId(nil)
	state := GenerateRequestId(5, 32, goutil.NumCharset)
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
		"state":       state,
	}

	qs, err := goutil.GenerateQueryString(queryParams)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function": "GetCustomerAuthCode",
			"message":  "error when generate query string",
			"error":    err,
			"params":   queryParams,
		})

		return nil, nil, err
	}

	requestUrl := fmt.Sprintf("%s/%s?%s", c.Config.WebUrl, URLGetAuthCode, *qs)

	c.log("debug", map[string]interface{}{
		"function": "GetCustomerAuthCode",
		"url":      requestUrl,
	})

	return &externalId, &requestUrl, nil
}

func (c *Client) CustomerApplyToken(token string, granType *string) (*CustomerApplyTokenResponse, error) {
	timestamp := c.getTimestamp()
	strToSign := fmt.Sprintf("%s|%s", c.Config.ClientId, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "CustomerApplyToken",
			"message":      "error when sign request",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

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

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLApplyToken)

	var result CustomerApplyTokenResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "CustomerApplyToken",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "CustomerApplyToken",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &result, nil
}

func (c *Client) CustomerApplyOTT(customerAccessToken *AccessToken) (*string, *CustomerApplyOTTResponse, error) {
	timestamp := c.getTimestamp()
	externalId := c.getRequestId(nil)
	accessToken := c.getCustomerAccessToken(customerAccessToken)

	requestBody := map[string]interface{}{
		"userResources": []string{"OTT"},
		"additionalInfo": map[string]string{
			"accessToken": accessToken.AccessToken,
		},
	}

	encodeRequestBody := EncodeRequestBody(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLApplyOTT), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "CustomerApplyOTT",
			"message":      "error when sign request",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

		return nil, nil, err
	}

	requestHeaders := map[string]string{
		"Content-type":           "application/json",
		"Authorization-Customer": fmt.Sprintf("%s %s", accessToken.TokenType, accessToken.AccessToken),
		"X-TIMESTAMP":            timestamp,
		"X-SIGNATURE":            *signature,
		"ORIGIN":                 c.getOrigin(),
		"X-PARTNER-ID":           c.Config.ClientId,
		"X-EXTERNAL-ID":          externalId,
		"X-IP-ADDRESS":           c.getIpAddress(),
		"X-DEVICE-ID":            c.getDeviceId(),
		"X-LATITUDE":             c.getLatitude(),
		"X-LONGITUDE":            c.getLongitude(),
		"CHANNEL-ID":             "95221",
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLApplyOTT)

	var result CustomerApplyOTTResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "CustomerApplyOTT",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "CustomerApplyOTT",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &externalId, &result, nil
}

func (c *Client) CustomerUnbindAccount(customerAccessToken *AccessToken) (*string, *GeneralResponse, error) {
	timestamp := c.getTimestamp()
	externalId := c.getRequestId(nil)
	accessToken := c.getCustomerAccessToken(customerAccessToken)

	requestBody := map[string]interface{}{
		"merchantId": c.Config.MerchantId,
		"additionalInfo": map[string]string{
			"accessToken": accessToken.AccessToken,
		},
	}

	encodeRequestBody := EncodeRequestBody(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLUnbindToken), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "CustomerUnbindAccount",
			"message":      "error when sign request",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

		return nil, nil, err
	}

	requestHeaders := map[string]string{
		"Content-type":           "application/json",
		"Authorization-Customer": fmt.Sprintf("%s %s", accessToken.TokenType, accessToken.AccessToken),
		"X-TIMESTAMP":            timestamp,
		"X-SIGNATURE":            *signature,
		"ORIGIN":                 c.getOrigin(),
		"X-PARTNER-ID":           c.Config.ClientId,
		"X-EXTERNAL-ID":          externalId,
		"X-IP-ADDRESS":           c.getIpAddress(),
		"X-DEVICE-ID":            c.getDeviceId(),
		"X-LATITUDE":             c.getLatitude(),
		"X-LONGITUDE":            c.getLongitude(),
		"CHANNEL-ID":             "95221",
	}

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLUnbindToken)

	var result GeneralResponse

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "CustomerUnbindAccount",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "CustomerUnbindAccount",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &externalId, &result, nil
}
