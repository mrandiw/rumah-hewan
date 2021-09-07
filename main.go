package main

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Kucing struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World Home!")
}

func login(c echo.Context) error {

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

func getKucingFunc(c echo.Context) error {

	kucing := c.QueryParam("kucing")
	status := c.QueryParam("status")

	dataType := c.Param("type")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Ini kucingnya %s \nDan ini statusnya %s", kucing, status))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"kucing": kucing,
			"status": status,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "tipe harus string atau json",
	})
}

func addKucingFunc(c echo.Context) error {
	kucing := Kucing{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&kucing)
	if err != nil {
		log.Printf("Gagal melakukan decode %s", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status": "Gagal melakukan decode",
		})
	}

	// == save to database here ==
	log.Printf("Berhasil Menyimpan kucing dari request %v", kucing)
	return c.JSON(http.StatusOK, map[string]string{
		"status": "Success",
	})
}

func getDashboard(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "Success",
		"page":   "dashboard",
	})
}

func getDashboardCookie(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "Success",
		"page":   "cookie",
	})
}

func getDashboardJwt(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status": "Success",
		"page":   "JWT",
	})
}

// === MIDDLEWARE ===
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Rumah Hewan 1.1")
		c.Response().Header().Set("Developer", "Andi Wibowo")
		return next(c)
	}
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("SessionName")
		if err != nil {
			return c.String(http.StatusUnauthorized, "Authentication Failed")
		}

		if cookie.Value == "SessionValue" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "Authentication Failed")
	}
}

func main() {
	fmt.Println("Hello World.")

	e := echo.New()

	// Server header
	e.Use(ServerHeader)

	g := e.Group("/api/v1")
	gCookie := e.Group("/cookie/v1")
	gJwt := e.Group("/jwt/v1")

	gCookie.Use(checkCookie)
	gJwt.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("MySecret"),
		TokenLookup:   "cookie:JwtCookie",
	}))

	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${method}], host=${host}${path}, status=${status} latency=${latency}\n",
	}))

	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {

		// query ke database

		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("andi")) == 1 && subtle.ConstantTimeCompare([]byte(password), []byte("123456")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/dashboard", getDashboard)
	gCookie.GET("/main", getDashboardCookie)
	gJwt.GET("/main", getDashboardJwt)

	e.GET("/", home)
	e.GET("/login", login)

	e.GET("/getKucing/:type", getKucingFunc)

	e.POST("/addKucing", addKucingFunc)

	e.Start(":8080")
}
