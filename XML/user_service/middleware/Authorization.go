package middleware

import (
	"context"
	"log"
	"net/http"
	"os"

	my_auth "user/module/auth"
	"user/module/service"

	"github.com/euroteltr/rbac"
)

type AuthorizationHandler struct {
	Rbac             rbac.RBAC
	AdminPermissions *rbac.Permission
	Actions          []rbac.Action
	UserService      *service.UserService
}

func NewAuthorizationHandler(Rbac rbac.RBAC, AdminPermissions *rbac.Permission, Actions []rbac.Action, UserService *service.UserService) *AuthorizationHandler {
	return &AuthorizationHandler{Rbac, AdminPermissions, Actions, UserService}
}

type KeyUser struct{}

func ValidateToken(next http.Handler) http.Handler {

	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		tokenString := my_auth.ExtractToken(r)
		err, username := my_auth.TokenIsValid(tokenString)

		if err != nil {
			http.Error(rw, "Error middleware validating token: "+err.Error()+"\n", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), KeyUser{}, username)
		r = r.WithContext(ctx)
		rw.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(rw, r)
	})
}

func (handler *AuthorizationHandler) PermissionGranted(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		username, _ := r.Context().Value(KeyUser{}).(string)
		l := log.New(os.Stdout, "products-api ", log.LstdFlags)

		l.Printf("USERNAME FORM CONTEXT:" + username)
		err := handler.UserService.UserExists(username)
		if err != nil {
			http.Error(rw, err.Error()+"\n", http.StatusUnauthorized)
			return
		}

		userRole, err := handler.UserService.GetUserRole(username)

		if err != nil {
			http.Error(rw, err.Error()+"\n", http.StatusUnauthorized)
			return
		}

		handler.Rbac.IsGranted(userRole, handler.AdminPermissions, handler.Actions...)
		if !handler.Rbac.IsGranted(userRole, handler.AdminPermissions, handler.Actions...) {
			http.Error(rw, "User does not have permission", http.StatusUnauthorized)
			rw.WriteHeader(http.StatusForbidden)
			return
		}

		rw.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
