package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func Login(c echo.Context) error {

	username := c.QueryParam("username")
	password := c.QueryParam("password")

	if username == "andi" && password == "123456" {
		cookie := new(http.Cookie)
		cookie.Name = "SessionName"
		cookie.Value = "SessionValue"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)

		// create jwt token
		token, err := createJwtToken()
		if err != nil {
			log.Println("Failed Create JWT Token", err)
			return c.String(http.StatusInternalServerError, "Failed Create JWT Token.")
		}

		JwtCookie := new(http.Cookie)
		JwtCookie.Name = "JwtCookie"
		JwtCookie.Value = token
		JwtCookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(JwtCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"token":  token,
			"status": "Success",
		})
	}

	return c.String(http.StatusOK, "Login Failed.")
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"andi",
		jwt.StandardClaims{
			Id:        "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("MySecret"))
	if err != nil {
		return "", err
	}

	return token, nil
}
