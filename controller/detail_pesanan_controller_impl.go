package controller

import (
	"EcommersAPIHP/helper"
	"EcommersAPIHP/model/web"
	"EcommersAPIHP/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type DetailPesananControllerImpl struct {
	DetailPesananService service.DetailPesananService
}

func NewDetailPesananController(detailPesananService service.DetailPesananService) DetailPesananController {
	return &DetailPesananControllerImpl{
		DetailPesananService: detailPesananService,
	}
}

func (controller *DetailPesananControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	detailPesananCreateRequest := web.DetailPesananCreateRequest{}
	helper.ReadRequestBody(request, &detailPesananCreateRequest)

	detailResponse := controller.DetailPesananService.Create(request.Context(), detailPesananCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Data: detailResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *DetailPesananControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	detailResponse := controller.DetailPesananService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Data: detailResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *DetailPesananControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	detailPesananId := params.ByName("detailPesananId")
	id, err := strconv.Atoi(detailPesananId)
	helper.PanicIfError(err)

	detailResponse := controller.DetailPesananService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Data: detailResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *DetailPesananControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	detailPesananUpdateRequest := web.DetailPesananUpdateRequest{}
	helper.ReadRequestBody(request, &detailPesananUpdateRequest)
	detailPesananId := params.ByName("detailPesananId")
	id, err := strconv.Atoi(detailPesananId)
	helper.PanicIfError(err)

	detailPesananUpdateRequest.Id = id
	detailResponse := controller.DetailPesananService.Update(request.Context(), detailPesananUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Data: detailResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *DetailPesananControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	detailPesananId := params.ByName("detailPesananId")
	id, err := strconv.Atoi(detailPesananId)
	helper.PanicIfError(err)

	controller.DetailPesananService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Data: "success",
	}
	helper.WriteResponBody(writer, webResponse)
}
