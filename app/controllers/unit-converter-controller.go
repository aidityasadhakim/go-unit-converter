package controllers

import (
	"net/http"

	"github.com/aidityasadhakim/go-unit-converter/app/pages"
	"github.com/aidityasadhakim/go-unit-converter/app/services"
)

type UnitConverterController struct{}

func NewUnitConverterController(service *services.UnitConverterService) *UnitConverterController {
	return &UnitConverterController{}
}

func (uc *UnitConverterController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	pages.Home(w, pages.HomeData{Title: "Home"})
}