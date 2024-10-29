package main

import (
	"log"
	"net/http" // Standard library for HTTP handling
    "go-crud-challenge/internal/utils"
	"go-crud-challenge/internal/adapters/gormdb"
	httpAdapter "go-crud-challenge/internal/adapters/http" // Alias to avoid naming conflict
	// "go-crud-challenge/internal/adapters/memory"
	"go-crud-challenge/internal/service"


	"github.com/gorilla/mux"
)

func main() {
	db:=utils.InitDB()
	// you can use these InMemoryy Database
	// repo := memory.NewInMemoryPersonRepository()
	repo:=gormdb.NewPostgresPersonRepository(db)
	personService := service.NewPersonService(repo)
	personHandler := httpAdapter.NewPersonHandler(personService) // Use httpAdapter here

	r := mux.NewRouter()
	r.HandleFunc("/person", personHandler.GetPersons).Methods("GET")
	r.HandleFunc("/person/{id}", personHandler.GetPerson).Methods("GET") // Renamed to GetPerson to avoid handler conflict
	r.HandleFunc("/person", personHandler.CreatePerson).Methods("POST")
	r.HandleFunc("/person/{id}", personHandler.UpdatePerson).Methods("PUT")
	r.HandleFunc("/person/{id}", personHandler.DeletePerson).Methods("DELETE")

// Middleware to handle CORS
	r.Use(corsMiddleware)

	// Catch-all route for non-existing endpoints
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r)) // Listen and serve on port 8080
}

// CORS Middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins, adjust as needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Not Found Handler
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Not Found", http.StatusNotFound)
}
