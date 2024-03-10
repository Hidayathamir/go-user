// Package auth contains func related to auth.
package auth

import (
	"fmt"
	"strconv"
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
func GenerateUserJWTToken(userID int64) string {
	expireIn := time.Hour * time.Duration(config.JWT.ExpireHour)
	token := jwt.NewWithClaims(signingMethod, jwt.MapClaims{
		keyUserID: strconv.FormatInt(userID, 10),
		"exp":     time.Now().Add(expireIn).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.JWT.SignedKey))
	if err != nil {
		logrus.Warnf("jwt.Token.SigndString: %v", err)
	}

	tokenString = "Bearer " + tokenString

	return tokenString
}

// ValidateUserJWTToken -.
func ValidateUserJWTToken(tokenString string) (jwt.MapClaims, error) {
	keyFunc := func(*jwt.Token) (interface{}, error) {
		return []byte(config.JWT.SignedKey), nil
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
func GetUserIDFromJWTClaims(claims jwt.MapClaims) (int, error) {
	userIDAny, ok := claims[keyUserID]
	if !ok {
		return 0, errors.New("jwt.MapClaims[keyUserID]")
	}

	userID, ok := userIDAny.(int)
	if !ok {
		return 0, errors.New("userIDAny.(int)")
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
