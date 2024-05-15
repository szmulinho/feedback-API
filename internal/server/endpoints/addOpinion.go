package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/szmulinho/feedback/internal/model"
	"log"
	"net/http"
)


func (h *handlers) AddOpinion(w http.ResponseWriter, r *http.Request) {

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

	result := h.db.Create(&newOpinion)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("added new opinion %+v\n", model.Feedback)
	log.Printf("%+v", model.Feedback)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(model.Feedback)
}
