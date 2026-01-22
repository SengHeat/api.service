package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api.drsb-purchase-service/internal/domain/user"
	"api.drsb-purchase-service/pkg/response"
	"api.drsb-purchase-service/pkg/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	service   user.Service
	validator *validator.Validator
}

func NewHandler(service user.Service) *Handler {
	return &Handler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req user.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.validator.Validate(req); err != nil {
		response.ValidationError(w, err)
		return
	}

	u, err := h.service.CreateUser(r.Context(), &req)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(w, http.StatusCreated, u, "User created successfully")
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	u, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "User not found")
		return
	}

	response.Success(w, http.StatusOK, u, "")
}
