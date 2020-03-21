package auth

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// JWTAuthNZ returns an authentication and authorization handler.
func JWTAuthNZ(audience string, accPublicKeyURLTemplate string, secretHMAC string) AuthFunc {
	return func(ctx context.Context) (context.Context, error) {
		tokenStr, err := BearerFromMD(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, &jwt.StandardClaims{})
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		var authenticator Authenticator

		if _, ok := token.Header["kid"]; ok {
			authenticator = ServiceAccountAuthN(audience, accPublicKeyURLTemplate)
		} else {
			authenticator = HMACAuthN(secretHMAC)
		}

		sub, err := authenticator.Authenticate(ctx)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}

		grpc_ctxtags.Extract(ctx).Set("auth.sub", sub)
		newCtx := context.WithValue(ctx, "token", token)
		return newCtx, nil
	}
}
