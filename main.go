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

	router := httprouter.New()
	router.POST("/api/produks", produkController.Create)
	router.GET("/api/produks", produkController.FindAll)
	router.GET("/api/produks/:produkId", produkController.FindById)
	router.PUT("/api/produks/:produkId", produkController.Update)
	router.DELETE("/api/produks/:produkId", produkController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
