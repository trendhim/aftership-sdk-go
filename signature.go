package aftership

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/url"
	"sort"
	"strings"
)

type AuthenticationType int64

const (
	APIKey AuthenticationType = iota
	AES
)

const (
	HeaderAsSignatureHMAC = "as-signature-hmac-sha256"
)

func GetSignature(authenticationType AuthenticationType, secretKey []byte, asHeaders map[string]string, contentType, uri, method, date, body string) (string, string, error) {
	canonicalizedAmHeaders := GetCanonicalizedHeaders(asHeaders)
	canonicalizedResource, err := GetCanonicalizedResource(uri)
	if err != nil {
		return "", "", err
	}
	signString, err := GetSignString(method, body, contentType, date, canonicalizedAmHeaders, canonicalizedResource)
	if err != nil {
		return "", "", err
	}

	if authenticationType == AES {
		return HeaderAsSignatureHMAC, GetHMACSignature(signString, secretKey), nil
	}
	return "", "", errors.New("authenticationType incorrect")
}

func GetHMACSignature(signString string, secret []byte) string {
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(signString))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func GetSignString(method, body, contentType, date, canonicalizedAmHeaders, canonicalizedResource string) (signature string, err error) {
	result := ""

	result += method + "\n"
	if body != "" {
		body, err = Md5Encode(body)
		if err != nil {
			return "", err
		}
	} else {
		contentType = ""
	}
	result += body + "\n"
	result += contentType + "\n"
	result += date + "\n"
	result += canonicalizedAmHeaders + "\n"
	result += canonicalizedResource
	return result, nil
}

func GetCanonicalizedHeaders(headers map[string]string) string {
	if headers == nil {
		return ""
	}
	keys := make([]string, 0, len(headers))
	newHeaders := make(map[string]string, len(headers))
	for key, value := range headers {
		newKey := strings.ToLower(key)
		if !strings.HasPrefix(newKey, "as-") {
			continue
		}
		keys = append(keys, newKey)
		newHeaders[newKey] = strings.TrimLeft(value, " ")
	}

	sort.Strings(keys)

	result := make([]string, 0, len(keys))
	for _, key := range keys {
		result = append(result, key+":"+newHeaders[key])
	}

	return strings.Join(result, "\n")
}

func GetCanonicalizedResource(rawUrl string) (result string, err error) {
	url, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	result += url.Path
	params := url.Query().Encode()

	if params != "" {
		result += "?" + params
	}

	return result, nil
}

func Md5Encode(source string) (hashed string, err error) {
	h := md5.New()
	_, err = h.Write([]byte(source))
	if err != nil {
		return
	}
	sum := h.Sum(nil)
	hashed = strings.ToUpper(hex.EncodeToString(sum))
	return
}
