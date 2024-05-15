package endpoints

import (
	"bytes"
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

func TestAddOpinion(t *testing.T) {
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

	testOpinion := model.Opinion{
	}

	testOpinionJSON, err := json.Marshal(testOpinion)
	if err != nil {
		t.Fatal(err)
	}

	request, err := http.NewRequest("POST", "/add_opinion", bytes.NewBuffer(testOpinionJSON))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	h.AddOpinion(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var addedOpinionFromDB model.Opinion
	db.First(&addedOpinionFromDB, testOpinion.ID)

	assert.Equal(t, testOpinion.ID, addedOpinionFromDB.ID)
	assert.Equal(t, testOpinion.Login, addedOpinionFromDB.Login)
	assert.Equal(t, testOpinion.Comment, addedOpinionFromDB.Comment)
	assert.Equal(t, testOpinion.Rating, addedOpinionFromDB.Rating)
}
