package handler

import (
	"department-api/internal/dto"
	"department-api/internal/models"
	"department-api/internal/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type DepartmentHandler struct {
	departmentService service.DepartmentService
	employeeService   service.EmployeeService
}

func NewDepartmentHandler(depService service.DepartmentService, emplService service.EmployeeService) *DepartmentHandler {
	return &DepartmentHandler{
		departmentService: depService,
		employeeService:   emplService,
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

	response := dto.ToDepartmentResponse(department)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		log.Println(err)
	}
}

func (h *DepartmentHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req dto.EmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	employee := &models.Employee{
		FullName: req.FullName,
		Position: req.Position,
		HiredAt:  req.HiredAt,
	}

	err = h.employeeService.CreateEmployee(employee, uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.ToEmployeeResponse(employee)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		log.Println(err)
	}
}

func (h *DepartmentHandler) UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	idString := vars["id"]

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req dto.UpdateDepartmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	department := &models.Department{}
	if req.Name != nil {
		department.Name = *req.Name
	}
	department.ParentID = req.ParentId

	updateDepartment, err := h.departmentService.UpdateDepartment(uint(id), department)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := dto.ToDepartmentResponse(updateDepartment)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
