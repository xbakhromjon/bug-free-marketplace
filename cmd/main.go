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

	router := createRouter(db)

	server := &http.Server{Addr: os.Getenv("RPC_PORT"), Handler: router}
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic(err)
	}
	defer server.Close()
	log.Println("Starting server...")

}

func createRouter(db *pgx.Conn) *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	//Testing router
	router.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("Hello, Chi!"))
	})

	// User router
	// ...

	// Shop router
	shopRepo := adapters.NewShopRepository(db)
	shopFactory := domain.NewShopFactory(256)
	shopService := service.NewShopService(shopRepo, shopFactory)
	shopHandler := shophandler.ShopHandler{ShopService: shopService}
	router.Route("/shop", func(r chi.Router) {
		r.Post("/", shopHandler.CreateShop)
	})

	return router
}
