// Package main_import contains the main import app logic and acts as a starter point
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	geolocationservice "go-ddd/application/geolocation"
	"go-ddd/config"
	"go-ddd/domain/geolocation"
	"go-ddd/infrastructure/persistence/mysql"
	geolocationrepository "go-ddd/infrastructure/persistence/mysql/geolocation"
	"go-ddd/internal/import/csv"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Printf("[err] set up configuration: %s", err.Error())
		os.Exit(1)
	}

	app := NewImportApp(cfg)
	if err = app.RunImport(); err != nil {
		log.Printf("[err] app run: %s", err.Error())
		os.Exit(1)
	}
}

// ImportApp is the main import app.
type ImportApp struct {
	cfg        *config.Config
	dbClient   mysql.Client
	repository geolocation.Repository
}

// NewImportApp creates a new ImportApp.
func NewImportApp(cfg *config.Config) *ImportApp {
	return &ImportApp{cfg: cfg}
}

// RunImport runs the import app.
func (app ImportApp) RunImport() error {
	err := app.initDB(app.cfg)
	if err != nil {
		return fmt.Errorf("[err] set up database: %w\n", err)
	}

	app.registerRepository()

	if app.cfg.ImportService.CanRun() {
		err = app.runCSVImport()
		if err != nil {
			fmt.Printf("[err] import service: %v \n", err)
		}
	} else {
		fmt.Println("Could not be run, please check the configuration")
	}

	return nil
}

func (app *ImportApp) initDB(cfg *config.Config) error {
	client, err := mysql.NewClient(cfg.Database.ConnectionString)
	if err != nil {
		return fmt.Errorf("[err] mysql client initialization err: %w", err)
	}

	fmt.Println("[RUN] >>> create database started")

	client.CreateDatabase(app.cfg.Database.Name)
	client.UseDatabase(app.cfg.Database.Name)

	fmt.Println("[RUN] >>> migrations started")

	err = client.RunGormEntityMigrations()
	if err != nil {
		return fmt.Errorf("[err] gorm migrations: %w", err)
	}

	app.dbClient = client

	return nil
}

func (app *ImportApp) registerRepository() {
	app.repository = geolocationrepository.NewRepository(app.dbClient)
}

func (app ImportApp) runCSVImport() error {
	fmt.Println("[RUN] >>> import service started")

	service := geolocationservice.NewImportService(app.repository)

	parser := csv.NewParser(app.cfg.ImportService.File)

	start := time.Now()

	records, err := parser.Parse()
	if err != nil {
		fmt.Println("[RUN] >>> Finished import with error")
		return fmt.Errorf("[err] csv parser: %w", err)
	}

	stats, errors := service.Save(records)

	elapsed := time.Since(start)

	fmt.Println("[RUN] >>> Finished import")
	fmt.Printf("[RUN] >>> Time elapsed: %s\n", elapsed)
	fmt.Printf("[RUN] >>> Total rows: %v\n", len(records))
	fmt.Printf("[RUN] >>> Discared rows: %v\n", stats["discarded"])
	fmt.Printf("[RUN] >>> Accepted rows: %v\n", stats["accepted"])

	if len(errors) > 0 {
		fmt.Printf("[RUN] >>> With errors %d \n", len(errors))

		for errorEntity, errorRec := range errors {
			fmt.Println("[FAIL] >>>", errorEntity, "reason:", errorRec)
		}
	}

	return nil
}
