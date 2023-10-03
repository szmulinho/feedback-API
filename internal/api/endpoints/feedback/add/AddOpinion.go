package add

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/feedback/internal/database"
	"github.com/szmulinho/feedback/internal/model"
	"log"
	"net/http"
)

type errResponse struct {
	Error string `json:"error"`
}

func AddOpinion(w http.ResponseWriter, r *http.Request) {

	var newOpinion model.Opinion

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
	}
	err = json.Unmarshal(buf.Bytes(), &newOpinion)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(err)
		log.Printf("Invalid body")
	}

	result := database.DB.Create(&newOpinion)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	for _, singlePresc := range model.Feed {
		fmt.Println(singlePresc)
		if singlePresc.ID == model.Feedback.ID {
			w.WriteHeader(http.StatusConflict)
			json.NewEncoder(w).Encode(errResponse{Error: fmt.Sprintf("Prescription %model.already exist", model.Feedback.ID)})
			return
		}
	}

	fmt.Printf("created new prescription %+v\n", model.Feedback)
	log.Printf("%+v", model.Feedback)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.Feedback)
}
