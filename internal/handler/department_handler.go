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

func (h *DepartmentHandler) getIdParameter(w http.ResponseWriter, r *http.Request) (uint, error) {
	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil || id <= 0 {
		http.Error(w, "id must be a positive number", http.StatusBadRequest)
		return 0, err
	}
	return uint(id), nil
}

func (h *DepartmentHandler) GetDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := h.getIdParameter(w, r)
	if err != nil {
		log.Println(err)
	}

	depth := 1
	depthString := r.URL.Query().Get("depth")

	if depthString != "" {
		depth, err = strconv.Atoi(depthString)
		if err != nil || depth <= 0 {
			http.Error(w, "depth must be a positive number", http.StatusBadRequest)
			log.Println(err)
			return
		}
		if depth > 5 {
			http.Error(w, "depth must not exceed 5", http.StatusBadRequest)
			log.Println(err)
			return
		}
	}

	includeEmployees := true
	includeString := r.URL.Query().Get("include_employees")
	if includeString != "" {
		includeEmployees, err = strconv.ParseBool(includeString)
		if err != nil {
			http.Error(w, "include_employee parameter must be true or false", http.StatusBadRequest)
			log.Println(err)
			return
		}
	}

	var department *models.Department
	department, err = h.departmentService.GetDepartment(id, depth, includeEmployees)
	if err != nil {
		http.Error(w, "an error occurred while retrieving departments", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response := dto.ToDepartmentDetailResponse(department)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *DepartmentHandler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req dto.CreateDepartmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "an error occurred with the request body", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	department := &models.Department{
		Name:     req.Name,
		ParentID: req.ParentID,
	}

	err := h.departmentService.CreateDepartment(department)
	if err != nil {
		http.Error(w, "an error occurred while creating department", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response := dto.ToDepartmentResponse(department)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, "an error occurred with the response", http.StatusInternalServerError)
		log.Println(err)
	}
}

func (h *DepartmentHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := h.getIdParameter(w, r)
	if err != nil {
		log.Println(err)
	}

	var req dto.EmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "an error occurred with the request body", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	parsed, err := req.ParseHiredAt()
	if err != nil {
		http.Error(w, "the hired_at field must be in the format: dd/mm/yyyy", http.StatusBadRequest)
		log.Println(err)
		return
	}

	employee := &models.Employee{
		FullName: req.FullName,
		Position: req.Position,
		HiredAt:  parsed,
	}

	if err = h.employeeService.CreateEmployee(employee, id); err != nil {
		http.Error(w, "an error occurred while creating the new employee", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response := dto.ToEmployeeResponse(employee)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, "an error occurred with the response", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (h *DepartmentHandler) UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := h.getIdParameter(w, r)
	if err != nil {
		log.Println(err)
	}

	var req dto.UpdateDepartmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "an error occurred with the request body", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	department := &models.Department{}
	if req.Name != nil {
		department.Name = *req.Name
	}
	department.ParentID = req.ParentId

	updateDepartment, err := h.departmentService.UpdateDepartment(uint(id), department)
	if err != nil {
		http.Error(w, "an error occurred while updating the department", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	response := dto.ToDepartmentResponse(updateDepartment)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		http.Error(w, "an error occurred with the response", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}

func (h *DepartmentHandler) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := h.getIdParameter(w, r)
	if err != nil {
		log.Println(err)
	}

	mode := r.URL.Query().Get("mode")
	var reassignId int

	if mode == "" {
		http.Error(w, "mode parameter is required", http.StatusBadRequest)
		return
	}

	if mode == "reassign" {
		reassignIdString := r.URL.Query().Get("reassign_id")

		if reassignIdString != "" {
			reassignId, err = strconv.Atoi(reassignIdString)
			if err != nil || reassignId <= 0 {
				http.Error(w, "reassign_id must be a positive number", http.StatusBadRequest)
				log.Println(err)
				return
			}
		} else {
			http.Error(w, "reassign_id is requierd if mode is \"reassign\"", http.StatusBadRequest)
			log.Println(err)
			return
		}
	}

	if err = h.departmentService.DeleteDepartment(id, mode, uint(reassignId)); err != nil {
		http.Error(w, "an error occured while deleting the department", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
