package handlers

import (
	"context"
	"fiappos/ViniAlvesMartins/tech-challenge-fiap/infra/database/postgres"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// "fmt"

// "github.com/ViniAlvesMartins/tech-challenge-fiap/internal/adapter/handler"

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

/*type Handler struct {
	Conn *gorm.DB
}*/

func Execute() {
	cfg, err := postgres.NewConfig()

	if err != nil {
		fmt.Println(err)
	}

	db, err := postgres.NewConnection(context.Background(), slog.New(slog.NewTextHandler(os.Stderr, nil)), cfg)

	if err != nil {
		fmt.Println(err)
	}

	h := New(db)

	router := mux.NewRouter()

	// router.HandleFunc("/", Chain(h.getUser, Method("GET"), Logging()))

	router.HandleFunc("/clients", h.GetAllClients).Methods(http.MethodGet)
	router.HandleFunc("/clients/{id}", h.GetClient).Methods(http.MethodGet)
	router.HandleFunc("/clients", h.AddClient).Methods(http.MethodPost)
	router.HandleFunc("/clients/{id}", h.UpdateClient).Methods(http.MethodPut)
	router.HandleFunc("/clients/{id}", h.DeleteClient).Methods(http.MethodDelete)

	http.ListenAndServe(":8080", router)

	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World! teste1s2122",
	// 	})
	// })

	// router.GET("/os", func(c *gin.Context) {
	// 	c.String(200, runtime.GOOS)
	// })

	// router.Run(":8080")
}
