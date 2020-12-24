package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bayu-aditya/myfacilities-backend/lib/model"
	"github.com/bayu-aditya/myfacilities-backend/lib/variable"
	"github.com/dgrijalva/jwt-go"
)

// AuthorizationRequired using JWT claims
func AuthorizationRequired(ctx context.Context) error {
	c, err := GinContextFromContext(ctx)
	if err != nil {
		log.Fatal("AuthorizationRequired: gin context not include in this middleware")
	}

	// check if Authorization header request
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.Status(http.StatusUnauthorized)
		return errors.New("required authorization")
	}

	// check authorization header structure
	vals := strings.Split(authHeader, " ")
	if len(vals) < 2 || vals[0] != "Bearer" {
		c.Status(http.StatusUnauthorized)
		return errors.New("invalid credential")
	}
	token := vals[1]

	// get claims
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(variable.Project.JWT.Key), nil
	})
	if err != nil {
		c.Status(http.StatusUnauthorized)
		return errors.New("invalid credential")
	}

	c.Set("JWT_USER_ID", claims["user_id"])
	c.Set("JWT_EMAIL", claims["email"])
	c.Set("JWT_EXP", claims["exp"])

	return nil
}

// GetClaims from graphql context
func GetClaims(ctx context.Context) *variable.JwtClaims {
	c, err := GinContextFromContext(ctx)
	if err != nil {
		log.Fatal("GetClaims: gin context not include in this middleware")
	}

	return &variable.JwtClaims{
		UserID: c.GetString("JWT_USER_ID"),
		Email:  c.GetString("JWT_EMAIL"),
		Exp:    int64(c.GetFloat64("JWT_EXP")),
	}
}

// GenerateToken JWT
func GenerateToken(user *model.User) string {
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	claims["email"] = user.Email
	claims["user_id"] = user.ID

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(variable.Project.JWT.Key))
	if err != nil {
		log.Panic("middleware.GenerateToken Error")
	}
	return token
}
