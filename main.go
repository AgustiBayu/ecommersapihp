package main

import (
	"EcommersAPIHP/app"
	"EcommersAPIHP/controller"
	"EcommersAPIHP/helper"
	"EcommersAPIHP/repository"
	"EcommersAPIHP/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	validate := validator.New()
	db := app.NewDB()
	produkRepository := repository.NewProdukRepository()
	produkService := service.NewProdukService(produkRepository, db, validate)
	produkController := controller.NewProdukController(produkService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	pesananRepository := repository.NewPesananRepository()
	pesananService := service.NewPesananService(pesananRepository, userRepository, db, validate)
	pesananController := controller.NewPesananController(pesananService)

	router := httprouter.New()
	router.POST("/api/produks", produkController.Create)
	router.GET("/api/produks", produkController.FindAll)
	router.GET("/api/produks/:produkId", produkController.FindById)
	router.PUT("/api/produks/:produkId", produkController.Update)
	router.DELETE("/api/produks/:produkId", produkController.Delete)
	router.POST("/api/users/login", userController.Login)
	router.POST("/api/users/register", userController.Register)
	router.POST("/api/pesanans", pesananController.Create)
	router.GET("/api/pesanans", pesananController.FindAll)
	router.GET("/api/pesanans/:pesananId", pesananController.FindById)
	router.PUT("/api/pesanans/:pesananId", pesananController.Update)
	router.DELETE("/api/pesanans/:pesananId", pesananController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
