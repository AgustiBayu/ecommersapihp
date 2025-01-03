package controller

import (
	"EcommersAPIHP/helper"
	"EcommersAPIHP/model/web"
	"EcommersAPIHP/service"
	"encoding/base64"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Register(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRegisterRequest := web.UserRegisterRequest{}
	helper.ReadRequestBody(request, &userRegisterRequest)

	userResponse := controller.UserService.Register(request.Context(), userRegisterRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Data: userResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := web.UserLoginRequest{}
	helper.ReadRequestBody(request, &userLoginRequest)
	userResponse, err := controller.UserService.Login(request.Context(), userLoginRequest)
	helper.PanicIfError(err)

	token := base64.StdEncoding.EncodeToString([]byte(userLoginRequest.Email + ":" + userLoginRequest.Password))
	writer.Header().Set("Authorization", "Basic "+token)
	webResponse := web.WebResponse{
		Code: 200,
		Data: userResponse,
	}
	helper.WriteResponBody(writer, webResponse)
}
