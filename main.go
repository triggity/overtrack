package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	log "github.com/sirupsen/logrus"
)

var (
	address        = flag.String("address", fmt.Sprintf(":%s", getEnv("PORT", "8000")), "address to listen on")
	dbURL          = flag.String("dbUrl", getEnv("DATABASE_URL", "postgres://postgres@127.0.0.1:5432/postgres?sslmode=disable"), "database connection string")
	migrationsPath = flag.String("migrations", getEnv("MIGRATIONS_PATH", "migrations"), "path to migrations folder")
)

func getEnv(key string, backup string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return backup
	}
	return value
}
func getDB() *sqlx.DB {
	log.Infof("connecting to postgres at %s", *dbURL)
	db, err := sqlx.Open("postgres", *dbURL)
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})

	if err != nil {
		log.Panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprint("file://%s", *migrationsPath),
		"postgres", driver)

	if err != nil {
		log.Panic(err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Panic(err)
	}
	return db
}

func main() {
	flag.Parse()
	db := getDB()
	r := mux.NewRouter()
	Server(r, db)
	w := log.New().Writer()
	defer w.Close()
	loggedHandler := handlers.LoggingHandler(w, r)

	log.Infof("starting overtrack at %s", *address)
	log.Fatal(http.ListenAndServe(*address, loggedHandler))

}
