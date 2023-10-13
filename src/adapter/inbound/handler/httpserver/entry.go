package cli

import (
	"context"
	"encoding/json"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/infra/database/postgres"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/adapter/inbound/controller"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/src/adapter/outbound/repository"
	service "fiappos/ViniAlvesMartins/tech-challenge-fiap/src/core/service/product"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"gorm.io/gorm"

	"github.com/gorilla/mux"
)

// "fmt"

// "github.com/ViniAlvesMartins/tech-challenge-fiap/internal/adapter/handler"
type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

type Handler struct {
	Conn *gorm.DB
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	peter := User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
	}

	json.NewEncoder(w).Encode(peter)
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Logging logs all requests with its path and the time it took to process
func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Method(m string) Middleware {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Execute() {
	cfg, err := postgres.NewConfig()

	if err != nil {
		fmt.Println(err)
	}

	db, err := postgres.NewConnection(context.Background(), slog.New(slog.NewTextHandler(os.Stderr, nil)), cfg)

	if err != nil {
		fmt.Println(err)
	}

	h := Handler{
		Conn: db,
	}

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	router := mux.NewRouter()

	router.HandleFunc("/", Chain(h.getUser, Method("GET"), Logging()))
	router.HandleFunc("/product", Chain(productController.CreateProduct, Method("POST"), Logging()))

	http.ListenAndServe(":8080", router)

}
