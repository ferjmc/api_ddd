package http

// MapUserRoutes
func (h *userHandlers) MapUserRoutes() {
	h.group.POST("/register", h.Register())
	h.group.POST("/login", h.Login())
	h.group.GET("/:id", h.GetUserByID())
	h.group.PUT("/:id", h.Update())
	h.group.GET("/me", h.GetMe())
}
