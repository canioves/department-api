package main

import (
	"department-api/internal/config"
	"department-api/internal/database"
	"department-api/internal/handler"
	"department-api/internal/repository"
	"department-api/internal/service"
	"net/http"
)

func main() {
	config := config.LoadConfig()
	db := database.Connect(config)
	depRepo := repository.NewDepartmentRepository(db)
	emplRepo := repository.NewEmployeeRepository(db)

	depService := service.NewDepartmentService(depRepo, emplRepo)
	handler := handler.NewDepartmentHandler(depService)

	mux := http.NewServeMux()

	mux.HandleFunc("/departments/", handler.GetDepartment)

	http.ListenAndServe(":8080", mux)

}
