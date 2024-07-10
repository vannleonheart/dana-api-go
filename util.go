package dana

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/vannleonheart/goutil"
	"strings"
)

func GenerateRequestId(minLength, maxLength int, charset string) string {
	r := goutil.NewRandomString(charset)

	return r.WithCharset(goutil.NumCharset).GenerateRange(minLength, maxLength)
}

func EncodeRequestBody(data interface{}) string {
	by, _ := json.Marshal(data)
	hash := sha256.New()
	hash.Write(by)
	str := hex.EncodeToString(hash.Sum(nil))

	return strings.ToLower(str)
}
