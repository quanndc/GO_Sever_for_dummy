package middlewares

import (
	"firebase.google.com/go/v4/auth"
	"github.com/labstack/echo"
	"itss.edu.vn/todo/services/utilities"
)

func NewAuthMiddleware() func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rawToken := c.Request().Header.Get("Authorization")
			if rawToken == "" {
				return c.JSON(401, "Unauthorized")
			}
			uid, err := utilities.JWTVerify(rawToken)
			if err != nil {
				return c.JSON(401, "Unauthorized")
			}
			c.Set("uid", uid)
			return next(c)
		}
	}
}

func NewFirebaseAuthMiddleware(authClient *auth.Client) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rawToken := c.Request().Header.Get("Authorization")
			if rawToken == "" {
				return c.JSON(401, "Unauthorized")
			}
			uid, err := utilities.FirebaseUserVerify(rawToken, authClient)
			if err != nil {
				return c.JSON(401, "Unauthorized")
			}
			c.Set("uid", uid)
			return next(c)
		}
	}
}
