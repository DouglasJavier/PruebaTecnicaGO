package main

import (
	"net/http"

	"github.com/douglascolque/go-gorm-rest-api/db"
	"github.com/douglascolque/go-gorm-rest-api/models"
	"github.com/douglascolque/go-gorm-rest-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	db.DB.AutoMigrate(models.Category{})
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.Task_Category{})
	router := mux.NewRouter()
	router.HandleFunc("/ping", routes.PingHandler)

	//rutas para user
	router.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	//rutas para task
	//task routes
	router.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	router.HandleFunc("/tasks", routes.PostTaskHandler).Methods("POST")
	router.HandleFunc("/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")

	http.ListenAndServe(":8082", enableCORS(router))
}
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Permitir cualquier origen, puedes cambiarlo a "http://localhost:4200"
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Si la solicitud es un preflight, responder y no continuar
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Continuar al siguiente manejador
		next.ServeHTTP(w, r)
	})
}
