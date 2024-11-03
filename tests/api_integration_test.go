package tests

import (
	"database/sql"
	"encoding/json"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"lab-server/internal/config"
	"lab-server/internal/models/responses"
	"lab-server/internal/router"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// Fixture хранит необходимые ресурсы для тестов
type Fixture struct {
	DB     *sql.DB
	Config *config.Config
	Router *gin.Engine
}

// SetupFixture создает фикстуру и возвращает указатель на нее
func SetupFixture(t *testing.T) *Fixture {
	t.Helper()

	cfg := config.LoadConfig()
	u, err := url.Parse(cfg.DBConnStr)
	if err != nil {
		t.Fatalf("Failed to parse DB connection string: %v", err)
	}

	db, err := sql.Open("postgres", u.String())
	if err != nil {
		t.Fatalf("Failed to connect to DB: %v", err)
	}

	return &Fixture{
		DB:     db,
		Config: cfg,
		Router: router.SetupEngine(db),
	}
}

// TearDown закрывает ресурсы, когда они больше не нужны
func (f *Fixture) TearDown(t *testing.T) {
	t.Helper()
	if err := f.DB.Close(); err != nil {
		t.Fatalf("Failed to close DB: %v", err)
	}
}

// TestGetServer проверяет ответ от конечной точки /info/server
func TestGetServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fixture := SetupFixture(t)
	defer fixture.TearDown(t)

	req, _ := http.NewRequest("GET", "/info/server", nil)

	w := httptest.NewRecorder()
	fixture.Router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
	}

	var ResponseModel responses.InfoServerResponse
	err := json.Unmarshal(w.Body.Bytes(), &ResponseModel)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
}

func TestGetClient(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fixture := SetupFixture(t) // Используем фикстуру
	defer fixture.TearDown(t)  // Очистка ресурсов

	req, _ := http.NewRequest("GET", "/info/client", nil)
	req.Header.Set("User-Agent", "test-useragent")

	w := httptest.NewRecorder()
	fixture.Router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
	}

	var ResponseModel responses.InfoClientResponse
	err := json.Unmarshal(w.Body.Bytes(), &ResponseModel)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	if ResponseModel.Data.Useragent != "test-useragent" {
		t.Errorf("got header User-Agent %s, want test-useragent", ResponseModel.Data.Useragent)
	}
}

func TestGetDatabase(t *testing.T) {
	gin.SetMode(gin.TestMode)

	fixture := SetupFixture(t) // Используем фикстуру
	defer fixture.TearDown(t)  // Очистка ресурсов

	req, _ := http.NewRequest("GET", "/info/database", nil)

	w := httptest.NewRecorder()
	fixture.Router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("got status %d, want %d", w.Code, http.StatusOK)
	}

	var ResponseModel responses.InfoDatabaseResponse
	err := json.Unmarshal(w.Body.Bytes(), &ResponseModel)
	if err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}
}
