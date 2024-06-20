package dana

import (
	"fmt"
	"github.com/vannleonheart/goutil"
)

func (c *Client) GetAccessToken() (*AccessTokenResponse, error) {
	timestamp := c.getTimestamp()
	strToSign := fmt.Sprintf("%s|%s", c.Config.ClientId, timestamp)
	requestUrl := fmt.Sprintf("%s/%s", c.Config.BaseUrl, URLAccessToken)

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
