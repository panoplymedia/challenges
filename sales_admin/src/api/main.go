package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"sales-portal-api/database"
	"sales-portal-api/endpoints/user"
	"sales-portal-api/environment"
	"time"
)

func login(e *database.DbEngine) func(echo.Context) error {
	return func(c echo.Context) error {
		username := c.FormValue("username")
		password := c.FormValue("password")

		// Verify User Exists
		if ok := e.VerifyUser(username, password); !ok {
			return echo.ErrUnauthorized
		}

		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	env := environment.NewEnvironment()
	logger := logrus.New()
	engine, err := database.NewEngine(env, logger)
	if err != nil {
		logger.Panicln(err)
	}

	userSvc := user.NewSvc(engine)
	userHandler := user.NewHandler(userSvc)

	e := echo.New()

	// Non-Auth Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{fmt.Sprintf("http://%s:%s", env.AppHost, env.AppPort)},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/login", login)

	// JWT Middleware
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(env.JwtSecret),
	}))

	e.POST("/users", userHandler.CreateUser)
	e.GET("/users/:id", userHandler.GetUser)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", env.ApiPort)))
}
