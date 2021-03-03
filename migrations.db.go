package main

import (
	"fmt"

	"github.com/brianmwas/rest-mux/src/app"
	"github.com/brianmwas/rest-mux/src/models"
)

func autoMigrate(db app.App) {
	err := db.DB.AutoMigrate(&models.Product{}, &models.User{})
	if err != nil {
		fmt.Printf("Error from automigrating %s", err)
	}
}
