package router

func UserRoutes(router *gin.Engine, userUsecase domain.UserUseCase) {

	userHandler := handler.UserHandler{UserUsecase: userUsecase}

	users := router.Group("/users")

	{
		// implement user endpoints here
		// users.POST("/", userHandler.CreateUser)
	}
}
