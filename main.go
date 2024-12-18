package main

import (
	"EcommersAPIHP/app"
	"EcommersAPIHP/controller"
	"EcommersAPIHP/exception"
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

	keranjangRepository := repository.NewKeranjangRepository()
	keranjangService := service.NewKeranjangService(keranjangRepository, userRepository, produkRepository, db, validate)
	keranjangController := controller.NewKeranjangController(keranjangService)

	detailPesananRepository := repository.NewDetailPesananRepository()
	detailPesananService := service.NewDetailPesananService(detailPesananRepository, pesananRepository, userRepository, produkRepository, db, validate)
	detailPesananController := controller.NewDetailPesananController(detailPesananService)

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
	router.POST("/api/keranjangs", keranjangController.Create)
	router.GET("/api/keranjangs", keranjangController.FindAll)
	router.GET("/api/keranjangs/:keranjangId", keranjangController.FindById)
	router.PUT("/api/keranjangs/:keranjangId", keranjangController.Update)
	router.DELETE("/api/keranjangs/:keranjangId", keranjangController.Delete)
	router.POST("/api/detailPesanans", detailPesananController.Create)
	router.GET("/api/detailPesanans", detailPesananController.FindAll)
	router.GET("/api/detailPesanans/:detailPesananId", detailPesananController.FindById)
	router.PUT("/api/detailPesanans/:detailPesananId", detailPesananController.Update)
	router.DELETE("/api/detailPesanans/:detailPesananId", detailPesananController.Delete)

	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
