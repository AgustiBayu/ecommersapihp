package main

import (
	"EcommersAPIHP/app"
	"EcommersAPIHP/controller"
	"EcommersAPIHP/exception"
	"EcommersAPIHP/helper"
	"EcommersAPIHP/middleware"
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
	router.POST("/api/users/login", userController.Login)
	router.POST("/api/users/register", userController.Register)

	router.POST("/api/produks", middleware.BasicAuth(userService, produkController.Create))
	router.GET("/api/produks", middleware.BasicAuth(userService, produkController.FindAll))
	router.GET("/api/produks/:produkId", middleware.BasicAuth(userService, produkController.FindById))
	router.PUT("/api/produks/:produkId", middleware.BasicAuth(userService, produkController.Update))
	router.DELETE("/api/produks/:produkId", middleware.BasicAuth(userService, produkController.Delete))
	router.POST("/api/pesanans", middleware.BasicAuth(userService, pesananController.Create))
	router.GET("/api/pesanans", middleware.BasicAuth(userService, pesananController.FindAll))
	router.GET("/api/pesanans/:pesananId", middleware.BasicAuth(userService, pesananController.FindById))
	router.PUT("/api/pesanans/:pesananId", middleware.BasicAuth(userService, pesananController.Update))
	router.DELETE("/api/pesanans/:pesananId", middleware.BasicAuth(userService, pesananController.Delete))
	router.POST("/api/keranjangs", middleware.BasicAuth(userService, keranjangController.Create))
	router.GET("/api/keranjangs", middleware.BasicAuth(userService, keranjangController.FindAll))
	router.GET("/api/keranjangs/:keranjangId", middleware.BasicAuth(userService, keranjangController.FindById))
	router.PUT("/api/keranjangs/:keranjangId", middleware.BasicAuth(userService, keranjangController.Update))
	router.DELETE("/api/keranjangs/:keranjangId", middleware.BasicAuth(userService, keranjangController.Delete))
	router.POST("/api/detailPesanans", middleware.BasicAuth(userService, detailPesananController.Create))
	router.GET("/api/detailPesanans", middleware.BasicAuth(userService, detailPesananController.FindAll))
	router.GET("/api/detailPesanans/:detailPesananId", middleware.BasicAuth(userService, detailPesananController.FindById))
	router.PUT("/api/detailPesanans/:detailPesananId", middleware.BasicAuth(userService, detailPesananController.Update))
	router.DELETE("/api/detailPesanans/:detailPesananId", middleware.BasicAuth(userService, detailPesananController.Delete))

	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
