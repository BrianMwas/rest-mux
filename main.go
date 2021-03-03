package main

import (
	"os"

	"github.com/brianmwas/rest-mux/src/app"
)

func main() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "localhost:3306")
	os.Setenv("DB_NAME", "mux")
	a := app.App{}

	a.Initialize(
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
	)

	autoMigrate(a)

	a.Run(":8090")
}
