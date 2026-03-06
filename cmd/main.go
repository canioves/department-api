package main

import (
	"department-api/internal/config"
	"department-api/internal/database"
	"department-api/internal/repository"
	"fmt"
)

func main() {
	config := config.LoadConfig()
	db := database.Connect(config)
	r := repository.NewDepartmentRepository(db)
	deps, _ := r.GetChildren(1)
	for _, x := range deps {
		fmt.Printf("%+v\n", x)
	}
}
