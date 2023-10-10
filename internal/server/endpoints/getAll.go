package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/feedback/internal/model"
	"net/http"
)

func (h *handlers) GetAllOpinions(w http.ResponseWriter, r *http.Request) {
	var prescriptions []model.Opinion
	if err := h.db.Find(&prescriptions).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(prescriptions)
}
