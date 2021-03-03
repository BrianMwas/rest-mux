package controllers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/brianmwas/rest-mux/helpers"
	"github.com/brianmwas/rest-mux/src/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetProductOr404(w http.ResponseWriter, r *http.Request) {
	db := &gorm.DB{}
	vars := mux.Vars(r)

	if vars == nil {
		helpers.RespondError(w, http.StatusBadRequest, "Please enter")
	}

	product := models.Product{}
	if err := db.First(&product, models.Product{ID: vars["id"]}).Error; err != nil {
		helpers.RespondError(w, http.StatusNotFound, err.Error())
	}

	helpers.RespondJSON(w, http.StatusOK, &product)
}

func GetProductListOr404(w http.ResponseWriter, r *http.Request) {
	// Route parameters
	vars := mux.Vars(r)
	page, errPage := strconv.Atoi(vars["page"])
	size, errSize := strconv.Atoi(vars["size"])
	limit, errLimit := strconv.Atoi(vars["limit"])

	if errPage != nil {
		helpers.RespondError(w, http.StatusBadRequest, "Invalid page")
	}

	if errSize != nil {
		helpers.RespondError(w, http.StatusBadRequest, "Invalid size")
	}

	if errLimit != nil {
		helpers.RespondError(w, http.StatusBadRequest, "Invalid limit")
	}

	products := []models.Product{}
	db := &gorm.DB{}
	// Paginate data
	if page == 0 {
		page = 1
	}

	switch {
	case size > 100:
		size = 100
	case size <= 0:
		size = 10
	}

	offset := (page - 1) * size
	dbProduct := db.Find(&products).Offset(offset).Limit(limit).Order("price")
	if err := dbProduct.Error; err != nil {
		helpers.RespondError(w, http.StatusNotFound, err.Error())

	}

	helpers.RespondJSON(w, http.StatusOK, &products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := helpers.DecodeJsonBody(w, r, &product)

	if err != nil {
		var mr *helpers.MalformedRequest
		if errors.As(err, &mr) {
			helpers.RespondError(w, mr.Status, mr.Msg)

		} else {
			log.Println(err.Error())
			helpers.RespondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		return
	}

	helpers.RespondJSON(w, http.StatusCreated, helpers.SuccessResponse{Msg: "Successfully created product"})
}
