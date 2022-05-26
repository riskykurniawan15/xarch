package jwt

import (
	"fmt"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/riskykurniawan15/xarch/config"
)

type JwtToken struct {
	cfg config.Config
}

func NewJwtToken(cfg config.Config) *JwtToken {
	return &JwtToken{
		cfg,
	}
}

type JwtCustomClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

func (JT JwtToken) GenerateToken(Structure *JwtCustomClaims) (string, error) {
	var mySigningKey = []byte(JT.cfg.JWT.SecretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = Structure.ID
	claims["email"] = Structure.Email
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(JT.cfg.JWT.Expired)).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (JT JwtToken) ClaimsToken(token string) (*JwtCustomClaims, error) {
	var mySigningKey = []byte(JT.cfg.JWT.SecretKey)

	TokenClaims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("your token has expired")
	}

	claims, ok := TokenClaims.Claims.(jwt.MapClaims)
	if ok && TokenClaims.Valid {
		id, err := strconv.Atoi(fmt.Sprint(claims["id"]))
		if err != nil {
			return nil, fmt.Errorf("malformed id")
		}

		data := &JwtCustomClaims{
			ID:    uint(id),
			Email: fmt.Sprint(claims["email"]),
		}
		return data, nil
	}

	return nil, fmt.Errorf("token not valid")
}
