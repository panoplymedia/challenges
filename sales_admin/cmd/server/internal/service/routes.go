package service

func (s *Service) setupRoutes() {
	// s.echo.GET("/login", s.login)
	s.echo.GET("/sales", s.listSales)
	s.echo.POST("/sales/upload", s.uploadSales)
}
