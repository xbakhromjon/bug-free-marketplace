package handler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"golang-project-template/internal/common"
	"golang-project-template/internal/shop/app"
	"golang-project-template/internal/shop/domain"
	"golang-project-template/internal/shop/ports/rest"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service        app.ProductService
	productFactory domain.ProductFactory
}

func (p *ProductHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		rest.ErrorHandler(w, r, errors.New("id must be number"))
		return
	}
	response, err := p.service.GetOne(id)
	if err != nil {
		rest.ErrorHandler(w, r, err)
		return
	}

	render.JSON(w, r, response)
}

func (p *ProductHandler) Filter(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	priceFrom, _ := strconv.Atoi(r.URL.Query().Get("priceFrom"))
	priceTo, _ := strconv.Atoi(r.URL.Query().Get("priceTo"))
	page, err := strconv.ParseUint(r.URL.Query().Get("page"), 10, 64)
	if err != nil {
		page = 1
	}
	size, err := strconv.ParseUint(r.URL.Query().Get("size"), 10, 64)
	if err != nil {
		size = 10
	}
	searchModel := p.productFactory.CreateNewSearchModel(name, priceFrom, priceTo)

	response, err := p.service.FilterByPageable(*searchModel, *common.CreatePageableRequest(page, size))
	if err != nil {
		rest.ErrorHandler(w, r, err)
	}

	render.JSON(w, r, response)
}
