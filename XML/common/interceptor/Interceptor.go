package interceptor

import (
	"common/module/logger"
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"os"
	"strings"
)

type AuthInterceptor struct {
	accessibleRoles map[string][]string
	publicKey       *rsa.PublicKey
	logError        *logger.Logger
}

func NewAuthInterceptor(accessibleRoles map[string][]string, publicKey *rsa.PublicKey, logError *logger.Logger) *AuthInterceptor {
	return &AuthInterceptor{
		accessibleRoles: accessibleRoles,
		publicKey:       publicKey,
		logError:        logError,
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
		interceptor.logError.Logger.Errorf("ERR:UNOTHORIZED:NO METADATA")
		return ctx, status.Errorf(codes.Unauthenticated, "Unauthorized")

	}

	err, tokenString := parseToken(md, interceptor.logError)
	if err != nil {
		return ctx, err

	}

	err, claimsRoles := TokenIsValid(ctx, tokenString)
	if err != nil {
		interceptor.logError.Logger.Errorf("ERR:UNOTHORIZED:TOKEN INVALID")
		return ctx, status.Errorf(codes.Unauthenticated, "Unauthorized")
	}

	userName := getUsernameFromClaim(tokenString)

	for _, claimsRole := range claimsRoles {
		for _, role := range accessibleRoles {
			if role == claimsRole {
				fmt.Println(role)
				//k := loggedInUser("loggedIn")
				return context.WithValue(ctx, LoggedInUserKey{}, userName), nil
				//return context.WithValue(ctx, "loggedIn", getUsernameClaim(tokenString)), nil //zamjenila sam da umesto role ide username
			}
		}
	}

	interceptor.logError.Logger.WithFields(logrus.Fields{
		"user": userName,
	}).Errorf("ERR:FORBIDEN")
	return ctx, status.Errorf(codes.PermissionDenied, "Forbidden")
}

func parseToken(md metadata.MD, logError *logger.Logger) (error, string) {
	var values []string
	values = md.Get("Authorization")
	if len(values) == 0 {
		fmt.Println("NEMA AUTH")
		logError.Logger.Errorf("ERR:UNOTHORIZED:NO AUTH")
		return status.Errorf(codes.Unauthenticated, "Unauthorized"), ""
	}

	authHeader := values[0]
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 {
		logError.Logger.Errorf("ERR:UNOTHORIZED:NO JWT") // samo stoji Bearer bez tokena
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

func getUsernameFromClaim(tokenString string) string {
	myClaims := &JwtClaims{}
	_, err := jwt.ParseWithClaims(tokenString, myClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte("<YOUR VERIFICATION KEY>"), nil
	})
	fmt.Println(err)
	//fmt.Println(myToken)
	//for key, val := range myClaims {
	//	if key == "username" {
	//		return my
	//	}
	//}
	return myClaims.Username
}
