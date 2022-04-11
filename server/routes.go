package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/tejashwikalptaru/tutorial/handler"
	"github.com/tejashwikalptaru/tutorial/middlewares"
	"net/http"
)

type Server struct {
	chi.Router
}

func SetupRoutes() *Server {
	router := chi.NewRouter()

	router.Route("/api", func(api chi.Router) {
		//api.Get("/welcome", handler.Greet)
		api.Post("/register", handler.Register)
		api.Post("/login", handler.Login)
		api.Get("/logout", handler.Logout)

		api.Route("/", func(api chi.Router) {
			api.Use(middlewares.Middleware)
			api.Post("/addTask", handler.AddTask)
			api.Get("/getTask", handler.GetTask)
			api.Put("/updateTask", handler.UpdateTask)
		})

	})

	return &Server{router}
}

func (svc *Server) Run(port string) error {
	return http.ListenAndServe(port, svc)
}
