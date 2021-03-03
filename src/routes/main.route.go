package routes

import (
	"github.com/brianmwas/rest-mux/src/app"
	"github.com/brianmwas/rest-mux/src/controllers"
)

// import "github.com/brianmwas/rest-mux/src/app"

func InitializeAllRoutes(a *app.App) {
	a.Post("/product", controllers.CreateProduct)
	a.Get("/products/{id}", controllers.GetProductOr404)
	a.Get("/products", controllers.GetProductListOr404)
}
