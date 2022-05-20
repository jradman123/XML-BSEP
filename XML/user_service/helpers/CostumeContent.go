package helpers

import "github.com/labstack/echo"

type CustomContext struct {
	echo.Context
}

func NewCustomContext() *CustomContext {
	return &CustomContext{}
}

// return u.echoContent.JSON(http.StatusOK, map[string]string{
// 	"accessToken": token,
// 	"roles":       string(strings.Join(claims.Roles, "")),
// 	"expireTime":  tokenCreationTime.String(),
// })
