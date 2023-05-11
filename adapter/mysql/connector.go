package mysql

import (
	"database/sql"
	"github.com/golang-migrate/migrate"
	"github.com/odanaraujo/golang/clean-architecture-hex/infra"
	"log"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", infra.StringConnectionDatabase)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func RunMigrations() {
	databasaURL := infra.StringConnectionDatabase

	migrate, err := migrate.New("file://database/migrations", "mysql"+databasaURL)

	if err != nil {
		log.Println(err)
	}

	if err := migrate.Up(); err != nil {
		log.Println(err)
	}
}
