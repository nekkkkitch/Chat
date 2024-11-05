package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	PrivateKey             *rsa.PrivateKey `yaml:"private_key" env-prefix:"PRIVATEKEY" env-default:""`
	PublicKey              *rsa.PublicKey  `yaml:"public_key" env-prefix:"PUBLICKEY" env-default:""`
	AccessTokenExpiration  time.Duration   `yaml:"access_token_expiration" env-prefix:"ACCESSTOKENEXPIRATION" env-default:"3600"`
	RefreshTokenExpiration time.Duration   `yaml:"private_key" env-prefix:"PRIVATEKEY" env-default:"36000"`
}

type Config struct {
	Key                    string
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
}

const (
	InvalidToken = "invalid token"
	ExpiredToken = "expired token"
)

func New(cfg *Config) (JWT, error) {
	jwt := JWT{}
	var err error
	privateKeyString := cfg.Key
	jwt.AccessTokenExpiration = cfg.AccessTokenExpiration * time.Second
	jwt.RefreshTokenExpiration = cfg.RefreshTokenExpiration * time.Second
	if privateKeyString == "" {
		jwt.PrivateKey, err = rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return JWT{}, err
		}
		jwt.PublicKey = &jwt.PrivateKey.PublicKey
		return jwt, nil
	}
	keyBytes := convertStringToBytesSlice(privateKeyString)
	jwt.PrivateKey, err = x509.ParsePKCS1PrivateKey(keyBytes)
	jwt.PublicKey = &jwt.PrivateKey.PublicKey
	if err != nil {
		return JWT{}, err
	}
	return jwt, nil
}

func (j *JWT) CreateTokens(user_id int) (string, string, error) {
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.AccessTokenExpiration)),
		Subject:   strconv.Itoa(user_id),
	}).SignedString(j.PrivateKey)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.RefreshTokenExpiration)),
	}).SignedString(j.PrivateKey)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func (j *JWT) ValidateToken(token string) (bool, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.PublicKey, nil
	})
	if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return false, fmt.Errorf(ExpiredToken)
		case errors.Is(err, jwt.ErrSignatureInvalid) || errors.Is(err, jwt.ErrTokenUnverifiable):
			return false, fmt.Errorf(InvalidToken)
		default:
			return false, err
		}
	}
	return true, nil
}

func (j *JWT) GetIDFromToken(token string) (int, error) {
	var id int
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.PublicKey, nil
	})
	if err != nil {
		return 0, err
	}
	id, err = strconv.Atoi(claims["sub"].(string))
	if err != nil {
		panic(err)
	}
	return id, nil
}

func convertStringToBytesSlice(line string) []byte {
	line = strings.Trim(line, "[]")
	parts := strings.Split(line, " ")
	var bytes []byte
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		bytes = append(bytes, byte(num))
	}
	return bytes
}