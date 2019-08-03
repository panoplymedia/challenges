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
	"sales-portal-api/endpoints/customer"
	"sales-portal-api/endpoints/merchant"
	"sales-portal-api/endpoints/product"
	"sales-portal-api/endpoints/sale"
	"sales-portal-api/environment"
	"sales-portal-api/migrations"
	"time"
)

type jwtCustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func login(e *database.DbEngine) func(echo.Context) error {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		password := c.FormValue("password")

		// Verify User Exists
		usr, ok := e.VerifyUser(email, password)
		if !ok {
			return echo.ErrUnauthorized
		}

		// Set claims
		claims := &jwtCustomClaims{
			email,
			usr.Role,
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
				Issuer:    "Acme Cult Hero Supplies Inc.",
			},
		}

		// Create token using custom claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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

func AuthorizeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwtCustomClaims)

		if claims.Issuer != "Acme Cult Hero Supplies Inc." {
			c.Error(fmt.Errorf("unkown issuer %s", claims.Issuer))
		}

		return next(c)
	}
}

func main() {
	env := environment.NewEnvironment()
	logger := logrus.New()
	engine, err := database.NewEngine(env, logger)
	if err != nil {
		logger.Panicln(err)
	}

	// Run migrations to ensure the db has the latest changes
	migrations.Migrate(engine.DB.DB(), logger)

	customerSvc := customer.NewSvc(engine)
	customerHandler := customer.NewHandler(customerSvc)

	merchantSvc := merchant.NewSvc(engine)
	merchantHandler := merchant.NewHandler(merchantSvc)

	productSvc := product.NewSvc(engine)
	productHandler := product.NewHandler(productSvc)

	saleSvc := sale.NewSvc(engine)
	saleHandler := sale.NewHandler(saleSvc)

	e := echo.New()

	// Non-Auth Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{env.AllowOrigins},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// Login handler for authentication
	e.POST("/login", login(engine))

	// JWT Middleware
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(env.JwtSecret),
		Claims:     &jwtCustomClaims{},
	})

	// sale endpoint handlers
	e.POST("/sale", saleHandler.CreateSale, jwtMiddleware, AuthorizeMiddleware)
	e.GET("/sale/:id", saleHandler.GetSale, jwtMiddleware, AuthorizeMiddleware)
	e.PUT("/sale/:id", saleHandler.UpdateSale, jwtMiddleware, AuthorizeMiddleware)
	e.GET("/sale", saleHandler.ListSales, jwtMiddleware, AuthorizeMiddleware)
	e.DELETE("/sale/:id", saleHandler.DeleteSale, jwtMiddleware, AuthorizeMiddleware)
	e.GET("/sale/summary", saleHandler.GetSalesSummary, jwtMiddleware, AuthorizeMiddleware)
	e.POST("/sale/upload", saleHandler.UploadSalesCsv, jwtMiddleware, AuthorizeMiddleware)

	// customer endpoint handlers
	e.POST("/customer", customerHandler.CreateCustomer, jwtMiddleware, AuthorizeMiddleware)
	e.GET("/customer/:id", customerHandler.GetCustomer, jwtMiddleware, AuthorizeMiddleware)
	e.GET("/customer", customerHandler.ListCustomers, jwtMiddleware, AuthorizeMiddleware)
	e.PUT("/customer/:id", customerHandler.UpdateCustomer, jwtMiddleware, AuthorizeMiddleware)
	e.DELETE("/customer/:id", customerHandler.DeleteCustomer, jwtMiddleware, AuthorizeMiddleware)

	// merchant endpoint handlers
	e.POST("/merchant", merchantHandler.CreateMerchant, jwtMiddleware, AuthorizeMiddleware)
	e.GET("/merchant/:id", merchantHandler.GetMerchant, jwtMiddleware, AuthorizeMiddleware)
	e.PUT("/merchant/:id", merchantHandler.UpdateMerchant, jwtMiddleware, AuthorizeMiddleware)
	e.DELETE("/merchant/:id", merchantHandler.DeleteMerchant, jwtMiddleware, AuthorizeMiddleware)

	// product endpoint handlers
	e.POST("/product", productHandler.CreateProduct, jwtMiddleware, AuthorizeMiddleware)
	e.GET("/product", productHandler.ListProducts, jwtMiddleware, AuthorizeMiddleware)
	e.GET("/product/:merchantId", productHandler.GetMerchantProducts, jwtMiddleware, AuthorizeMiddleware)
	e.PUT("/product/:id", productHandler.UpdateProduct, jwtMiddleware, AuthorizeMiddleware)
	e.DELETE("/product/:id", productHandler.DeleteProduct, jwtMiddleware, AuthorizeMiddleware)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", env.ApiPort)))
}
