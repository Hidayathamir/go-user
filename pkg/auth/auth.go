// Package auth contains func related to auth.
package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/Hidayathamir/go-user/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

var (
	signingMethod = jwt.SigningMethodHS256
)

const (
	keyUserID = "user_id"
)

// GenerateUserJWTToken return jwt string.
func GenerateUserJWTToken(userID int64, cfg config.Config) string {
	expireIn := time.Hour * time.Duration(cfg.JWT.ExpireHour)
	token := jwt.NewWithClaims(signingMethod, jwt.MapClaims{
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

// ValidateUserJWTToken -.
func ValidateUserJWTToken(cfg config.Config, tokenString string) (jwt.MapClaims, error) {
	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	keyFunc := func(*jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.SignedKey), nil
	}

	token, err := jwt.Parse(tokenString, keyFunc)
	if err != nil {
		err = fmt.Errorf("jwt.Parse: %w", err)
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return jwt.MapClaims{}, fmt.Errorf("auth is not a token: %w", err)
		}
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return jwt.MapClaims{}, fmt.Errorf("invalid signature: %w", err)
		}
		if errors.Is(err, jwt.ErrTokenExpired) {
			return jwt.MapClaims{}, fmt.Errorf("token expired: %w", err)
		}
		if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return jwt.MapClaims{}, fmt.Errorf("token not valid yet: %w", err)
		}
		return jwt.MapClaims{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return jwt.MapClaims{}, errors.New("jwt.Token.Claims.(jwt.MapClaims)")
	}

	return claims, nil
}

// GetUserIDFromJWTClaims return userID from jwt.MapClaims.
func GetUserIDFromJWTClaims(claims jwt.MapClaims) (int64, error) {
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
