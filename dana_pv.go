package dana

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/vannleonheart/goutil"
	"strings"
	"time"
)

func (c *Client) getChannelId() string {
	return defaultChannelId
}

func (c *Client) getCustomerAccessToken(accessToken *AccessToken) *AccessToken {
	if accessToken != nil {
		c.SetCustomerAccessToken(accessToken)
	}

	defer c.ClearCustomerAccessToken()

	return c.customerAccessToken
}

func (c *Client) getRequestId(requestId *string) string {
	if requestId != nil {
		c.SetRequestId(*requestId)
	}

	if c.requestId == nil {
		c.SetGeneratedRequestId()
	}

	defer c.ClearRequestId()

	return *c.requestId
}

func (c *Client) getOrigin() string {
	origin := c.Config.Origin
	if c.origin != nil {
		origin = *c.origin
	}

	return origin
}

func (c *Client) getDeviceId() string {
	deviceId := defaultDevideId
	if c.deviceId != nil {
		deviceId = *c.deviceId
	}

	return deviceId
}

func (c *Client) getIpAddress() string {
	ipAddress := c.Config.IpAddress
	if c.ipAddress != nil {
		ipAddress = *c.ipAddress
	}

	return ipAddress
}

func (c *Client) getLatitude() string {
	lat := c.Config.Latitude
	if c.lat != nil {
		lat = *c.lat
	}

	return lat
}

func (c *Client) getLongitude() string {
	lon := c.Config.Longitude
	if c.lon != nil {
		lon = *c.lon
	}

	return lon
}

func (c *Client) getTimestamp() string {
	now := time.Now()
	currentTimezone := defaultTimezone
	configTimezone := strings.TrimSpace(c.Config.Timezone)

	if len(configTimezone) > 0 {
		currentTimezone = configTimezone
	}

	loc, err := time.LoadLocation(currentTimezone)
	if err == nil {
		now = now.In(loc)
	}

	return now.Format(TimestampFormat)
}

func (c *Client) sign(strToSign string) (*string, error) {
	pk, err := c.parsePrivateKey(c.Config.PrivateKey)
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	if _, err = h.Write([]byte(strToSign)); err != nil {
		return nil, err
	}

	signed, err := rsa.SignPKCS1v15(rand.Reader, pk, crypto.SHA256, h.Sum(nil))
	if err != nil {
		return nil, err
	}

	signature := base64.StdEncoding.EncodeToString(signed)

	return &signature, nil
}

func (c *Client) parsePrivateKey(pvKey string) (*rsa.PrivateKey, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(pvKey)
	if err != nil {
		return nil, err
	}

	pvtKey, err := x509.ParsePKCS1PrivateKey(keyBytes)
	if err == nil {
		return pvtKey, nil
	}

	key, err2 := x509.ParsePKCS8PrivateKey(keyBytes)
	if err2 == nil {
		valPvtKey, ok := key.(*rsa.PrivateKey)
		if ok {
			return valPvtKey, nil
		}

		return nil, fmt.Errorf("expected *rsa.PrivateKey, got %T", key)
	}

	return nil, errors.Join(err, err2)
}

func (c *Client) log(level string, data interface{}) {
	if c.Config.Log != nil && c.Config.Log.Enable {
		if c.Config.Log.Level == "error" && level != "error" {
			return
		}

		msg := map[string]interface{}{
			"level": level,
			"data":  data,
		}

		_ = goutil.WriteJsonToFile(msg, c.Config.Log.Path, c.Config.Log.Filename, c.Config.Log.Extension, c.Config.Log.Rotation)
	}
}
