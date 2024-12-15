package controller

import (
	"EcommersAPIHP/helper"
	"EcommersAPIHP/model/web"
	"EcommersAPIHP/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProdukControllerImpl struct {
	ProdukService service.ProdukService
}

func NewProdukController(produkService service.ProdukService) ProdukController {
	return &ProdukControllerImpl{
		ProdukService: produkService,
	}
}

func (controller *ProdukControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	produkCreateRequest := web.ProdukCreateRequest{}
	helper.ReadRequestBody(request, &produkCreateRequest)

	produkResponse := controller.ProdukService.Create(request.Context(), produkCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Data: produkResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *ProdukControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	produkResponse := controller.ProdukService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Data: produkResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *ProdukControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	produkId := params.ByName("produkId")
	id, err := strconv.Atoi(produkId)
	helper.PanicIfError(err)

	produkResponse := controller.ProdukService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Data: produkResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *ProdukControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	produkUpdateRequest := web.ProdukUpdateRequest{}
	helper.ReadRequestBody(request, &produkUpdateRequest)

	produkId := params.ByName("produkId")
	id, err := strconv.Atoi(produkId)
	helper.PanicIfError(err)
	produkUpdateRequest.Id = id

	produkResponse := controller.ProdukService.Update(request.Context(), produkUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Data: produkResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *ProdukControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	produkId := params.ByName("produkId")
	id, err := strconv.Atoi(produkId)
	helper.PanicIfError(err)

	controller.ProdukService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Data: "success",
	}
	helper.WriteResponBody(writer, webResponse)
}
