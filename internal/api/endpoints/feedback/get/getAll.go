package get

import (
	"encoding/json"
	"github.com/szmulinho/feedback/internal/database"
	"github.com/szmulinho/feedback/internal/model"
	"net/http"
)

func GetAllOpinions(w http.ResponseWriter, r *http.Request) {
	var prescriptions []model.Opinion
	if err := database.DB.Find(&prescriptions).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(prescriptions)
}
