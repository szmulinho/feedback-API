package server

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/szmulinho/feedback/internal/server/endpoints"
	"gorm.io/gorm"

	"log"
	"net/http"
)

func Run(ctx context.Context , db *gorm.DB) {
	handler := endpoints.NewHandler(db)
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/add_opinion", handler.AddOpinion).Methods("POST")
	router.HandleFunc("/get_all", handler.GetAllOpinions).Methods("GET")
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"}),
		handlers.ExposedHeaders([]string{}),
		handlers.AllowCredentials(),
		handlers.MaxAge(86400),
	)
	go func() {
		err := http.ListenAndServe(":8083", cors(router))
		if err != nil {
			log.Fatal(err)
		}
	}()
	<-ctx.Done()

}
