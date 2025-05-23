package controllers

import (
	"net/http"
	"strconv"

	"github.com/aidityasadhakim/go-unit-converter/app/pages"
	"github.com/aidityasadhakim/go-unit-converter/app/services"
)

type UnitConverterController struct {
	service *services.UnitConverterService
}

func NewUnitConverterController(service *services.UnitConverterService) *UnitConverterController {
	return &UnitConverterController{service: service}
}

func (uc *UnitConverterController) HomeHandler(w http.ResponseWriter, r *http.Request) {
	pages.Home(w, pages.HomeData{Title: "Home"})
}

func (uc *UnitConverterController) LengthHandler(w http.ResponseWriter, r *http.Request) {
	pages.Length(w, pages.LengthData{Title: "Length"})
}

func (uc *UnitConverterController) LengthResultHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	type parameter struct {
		Value float64
		From string
		To string
	}
	param := parameter{
		Value: func() float64 {
			v, err := strconv.ParseFloat(r.FormValue("length"), 64)
			if err != nil {
				return 0
			}
			return v
		}(),
		From: r.FormValue("from"),
		To: r.FormValue("to"),
	}
	result, err := uc.service.ConvertLength(param.Value, param.From, param.To)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pages.LengthResult(w, pages.LengthResultData{
		Title:     "Length Result",
		Result:    strconv.FormatFloat(result, 'f', 2, 64),
		FromValue: strconv.FormatFloat(param.Value, 'f', 2, 64),
		FromUnit:  param.From,
		ToUnit:    param.To,
	})
}

func (uc *UnitConverterController) TemperatureHandler(w http.ResponseWriter, r *http.Request) {
	pages.Temperature(w, pages.TemperatureData{Title: "Temperature"})
}

func (uc *UnitConverterController) TemperatureResultHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	type parameter struct {
		Value float64
		From string
		To string
	}
	param := parameter{
		Value: func() float64 {
			v, err := strconv.ParseFloat(r.FormValue("temperature"), 64)
			if err != nil {
				return 0
			}
			return v
		}(),
		From: r.FormValue("from"),
		To: r.FormValue("to"),
	}
	result, err := uc.service.ConvertTemperature(param.Value, param.From, param.To)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pages.TemperatureResult(w, pages.TemperatureResultData{
		Title:     "Temperature Result",
		Result:    strconv.FormatFloat(result, 'f', 2, 64),
		FromValue: strconv.FormatFloat(param.Value, 'f', 2, 64),
		FromUnit:  param.From,
		ToUnit:    param.To,
	})
}

func (uc *UnitConverterController) WeightHandler(w http.ResponseWriter, r *http.Request) {
	pages.Weight(w, pages.WeightData{Title: "Weight"})
}

func (uc *UnitConverterController) WeightResultHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	type parameter struct {
		Value float64
		From string
		To string
	}
	param := parameter{
		Value: func() float64 {
			v, err := strconv.ParseFloat(r.FormValue("weight"), 64)
			if err != nil {
				return 0
			}
			return v
		}(),
		From: r.FormValue("from"),
		To: r.FormValue("to"),
	}
	result, err := uc.service.ConvertWeight(param.Value, param.From, param.To)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	pages.WeightResult(w, pages.WeightResultData{
		Title:     "Weight Result",
		Result:    strconv.FormatFloat(result, 'f', 2, 64),
		FromValue: strconv.FormatFloat(param.Value, 'f', 2, 64),
		FromUnit:  param.From,
		ToUnit:    param.To,
	})
}
