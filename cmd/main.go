package main

import (
	"department-api/internal/config"
	"department-api/internal/database"
	"department-api/internal/handler"
	"department-api/internal/repository"
	"department-api/internal/service"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config := config.LoadConfig()
	db := database.Connect(config)
	depRepo := repository.NewDepartmentRepository(db)
	emplRepo := repository.NewEmployeeRepository(db)

	depService := service.NewDepartmentService(depRepo, emplRepo)
	emplService := service.NewEmployeeService(depRepo, emplRepo)
	handler := handler.NewDepartmentHandler(depService, emplService)

	router := mux.NewRouter()

	router.HandleFunc("/departments", handler.CreateDepartment).Methods("POST")
	router.HandleFunc("/departments/{id}", handler.GetDepartment).Methods("GET")
	router.HandleFunc("/departments/{id}", handler.UpdateDepartment).Methods("PATCH")
	router.HandleFunc("/departments/{id}", handler.DeleteDepartment).Methods("DELETE")
	router.HandleFunc("/departments/{id}/employees", handler.CreateEmployee).Methods("POST")

	http.ListenAndServe(":"+config.AppPort, router)
}
