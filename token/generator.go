package token

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// Generator is a generator of Apple Music JWT token.
type Generator struct {
	// A 10-character key identifier (kid) key, obtained from your developer account.
	KeyId string

	// The issuer (iss) registered claim key, whose value is your 10-character Team ID,
	// obtained from your developer account.
	TeamId string

	// TTL (time-to-live), must not be greater than 15777000 (6 months in seconds).
	TTL int64

	// MusicKit private key.
	Secret []byte
}

// Generate generates a JWT token.
func (g Generator) Generate() (string, error) {
	now := time.Now()

	t := jwt.Token{
		Method: jwt.SigningMethodES256,
		Header: map[string]interface{}{
			"alg": jwt.SigningMethodES256.Alg(),
			"kid": g.KeyId,
		},
		Claims: jwt.MapClaims{
			"iss": g.TeamId,
			"iat": now.Unix(),
			"exp": now.Add(time.Second * time.Duration(g.TTL)).Unix(),
		},
		Signature: string(g.Secret),
	}

	key, err := ParsePKCS8PrivateKeyFromPEM(g.Secret)
	if err != nil {
		return "", err
	}

	return t.SignedString(key)
}

// ParsePKCS8PrivateKeyFromPEM parses PEM encoded PKCS8 Private Key Structure.
func ParsePKCS8PrivateKeyFromPEM(key []byte) (*ecdsa.PrivateKey, error) {
	var err error

	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, errors.New("Invalid Key: Key must be PEM encoded PKCS8 private key")
	}

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	var pkey *ecdsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*ecdsa.PrivateKey); !ok {
		return nil, errors.New("Key is not a valid PKCS8 private key")
	}

	return pkey, nil
}
