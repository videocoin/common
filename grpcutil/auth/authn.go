package auth

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	jwt "github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"github.com/google/uuid"
)

// AuthenticatorFunc ...
type AuthenticatorFunc func(ctx context.Context) (string, error)

// Authenticate ...
func (f AuthenticatorFunc) Authenticate(ctx context.Context) (string, error) {
	return f(ctx)
}

// Authenticator authenticates requests.
type Authenticator interface {
	Authenticate(ctx context.Context) (string, error)
}

// ServiceAccountAuthN handles authentication for service accounts.
func ServiceAccountAuthN(ctx context.Context, audience string, pubKeyURLTemplate string) AuthenticatorFunc {
	return func(ctx context.Context) (string, error) {
		tokenStr, err := BearerFromMD(ctx)
		if err != nil {
			return "", err
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			claims := token.Claims.(*jwt.StandardClaims)

			url, err := url.Parse(claims.Audience)
			if err != nil {
				return nil, err
			}

			if url.Hostname() != audience {
				return nil, fmt.Errorf("Unexpected audience: %s", claims.Audience)
			}

			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			kid := token.Header["kid"].(string)
			if kid == "" {
				return nil, fmt.Errorf("Key ID is missing")
			}

			if _, err := uuid.Parse(token.Header["kid"].(string)); err != nil {
				return nil, fmt.Errorf("Invalid kid: %v", kid)
			}

			return ServiceAccountPublicKey(pubKeyURLTemplate, claims.Subject, kid)
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					return "", errors.New("")
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					return "", errors.New("Token is either expired or not active yet")
				}
			}
			return "", fmt.Errorf("Couldn't handle this token: %v", err)
		}

		return token.Claims.(*jwt.StandardClaims).Subject, nil
	}
}

// JWTHMACAuthN handles authentication based on JWT with HMAC protection.
func JWTHMACAuthN(ctx context.Context, secret string) AuthenticatorFunc {
	return func(ctx context.Context) (string, error) {
		tokenStr, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return "", err
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			claims := token.Claims.(*jwt.StandardClaims)
			_, err := uuid.Parse(claims.Subject)
			if err != nil {
				return nil, err
			}

			return secret, nil
		})
		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok {
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					return "", errors.New("")
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					return "", errors.New("Token is either expired or not active yet")
				}
			}
			return "", fmt.Errorf("Couldn't handle this token: %v", err)
		}

		return token.Claims.(*jwt.StandardClaims).Subject, nil
	}
}

/*
// AuthNZ handles authentication and authrorization.
func AuthNZ() grpc_auth.AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		tokenStr, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}

		token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, &jwt.StandardClaims{})
		if err != nil {
			return nil, err
		}


		//
		var
		_, ok := token.Header["kid"]
		if ok {
			authenticator = a.SvcAccAuth
		} else {
			authenticator = a.UserAuth
		}

	}
}
*/
