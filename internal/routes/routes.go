package routes

import (
	"net/http"

	services "github.com/Nithya-kalyani/go-boilerplate/internal/services/user-service"
	"github.com/Nithya-kalyani/go-boilerplate/pkg/utils"
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", defaultHandler)

	// Public route
	router.HandleFunc("/login", loginHandler).Methods("POST")

	// Private routes (e.g., user routes)
	privateRoute := router.PathPrefix("/").Subrouter()
	// privateRoute.Use(middleware.AuthMiddleware) // Apply middleware if needed
	privateRoute.HandleFunc("/user/list", services.ListUsers).Methods("GET")
	privateRoute.HandleFunc("/user/create", services.CreateUser).Methods("POST")

	return router
}

// loginHandler handles login without authentication
func loginHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Logged in successfully"}
	utils.RenderJSON(w, http.StatusOK, response)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the Home Page"))
}
