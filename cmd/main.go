package main

import (
	"golang-project-template/internal/common"
	shopAdapters "golang-project-template/internal/shop/adapters"
	service "golang-project-template/internal/shop/app"
	"golang-project-template/internal/shop/domain"
	shophandler "golang-project-template/internal/shop/ports/rest/handler"
	userAdapters "golang-project-template/internal/users/adapters"
	userApp "golang-project-template/internal/users/app"
	userController "golang-project-template/internal/users/ports/http/controller"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// app.Execute()
	httpServer()

}

func httpServer() *chi.Mux {
	db, err := common.ConnectToDb(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	//Testing router
	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, Chi!"))
	})

	// User router
	//userRepo := userAdapters.NewUserRepository(db)
	userRepo := userAdapters.NewUserRepository(db)
	userUsecase := userApp.NewUserUsecase(userRepo)
	userHandler := userController.NewUserController(userUsecase)

	// Shop router
	shopRepo := shopAdapters.NewShopRepository(db)
	shopFactory := domain.NewShopFactory(256)
	shopService := service.NewShopService(shopRepo, shopFactory, userRepo)
	shopHandler := shophandler.ShopHandler{ShopService: shopService}

	// Routers
	router.Route("/api", func(r chi.Router) {

		r.Route("/user", func(r chi.Router) {
			r.Post("/register-admin/", userHandler.RegisterAdminUserHandler)
			r.Post("/register-merchant/", userHandler.RegisterMerchantHandler)
			r.Post("/register-customer/", userHandler.RegisterCustomerHandler)
			r.Post("/login/", userHandler.LoginUserHandler)
		})

		r.Route("/shop", func(r chi.Router) {

			r.Post("/", shopHandler.CreateShop)
		})

	})

	server := &http.Server{Addr: os.Getenv("HTTP_PORT"), Handler: router}
	log.Println("Starting server on port...", os.Getenv("HTTP_PORT"))
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
	defer server.Close()

	return router
}
