package config

import (
	"os"
	"time"

	migrations "github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"github.com/kpango/glg"
)

type Config struct {
	DB *pg.DB
}

var Get *Config

// Load sets the db configurations and runs the migrations
func Load() {
	Get = &Config{}
	setConfig()
	col := migrations.NewCollection()
	err := col.DiscoverSQLMigrations("migrations")
	if err != nil {
		glg.Fatal("could not load migrations: err ", err.Error())
	}
	if os.Getenv("ENV") == "" {
		time.Sleep(5 * time.Second)
	}
	_, _, err = col.Run(Get.DB, "init")
	if err != nil {
		glg.Fatal("could not run migrations: err ", err.Error())
	}
	glg.Info("[POSTGRES] INIT OK ")
	v0, v1, err := col.Run(Get.DB, "up")
	if err != nil {
		glg.Fatal("could not run migrations: err ", err.Error())
	}
	glg.Infof("[POSTGRES] Migrated from '%v' to '%v'", v0, v1)
}

func setConfig() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL != "" {
		opts, err := pg.ParseURL(dbURL)
		if err != nil {
			glg.Fatal("load err: ", err.Error())
		}
		Get.DB = pg.Connect(opts)
		return
	}
	Get.DB = pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
	})
}
