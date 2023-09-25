package middleware

import (
	"golang_middleware/configs"
	"golang_middleware/helpers"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var SECRET = configs.InitAuthConfig().Secret

func GenerateJWTCookie(name string, c echo.Context) (string, error) {
	expTime := time.Now().Add(time.Hour * 24)
	claims := &configs.JWTClaim{
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Golang",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaim.SignedString(SECRET)
	if err != nil {
		return "", err
	}

	c.SetCookie(&http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	return token, nil
}

func Validate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwtCookie, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, "Unauthorized", nil))
			}
		}

		tokenString := jwtCookie.Value

		claims := &configs.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return SECRET, nil
		})

		if err != nil {
			switch err {
				case jwt.ErrSignatureInvalid:
					return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, "Unauthorized", nil))
				
				case jwt.ErrTokenExpired:
					return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, "Unauthorized, token expired", nil))
				default :
					return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, "Unauthorized", nil))
			}
		}

		if !token.Valid {
			return c.JSON(http.StatusUnauthorized, helpers.WebResponse(http.StatusUnauthorized, "Unauthorized", nil))
		}

		return next(c)
	}
}