package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	_ "lab-server/docs" // замените на ваш путь к docs
	"lab-server/internal/config"
	"lab-server/internal/router"
	"log"
	"net/url"
)

// @title           My API
// @version         1.0
// @description     This is a sample server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @host      localhost:8080
// @BasePath  /

func main() {

	// Чтение конфигурации
	cfg := config.LoadConfig()

	u, err := url.Parse(cfg.DBConnStr)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", u.String())
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	log.Fatal(router.SetupEngine(db).Run())
}
