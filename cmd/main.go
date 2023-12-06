package main

import (
	"golang-project-template/internal/common"
	"golang-project-template/internal/shop/adapters"
	service "golang-project-template/internal/shop/app"
	"golang-project-template/internal/shop/domain"
	shophandler "golang-project-template/internal/shop/ports/rest/handler"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jackc/pgx"
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

	server := &http.Server{Addr: os.Getenv("HTTP_PORT"), Handler: router}
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
	defer server.Close()
	log.Println("Starting server...")

	//Testing router
	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, Chi!"))
	})

	// User router
	// ...

	// Shop router
	shopRepo := adapters.NewShopRepository(db)
	userRepo := mockUserRepo{db: db}
	shopFactory := domain.NewShopFactory(256)
	shopService := service.NewShopService(shopRepo, shopFactory, userRepo)
	shopHandler := shophandler.ShopHandler{ShopService: shopService}

	// Routers
	router.Route("/api", func(r chi.Router) {

		r.Route("/shop", func(r chi.Router) {
			r.Post("/", shopHandler.CreateShop)
		})

	})

	return router
}

type mockUserRepo struct {
	db *pgx.Conn
}

func (u mockUserRepo) UserExists(id int) (bool, error) {
	if id == 99 {
		return false, domain.ErrUserNotExists
	}
	return true, nil
}
