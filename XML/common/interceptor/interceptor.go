package interceptor

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"os"
	"strings"
)

type AuthInterceptor struct {
	accessibleRoles map[string][]string
	publicKey       string
}

//TODO:ovdje vratiti na *rsa.PublicKey umesto stringa
func NewAuthInterceptor(accessibleRoles map[string][]string, publicKey string) *AuthInterceptor {
	return &AuthInterceptor{
		accessibleRoles: accessibleRoles,
		publicKey:       publicKey,
	}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		fmt.Println(info.FullMethod)
		ctx, err := interceptor.Authorize(ctx, info.FullMethod)
		if err != nil {

			return nil, err
		}

		return handler(ctx, req)
	}
}

type LoggedInUserKey struct {
}

func (interceptor *AuthInterceptor) Authorize(ctx context.Context, method string) (context.Context, error) {

	accessibleRoles, ok := interceptor.accessibleRoles[method]
	if !ok {
		return ctx, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("NEMA METADATA")
		return ctx, status.Errorf(codes.Unauthenticated, "Unauthorized")

	}

	err, tokenString := parseToken(md)
	if err != nil {
		return ctx, err

	}

	err, claimsRoles := TokenIsValid(ctx, tokenString)
	if err != nil {
		fmt.Println("TOKEN NOT VALID")
		return ctx, status.Errorf(codes.Unauthenticated, "Unauthorized")
	}

	for _, claimsRole := range claimsRoles {
		for _, role := range accessibleRoles {
			if role == claimsRole {
				fmt.Println(role)
				return context.WithValue(ctx, LoggedInUserKey{}, role), nil
			}
		}
	}

	return ctx, status.Errorf(codes.PermissionDenied, "Forbidden")
}

func parseToken(md metadata.MD) (error, string) {
	var values []string
	values = md.Get("Authorization")
	if len(values) == 0 {
		fmt.Println("NEMA AUTH")
		return status.Errorf(codes.Unauthenticated, "Unauthorized"), ""
	}

	authHeader := values[0]
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		fmt.Println("NIJE SPLIT")
		return status.Errorf(codes.Unauthenticated, "Unauthorized"), ""
	}
	return nil, parts[1]
}

func TokenIsValid(ctx context.Context, tokenString string) (error, []string) {

	claims, err := VerifyToken(tokenString)

	if err != nil {
		return status.Errorf(codes.Unauthenticated, "Unauthorized"), nil
	}
	err = claims.Valid()
	if err != nil {
		fmt.Println("CLAIMS NOT VALID")
		return status.Errorf(codes.Unauthenticated, "Unauthorized"), nil
	}
	return nil, claims.Roles
}

func VerifyToken(tokenString string) (*JwtClaims, error) {
	claims := &JwtClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		fmt.Println("Error parsing claims")
		return nil, err
	}

	return claims, nil
}
