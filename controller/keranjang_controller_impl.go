package controller

import (
	"EcommersAPIHP/helper"
	"EcommersAPIHP/model/web"
	"EcommersAPIHP/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type KeranjangControllerImpl struct {
	KeranjangService service.KeranjangService
}

func NewKeranjangController(keranjangController service.KeranjangService) KeranjangController {
	return &KeranjangControllerImpl{
		KeranjangService: keranjangController,
	}
}

func (controller *KeranjangControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	keranjangCreateRequest := web.KeranjangCreateRequest{}
	helper.ReadRequestBody(request, &keranjangCreateRequest)

	keranjangResponse := controller.KeranjangService.Create(request.Context(), keranjangCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Data: keranjangResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *KeranjangControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	keranjangResponse := controller.KeranjangService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Data: keranjangResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *KeranjangControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	keranjangId := params.ByName("keranjangId")
	id, err := strconv.Atoi(keranjangId)
	helper.PanicIfError(err)

	keranjangResponse := controller.KeranjangService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Data: keranjangResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *KeranjangControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	keranjangUpdateRequest := web.KeranjangUpdateRequest{}
	helper.ReadRequestBody(request, &keranjangUpdateRequest)

	keranjangId := params.ByName("keranjangId")
	id, err := strconv.Atoi(keranjangId)
	helper.PanicIfError(err)

	keranjangUpdateRequest.Id = id
	keranjangResponse := controller.KeranjangService.Update(request.Context(), keranjangUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Data: keranjangResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *KeranjangControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	keranjangId := params.ByName("keranjangId")
	id, err := strconv.Atoi(keranjangId)
	helper.PanicIfError(err)

	controller.KeranjangService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Data: "success",
	}
	helper.WriteResponBody(writer, webResponse)
}
