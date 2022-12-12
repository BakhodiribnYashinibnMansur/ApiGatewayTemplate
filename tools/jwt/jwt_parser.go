package jwt

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	UserID  int64
	Expires int64
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(ctx *gin.Context) (*TokenMetadata, error) {
	token, err := verifyToken(ctx)
	if err != nil {
		return nil, err
	}

	// Setting and checking token and credentials.
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// User ID.
		userID := claims["id"].(string)

		userIDInt64, err := strconv.ParseInt(userID, 10, 64)
		if err != nil {
			return nil, err
		}
		// Expires time.
		expires := int64(claims["expires"].(float64))

		// User credentials.

		return &TokenMetadata{
			UserID:  userIDInt64,
			Expires: expires,
		}, nil
	}

	return nil, err
}

func extractToken(ctx *gin.Context) string {
	bearToken := ctx.GetHeader("Authorization")
	token := fmt.Sprintf("%v", bearToken)
	onlyToken := strings.Split(token, " ")
	if len(onlyToken) != 2 {
		return ""
	}
	if onlyToken[0] != "Bearer" {
		return ""
	}
	return onlyToken[1]
	// return token
}

func verifyToken(ctx *gin.Context) (*jwt.Token, error) {
	tokenString := extractToken(ctx)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(conf.JWTSecretKey), nil
}

// ParseRefreshToken func for parse second argument from refresh token.
func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
