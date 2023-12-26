package app

import (
	"fmt"
	"golang-project-template/internal/common/postgres"
	"os"

	"golang-project-template/internal/common"
	"golang-project-template/internal/pkg/jwt"
	shopAdapters "golang-project-template/internal/shop/adapters"
	service "golang-project-template/internal/shop/app"
	"golang-project-template/internal/shop/domain"
	shophandler "golang-project-template/internal/shop/ports/rest/handler"
	userAdapters "golang-project-template/internal/users/adapters"
	userApp "golang-project-template/internal/users/app"
	"golang-project-template/internal/users/ports/grpc/server"
	userController "golang-project-template/internal/users/ports/http/controller"
	"log"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var rootCmd = &cobra.Command{Use: "run-grpc"}

var userGrpcServerCmd = &cobra.Command{
	Use: "user-grpc-server",
	Run: func(cmd *cobra.Command, args []string) {
		server.RunGRPCServer()
	},
}

var httpCmd = &cobra.Command{
	Use: "http-server",
	Run: func(cmd *cobra.Command, args []string) {
		httpServer()
	},
}

func Execute() {
	rootCmd.AddCommand(userGrpcServerCmd)
	rootCmd.AddCommand(httpCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
		os.Exit(1)
	}
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

	dbNew, err := postgres.New(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DATABASE"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		"disable",
	)

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
	shopRepo := shopAdapters.NewShopRepository(dbNew)
	shopFactory := domain.NewShopFactory(256, 256)
	shopService := service.NewShopService(shopRepo, shopFactory, userRepo)
	shopHandler := shophandler.ShopHandler{ShopService: shopService}

	// Routers
	router.Route("/api", func(r chi.Router) {

		r.Route("/users", func(r chi.Router) {
			r.Post("/register-admin/", userHandler.RegisterAdminUserHandler)
			r.Post("/register-merchant/", userHandler.RegisterMerchantHandler)
			r.Post("/register-customer/", userHandler.RegisterCustomerHandler)
			r.Post("/login/", userHandler.LoginUserHandler)
			r.With(jwt.AuthMiddleWare).Get("/get-user/{phone_number}", userHandler.GetUserByPhoneNumberHandler)
		})

		r.Route("/shop", func(r chi.Router) {

			r.Post("/", shopHandler.CreateShop)
			r.Get("/{id}", shopHandler.GetShopById)
			r.Get("/", shopHandler.GetAllShops)
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
