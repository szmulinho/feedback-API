package endpoints

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/szmulinho/feedback/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetAllOpinions(t *testing.T) {
	err := godotenv.Load(".env.test")
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " dbname=" + name + " sslmode=require password=" + password + " port=" + port

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	defer db.DB()

	db.AutoMigrate(&model.Opinion{})

	h := &handlers{db: db}

	request, err := http.NewRequest("GET", "/get_all", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router := http.NewServeMux()
	router.HandleFunc("/get_all", h.GetAllOpinions)

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var responseOpinions []model.Opinion
	err = json.Unmarshal(recorder.Body.Bytes(), &responseOpinions)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 4, len(responseOpinions))

}