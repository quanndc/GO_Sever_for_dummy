package utilities

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/cristalhq/jwt"
)

const (
	SECRET_KEY = "my-secret-key"
)

func JWTSign(uid string) (string, error) {

	signer, err := jwt.NewHS256([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	tokenBuilder := jwt.NewTokenBuilder(signer)
	token, err := tokenBuilder.Build(jwt.StandardClaims{
		Audience:  jwt.Audience{"services"},
		ExpiresAt: jwt.Timestamp(time.Now().Add(time.Minute * 30).UnixMilli()),
		ID:        uid,
		IssuedAt:  jwt.Timestamp(time.Now().UnixMilli()),
		Issuer:    "services",
		Subject:   "services",
	})

	if err != nil {
		return "", err
	}
	return token.String(), nil
}

func JWTVerify(stringToken string) (string, error) {
	signer, err := jwt.NewHS256([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	token, err := jwt.ParseAndVerify([]byte(stringToken), signer)
	if err != nil {
		return "", err
	}

	payLoadRaw := token.RawClaims()
	//decode payload

	// payLoad := &jwt.StandardClaims{}
	data := make(map[string]interface{})
	err = json.Unmarshal(payLoadRaw, &data)
	if err != nil {
		return "", err
	}

	uid, ok := data["jti"].(string)
	if !ok {
		return "", fmt.Errorf("invalid uid")
	}
	return uid, nil
}

func FirebaseUserVerify(token string, authClient *auth.Client) (string, error) {
	decodedToken, err := authClient.VerifyIDToken(context.Background(), token)
	if err != nil {
		return "", err
	}
	return decodedToken.UID, nil
}
