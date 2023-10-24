package main

import (
	"golang-crud/config"
	"golang-crud/controller"
	"golang-crud/helper"
	"golang-crud/model"
	"golang-crud/repository"
	"golang-crud/router"
	"golang-crud/service"
	"log"
	"net/http"
	"github.com/go-playground/validator/v10"
)

func main() {
	log.Println("Starting server...")
	// database
	db := config.DatabaseConnection()
	validate := validator.New()
	db.Table("tags").AutoMigrate(&model.Tags{})
	// repository
	tagsRepository := repository.NewTagsRepositoryImpl(db)
	// service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)
	// controller
	tagsController := controller.NewTagController(tagsService)
	// router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr: ":8888",
		Handler:  routes,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}