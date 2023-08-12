package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	
	"github.com/sjxiang/biz-demo/easy-note/pkg/consts"
)

type Auth2Claims struct {
	User   int64       `json:"user"`
	
	jwt.RegisteredClaims
}



// 生成
func GenerateAuth2Token(userID int64, redirectURL string) (string, error) {
	claims := &Auth2Claims{
		User:     userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "sjxiang",
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Minute * 15),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(consts.SecretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}



// 提取
func ExtractAuth2Token(stateToken string) (userID int64, err error) {
	authClaims := &Auth2Claims{}
	token, err := jwt.ParseWithClaims(stateToken, authClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(consts.SecretKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*Auth2Claims)
	if !(ok && token.Valid) {
		return 0, err
	}

	return claims.User, nil
}
