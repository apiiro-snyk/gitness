// Copyright 2021 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package token

import (
	"fmt"
	"time"

	"github.com/harness/gitness/types"
	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
)

// Claims defines custom token claims.
type Claims struct {
	Admin bool `json:"admin"`

	jwt.StandardClaims
}

// Generate generates a token with no expiration.
func Generate(user *types.User, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		user.Admin,
		jwt.StandardClaims{
			Subject:  fmt.Sprint(user.ID),
			IssuedAt: time.Now().Unix(),
		},
	})

	res, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.Wrap(err, "Failed to sign token")
	}

	return res, nil
}

// GenerateExp generates a token with an expiration date.
func GenerateExp(user *types.User, exp int64, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		user.Admin,
		jwt.StandardClaims{
			ExpiresAt: exp,
			Subject:   fmt.Sprint(user.ID),
			IssuedAt:  time.Now().Unix(),
		},
	})

	res, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.Wrap(err, "Failed to sign token")
	}

	return res, nil
}