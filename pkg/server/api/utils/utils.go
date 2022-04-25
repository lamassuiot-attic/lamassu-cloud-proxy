package utils

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
)

const (
	PublicKeyHeader = "-----BEGIN PUBLIC KEY-----"
	PublicKeyFooter = "-----END PUBLIC KEY-----"
)

func CreateCAPool(CAPath string) (*x509.CertPool, error) {
	caCert, err := ioutil.ReadFile(CAPath)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	return caCertPool, nil
}

func ParseKeycloakPublicKey(data []byte) (*rsa.PublicKey, error) {
	pubPem, _ := pem.Decode(data)
	parsedKey, err := x509.ParsePKIXPublicKey(pubPem.Bytes)
	if err != nil {
		return nil, errors.New("Unable to parse public key")
	}
	pubKey := parsedKey.(*rsa.PublicKey)
	return pubKey, nil
}

func InsertNth(s string, n int) string {
	if len(s)%2 != 0 {
		s = "0" + s
	}
	var buffer bytes.Buffer
	var n_1 = n - 1
	var l_1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
			buffer.WriteRune('-')
		}
	}
	return buffer.String()
}

func ToHexInt(n *big.Int) string {
	return fmt.Sprintf("%x", n) // or %X or upper case
}
func DecodeB64(message string) (string, error) {
	base64Text := make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	_, err := base64.StdEncoding.Decode(base64Text, []byte(message))
	return string(base64Text), err
}
