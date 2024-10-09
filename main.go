package main

import (
	"net/http"

	"github.com/douglascolque/go-gorm-rest-api/db"
	"github.com/douglascolque/go-gorm-rest-api/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection()
	router := mux.NewRouter()
	router.HandleFunc("/ping", routes.PingHandler)

	http.ListenAndServe(":3000", enableCORS(router))
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
