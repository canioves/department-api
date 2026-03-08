package handler

import (
	"department-api/internal/models"
	"department-api/internal/service"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type DepartmentHandler struct {
	departmentService service.DepartmentService
}

func NewDepartmentHandler(service service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{departmentService: service}
}

func (h *DepartmentHandler) GetDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idString := strings.TrimPrefix(r.URL.Path, "/departments/")
	id, _ := strconv.Atoi(idString)
	depth, _ := strconv.Atoi(r.URL.Query().Get("depth"))
	includeEmployees, _ := strconv.ParseBool(r.URL.Query().Get("include_employees"))

	var department *models.Department

	department, _ = h.departmentService.GetDepartment(uint(id), depth, includeEmployees)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(department)
}

func (h *DepartmentHandler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cpntent-Type", "application/json")
	var newDepartment struct {
		Name     string `json:"name"`
		ParentID *uint  `json:"parent_id,omitempty"`
	}
	json.NewDecoder(r.Body).Decode(&newDepartment)
	createdDepartment, _ := h.departmentService.CreateDepartment(newDepartment)
}
