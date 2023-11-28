package router

func InitRoutes(db *sqlx.DB) *gin.Engine {
	router := gin.New()

	userRepo := adapters.NewUserRepository(db)
	userUsecase := app.NewUserUseCase(userRepo)
	UserRoutes(router, userUsecase)

	// Implements other routes(shop and product) here like user

	return router
}
