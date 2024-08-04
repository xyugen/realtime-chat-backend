package main

import (
	"database/sql"
	"log"

	"github.com/xyugen/realtime-chat-backend/cmd/api"
	"github.com/xyugen/realtime-chat-backend/config"
	"github.com/xyugen/realtime-chat-backend/db"
)

// func corsMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		if r.Method == "OPTIONS" {
// 			return
// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }

func main() {
	db, err := db.NewMySQLiteStorage(config.Envs.TursoDatabaseUrl, config.Envs.TursoAuthToken)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}
