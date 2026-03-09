package handler

import (
	"department-api/internal/dto"
	"department-api/internal/models"
	"department-api/internal/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DepartmentHandler struct {
	departmentService service.DepartmentService
	employeeService   service.EmployeeService
}

func NewDepartmentHandler(depService service.DepartmentService, emplSerice service.EmployeeService) *DepartmentHandler {
	return &DepartmentHandler{
		departmentService: depService,
		employeeService:   emplSerice,
	}
}

func (h *DepartmentHandler) GetDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	depth := 1
	depthString := r.URL.Query().Get("depth")

	if depthString != "" {
		depth, err = strconv.Atoi(depthString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	includeEmployees := true
	includeString := r.URL.Query().Get("include_employees")
	if includeString != "" {
		includeEmployees, err = strconv.ParseBool(includeString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	var department *models.Department

	department, err = h.departmentService.GetDepartment(uint(id), depth, includeEmployees)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(department)
}

func (h *DepartmentHandler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req dto.CreateDepartmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}

	department := &models.Department{
		Name:     req.Name,
		ParentID: req.ParentID,
	}

	err := h.departmentService.CreateDepartment(department)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.ToDepartmentResponce(department)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		log.Println(err)
	}
}

func (h *DepartmentHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idString := vars["id"]
	fmt.Println(idString)
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req dto.EmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(err)
	}

	employee := &models.Employee{
		FullName: req.FullName,
		Position: req.Position,
		HiredAt:  req.HiredAt,
	}

	err = h.employeeService.CreateEmployee(employee, uint(id))
	response := dto.ToEmployeeResponce(employee)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		log.Println(err)
	}
}
