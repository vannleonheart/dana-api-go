package dana

import (
	"fmt"
	"github.com/vannleonheart/goutil"
	"net/http"
)

func (c *Client) CustomerBalanceInquiry(requestId *string, customerAccessToken *AccessToken) (*interface{}, error) {
	timestamp := c.getTimestamp()
	externalId := c.getRequestId(requestId)
	accessToken := c.getCustomerAccessToken(customerAccessToken)

	requestBody := map[string]interface{}{
		"additionalInfo": map[string]string{
			"accessToken": accessToken.AccessToken,
		},
	}

	encodeRequestBody := EncodeRequestBody(requestBody)
	strToSign := fmt.Sprintf("%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLBalanceInquiry), encodeRequestBody, timestamp)
	signature, err := c.sign(strToSign)
	if err != nil {
		c.log("error", map[string]interface{}{
			"function":     "CustomerBalanceInquiry",
			"message":      "error when sign request body",
			"error":        err.Error(),
			"stringToSign": strToSign,
		})

		return nil, err
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

	requestUrl := fmt.Sprintf("%s/%s", c.Config.ApiUrl, URLBalanceInquiry)

	var result interface{}

	if _, err = goutil.SendHttpPost(requestUrl, requestBody, &requestHeaders, &result, nil); err != nil {
		c.log("error", map[string]interface{}{
			"function": "CustomerBalanceInquiry",
			"message":  "error when send http post",
			"error":    err,
			"url":      requestUrl,
			"headers":  requestHeaders,
			"body":     requestBody,
		})

		return nil, err
	}

	c.log("debug", map[string]interface{}{
		"function": "CustomerBalanceInquiry",
		"result":   result,
		"url":      requestUrl,
		"headers":  requestHeaders,
		"body":     requestBody,
	})

	return &result, nil
}
