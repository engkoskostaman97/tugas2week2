package handlers

import (
	categorydto "dumbflix/dto/category"
	dto "dumbflix/dto/result"

	"dumbflix/models"
	"dumbflix/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


type handlerCategory struct {
	CategoryRepository repositories.CategoryRepository
}

func HandlerCategory(CategoryRepository repositories.CategoryRepository) *handlerCategory {
	return &handlerCategory{CategoryRepository}
}

func (h *handlerCategory) FindCategorys(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	films, err := h.CategoryRepository.FindCategorys()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: films}
	json.NewEncoder(w).Encode(response)
}

func (h *handlerCategory) GetCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var category models.Category
	category, err := h.CategoryRepository.GetCategory(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponseCategory(category)}
	json.NewEncoder(w).Encode(response)
}

func convertResponseCategory(u models.Category) categorydto.CategoryResponse {
	return categorydto.CategoryResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
