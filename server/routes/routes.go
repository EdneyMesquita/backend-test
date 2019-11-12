package routes

import (
	"backend-test/server/handlers/workflows"
	"net/http"

	cors "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//Run provides the routes to the application
func Run(port string) {
	//starting new router
	r := mux.NewRouter()

	allowedHeaders := cors.AllowedHeaders([]string{
		"X-Requested-With",
		"Content-Type",
		"Accept",
		"Access-Control-Allow-Origin",
		"Authorization",
		"Key",
		"Token",
	})
	allowedOrigins := cors.AllowedOrigins([]string{"*"})
	allowedMethods := cors.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"})

	r.HandleFunc("/workflow", workflows.List).Methods("GET")
	r.HandleFunc("/workflow/consume", workflows.Consume).Methods("GET")
	r.HandleFunc("/workflow", workflows.Create).Methods("POST", "OPTIONS")
	r.HandleFunc("/workflow/{uuid}", workflows.Change).Methods("PATCH")

	// r.Use(middlewares.Logging)
	http.ListenAndServe(":"+port, cors.CORS(allowedHeaders, allowedOrigins, allowedMethods)(r))
}
