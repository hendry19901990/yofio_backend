package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/hendry19901990/yofio_backend/controllers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	cont := controllers.Controller{
		DBType:        os.Getenv("DB_TYPE"),
		URLConnection: fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME")),
	}
	if err := cont.Init(); err != nil {
		panic(err)
	}

	r.Route("/api", func(r chi.Router) {
		r.Post("/credit-assignment", cont.CreditAssignment)
		r.Post("/statistics", cont.Statistics)
	})

	http.ListenAndServe(":9090", r)
}
