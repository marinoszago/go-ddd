// Package main contains the main api app logic and acts as a starter point
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"

	geolocationservice "go-ddd/application/geolocation"
	"go-ddd/config"
	"go-ddd/domain/geolocation"
	"go-ddd/infrastructure/persistence/mysql"
	geolocationrepository "go-ddd/infrastructure/persistence/mysql/geolocation"
	geolocationcontroller "go-ddd/interface/controllers/geolocation"
	geolocationrouter "go-ddd/interface/routes/geolocation"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Printf("[err] set up configuration: %s", err.Error())
		os.Exit(1)
	}

	app := NewApp(cfg)
	if err = app.Run(); err != nil {
		log.Printf("[err] app run: %s", err.Error())
		os.Exit(1)
	}
}

// App is the main app.
type App struct {
	cfg        *config.Config
	dbClient   mysql.Client
	router     *httprouter.Router
	repository geolocation.Repository
}

// NewApp creates a new App.
func NewApp(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

// Run runs the app.
func (app App) Run() error {
	fmt.Println("[RUN] >>> application started")

	err := app.initDB(app.cfg)
	if err != nil {
		return fmt.Errorf("[err] set up database: %w\n", err)
	}

	app.registerRepository()

	err = app.startServer()
	if err != nil {
		return fmt.Errorf("[err] set up server: %w", err)
	}

	return nil
}

func (app *App) initDB(cfg *config.Config) error {
	client, err := mysql.NewClient(cfg.Database.ConnectionString)
	if err != nil {
		return fmt.Errorf("[err] mysql client initialization err: %w", err)
	}

	client.UseDatabase(app.cfg.Database.Name)

	app.dbClient = client

	return nil
}

func (app App) startServer() error {
	log.Printf("[RUN] Server running at %s://%s:%d/\n",
		app.cfg.Server.Protocol, app.cfg.Server.Host, app.cfg.Server.Port)

	app.registerGeolocationRoutes()

	return http.ListenAndServe(fmt.Sprintf(":%d", app.cfg.Server.Port), app.router)
}

func (app *App) registerGeolocationRoutes() {
	service := geolocationservice.NewService(app.repository)
	controller := geolocationcontroller.NewController(service)

	gr := geolocationrouter.NewGeolocationRouter(controller)

	app.router = gr.Routes()
}

func (app *App) registerRepository() {
	app.repository = geolocationrepository.NewRepository(app.dbClient)
}
