// Package auth contains func related to auth.
package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/Hidayathamir/go-user/pkg/gouser"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const (
	keyUserID = "user_id"
)

// GenerateUserJWTToken return jwt string.
func GenerateUserJWTToken(userID int64, cfg config.Config) string {
	expireIn := time.Hour * time.Duration(cfg.JWT.ExpireHour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		keyUserID: userID,
		"exp":     time.Now().Add(expireIn).Unix(),
	})

	tokenString, err := token.SignedString([]byte(cfg.JWT.SignedKey))
	if err != nil {
		logrus.Warnf("jwt.Token.SigndString: %v", err)
	}

	tokenString = "Bearer " + tokenString

	return tokenString
}

// validateUserJWTToken parses and validates and verifies JWT token string.
func validateUserJWTToken(cfg config.Config, tokenString string) (jwt.MapClaims, error) {
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("token.Method.(*jwt.SigningMethodHMAC): %v", token.Header["alg"])
		}
		return []byte(cfg.JWT.SignedKey), nil
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		return jwt.MapClaims{}, fmt.Errorf("jwt.Parse: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return jwt.MapClaims{}, errors.New("jwt.Token.Claims.(jwt.MapClaims)")
	}

	return claims, nil
}

// getUserIDFromJWTClaims return userID from jwt.MapClaims.
func getUserIDFromJWTClaims(claims jwt.MapClaims) (int64, error) {
	userIDAny, ok := claims[keyUserID]
	if !ok {
		return 0, errors.New("jwt.MapClaims[keyUserID]")
	}

	if v, ok := userIDAny.(int); ok {
		return int64(v), nil
	} else if v, ok := userIDAny.(int32); ok {
		return int64(v), nil
	} else if v, ok := userIDAny.(int64); ok {
		return v, nil
	} else if v, ok := userIDAny.(float32); ok {
		return int64(v), nil
	} else if v, ok := userIDAny.(float64); ok {
		return int64(v), nil
	}

	return 0, errors.New("type assert user id in jwt map claims as int, int64, int32, float64, float32")
}

// GetUserIDFromJWTTokenString return userID from JWT token string.
func GetUserIDFromJWTTokenString(cfg config.Config, tokenString string) (int64, error) {
	claims, err := validateUserJWTToken(cfg, tokenString)
	if err != nil {
		err := fmt.Errorf("ValidateUserJWTToken: %w", err)
		return 0, fmt.Errorf("%w: %w", gouser.ErrJWTAuth, err)
	}

	userID, err := getUserIDFromJWTClaims(claims)
	if err != nil {
		return 0, fmt.Errorf("GetUserIDFromJWTClaims: %w", err)
	}

	return userID, nil
}

// GenerateHashPassword generate hashed password.
func GenerateHashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("bcrypt.GenerateFromPassword: %w", err)
	}
	return string(hashedPassword), nil
}

// CompareHashAndPassword compares hashed password with password, return error
// on failure.
func CompareHashAndPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return fmt.Errorf("bcrypt.CompareHashAndPassword: %w", err)
	}
	return nil
}
