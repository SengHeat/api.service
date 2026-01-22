package routes

import (
	"database/sql"
	"net/http"

	"api.drsb-purchase-service/config"
	"api.drsb-purchase-service/internal/api/handlers/user"
	"api.drsb-purchase-service/internal/infrastructure/cache"
	"api.drsb-purchase-service/internal/repository"
	"api.drsb-purchase-service/internal/service"

	//"api.drsb-purchase-service/internal/repository"
	//"api.drsb-purchase-service/internal/service"

	"github.com/gorilla/mux"
)

func Setup(db *sql.DB, redis *cache.Redis /*,log *log.Logger*/, cfg *config.Config) http.Handler {
	router := mux.NewRouter()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepo, cfg)

	// Initialize handlers
	userHandler := user.NewHandler(userService)

	// Initialize middleware
	//authMiddleware := middleware.NewAuthMiddleware(userService)
	//rateLimiter := middleware.NewRateLimiter(100, time.Minute)

	// Global middleware
	//router.Use(middleware.CORS)
	//router.Use(middleware.Logger(log))
	//router.Use(rateLimiter.Limit)

	// API v1 routes
	api := router.PathPrefix("/api/v1").Subrouter()

	// Public routes
	//api.HandleFunc("/health", healthCheck).Methods("GET")
	//api.HandleFunc("/auth/login", userHandler.Login).Methods("POST")
	api.HandleFunc("/users", userHandler.Create).Methods("POST")
	//
	//// Protected routes (require authentication)
	//protected := api.PathPrefix("").Subrouter()
	//protected.Use(authMiddleware.Authenticate)
	//
	//protected.HandleFunc("/auth/me", userHandler.Me).Methods("GET")
	//protected.HandleFunc("/users", userHandler.List).Methods("GET")
	api.HandleFunc("/users/{id}", userHandler.Get).Methods("GET")
	//protected.HandleFunc("/users/{id}", userHandler.Update).Methods("PUT")
	//
	//// Admin only routes
	//admin := protected.PathPrefix("").Subrouter()
	//admin.Use(authMiddleware.RequireAdmin)
	//
	//admin.HandleFunc("/users/{id}", userHandler.Delete).Methods("DELETE")

	return router
}
