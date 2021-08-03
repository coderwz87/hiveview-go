package utils

//用户密码认证
import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"hiveview/models"
	"time"
)

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte("hiveview"),
	}
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

func GenerateToken(user models.Users) (token string, err error) {
	j := NewJWT()
	claims := CustomClaims{
		user.Username,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),            // 生成时间
			ExpiresAt: time.Now().Unix() + 24*60*60, //过期时间1天
			Issuer:    "wangzhe",
		},
	}
	token, err = j.CreateToken(claims)
	return
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {

		return claims, nil
	}

	return nil, TokenInvalid
}
