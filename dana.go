package dana

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/vannleonheart/goutil"
	"strings"
	"time"
)

func New(config Config) *Client {
	return &Client{
		Config: config,
	}
}

func GenerateRequestId(length int) string {
	r := goutil.NewRandomString("")

	return r.WithCharset(goutil.NumCharset).Generate(length)
}

func (c *Client) SetAccessToken(accessToken *AccessToken) {
	c.accessToken = accessToken
}

func (c *Client) WithAccessToken(accessToken *AccessToken) *Client {
	c.SetAccessToken(accessToken)

	return c
}

func (c *Client) getTimestamp() string {
	now := time.Now()
	timezone := DefaultTimezone

	if c.Config.Timezone != nil && len(*c.Config.Timezone) > 0 {
		timezone = *c.Config.Timezone
	}

	loc, err := time.LoadLocation(timezone)
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

func (c *Client) encodeRequestData(data interface{}) string {
	by, _ := json.Marshal(data)
	hash := sha256.New()
	hash.Write(by)
	str := hex.EncodeToString(hash.Sum(nil))
	return strings.ToLower(str)
}

func (c *Client) getChannelId() string {
	channelId := DefaultChannelId

	return channelId
}
