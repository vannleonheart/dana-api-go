package dana

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/vannleonheart/goutil"
	"net/http"
)

func (c *Client) CreatePaymentRedirect(data *CreatePaymentRedirectRequest) (interface{}, error) {
	if err := c.EnsureAccessToken(); err != nil {
		return nil, err
	}

	accessToken := c.accessToken.AccessToken
	data.MerchantId = c.Config.MerchantId
	timestamp := c.getTimestamp()
	requestUrl := fmt.Sprintf("%s/%s", c.Config.BaseUrl, URLPaymentRedirect)
	encodeRequestBody := c.encodeRequestData(data)
	strToSign := fmt.Sprintf("%s:%s:%s:%s:%s", http.MethodPost, fmt.Sprintf("/%s", URLPaymentRedirect), accessToken, encodeRequestBody, timestamp)
	secretKey := []byte(c.Config.ClientSecret)
	hash := hmac.New(sha512.New, secretKey)
	hash.Write([]byte(strToSign))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	requestHeaders := map[string]string{
		"Content-type":  "application/json",
		"Authorization": fmt.Sprintf("Bearer %s", accessToken),
		"X-TIMESTAMP":   timestamp,
		"X-PARTNER-ID":  c.Config.ClientId,
		"X-EXTERNAL-ID": "1234567890",
		"X-SIGNATURE":   signature,
		"CHANNEL-ID":    "95221",
	}

	var result interface{}

	if _, err := goutil.SendHttpPost(requestUrl, data, &requestHeaders, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
