package app

import (
	"log"
	"net/http"
	"os"

	"github.com/aidityasadhakim/go-unit-converter/app/controllers"
	"github.com/aidityasadhakim/go-unit-converter/app/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	port string
	router *chi.Mux
	service *services.UnitConverterService
	controller *controllers.UnitConverterController
}

func NewServer() *Server {
	var server Server
	var set bool
	server.port, set = os.LookupEnv("PORT")
	if !set {
		server.port = "8080"
	}

	server.router = chi.NewRouter()
	server.service = services.NewUnitConverterService()
	server.controller = controllers.NewUnitConverterController(server.service)

	server.router.Use(middleware.Logger)
	server.router.Use(middleware.Recoverer)
	server.router.Route("/", func(r chi.Router) {
		r.Get("/", server.controller.HomeHandler)
		r.Get("/length", server.controller.LengthHandler)
		r.Get("/temperature", server.controller.TemperatureHandler)
		r.Get("/weight", server.controller.WeightHandler)
		r.Post("/length", server.controller.LengthResultHandler)
		r.Post("/temperature", server.controller.TemperatureResultHandler)
		r.Post("/weight", server.controller.WeightResultHandler)
	})

	return &server
}

func (s *Server) Start() {
	log.Printf("Starting server on port %s\n", s.port)
	log.Fatal(http.ListenAndServe(":" + s.port, s.router))	
}