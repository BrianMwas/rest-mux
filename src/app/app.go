package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/brianmwas/rest-mux/src/routes"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//App is Main Application of rest nux contains all functionality to initialize the application.
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) initializeRoutes() {
	routes.InitializeAllRoutes(a)
}

// Initialize is the db connection method.
func (a *App) Initialize(user string, password string, dbname string, port, host string) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	fmt.Printf("connection string %*s", 20, connectionString)
	var err error
	a.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Printf("format String %s", err)
		// Stop the server and print the problem
		panic("Failed to connect to the Database closing databasee")

	}

	a.Router = mux.NewRouter()
	db, dbError := a.DB.DB()
	if dbError != nil {
		fmt.Printf("DB connections fail")
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(100)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(10)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(time.Hour)
	// initialize routes to connect the App to backend.
	a.initializeRoutes()
}

// Run will open the application as a server address.
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}
