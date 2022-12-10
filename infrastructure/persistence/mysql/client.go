// Package mysql is a infrastructure level package that holds information about a new mysql Client
package mysql

import (
	"fmt"

	gormysql "gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-ddd/domain/geolocation"
)

// Client is mysql client to interact with the DB.
type Client struct {
	gorm.DB
}

// NewClient creates a new mysql client and tries to establish a connection with the mysql.
func NewClient(dataSource string) (Client, error) {
	db, err := gorm.Open(gormysql.Open(dataSource), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return Client{}, fmt.Errorf("open db connection: %w", err)
	}

	return Client{*db}, nil
}

// RunGormEntityMigrations runs the gorm migrations.
func (c Client) RunGormEntityMigrations() error {
	err := c.Migrator().DropTable(&geolocation.GeoData{})
	if err != nil {
		return err
	}

	err = c.Migrator().AutoMigrate(&geolocation.GeoData{})
	if err != nil {
		return err
	}

	return nil
}

func (c Client) CreateDatabase(dbName string) {
	c.DB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", dbName))
}

func (c Client) UseDatabase(dbName string) {
	c.DB.Exec(fmt.Sprintf("USE %s;", dbName))
}
