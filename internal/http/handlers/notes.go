package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"example.com/goprac11-borisovda/internal/core"
	"example.com/goprac11-borisovda/internal/repo"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Repo *repo.NoteRepoMem
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}

// CreateNote godoc
// @Summary      Создать заметку
// @Tags         notes
// @Accept       json
// @Produce      json
// @Param        input  body     core.NoteCreate  true  "Данные новой заметки"
// @Success      201    {object} core.Note
// @Failure      400    {object} map[string]string
// @Failure      500    {object} map[string]string
// @Router       /notes [post]
func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var n core.Note
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	id, _ := h.Repo.Create(n)
	n.ID = id
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(n)
}

// GetAllNotes godoc
// @Summary      Список заметок
// @Description  Возвращает список заметок с пагинацией и фильтром по заголовку
// @Tags         notes
// @Param        page   query  int     false  "Номер страницы"
// @Param        limit  query  int     false  "Размер страницы"
// @Param        q      query  string  false  "Поиск по title"
// @Success      200    {array}  core.Note
// @Header       200    {integer}  X-Total-Count  "Общее количество"
// @Failure      500    {object}  map[string]string
// @Router       /notes [get]
func (h *Handler) GetAllNotes(w http.ResponseWriter, r *http.Request) {
	notes, _ := h.Repo.GetAll()
	writeJSON(w, http.StatusOK, notes)
}

// GetNote godoc
// @Summary      Получить заметку
// @Tags         notes
// @Param        id   path   int  true  "ID"
// @Success      200  {object}  core.Note
// @Failure      404  {object}  map[string]string
// @Router       /notes/{id} [get]
func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	n, _ := h.Repo.Get(id)
	if n == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, n)
}

// UpdateNote godoc
// @Summary      Обновить заметку
// @Tags         notes
// @Accept       json
// @Param        id     path   int        true  "ID"
// @Param        input  body   core.NoteUpdate true  "Поля для обновления"
// @Success      200    {object}  core.Note
// @Failure      400    {object}  map[string]string
// @Failure      404    {object}  map[string]string
// @Router       /notes/{id} [patch]
func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	var upd core.Note
	if err := json.NewDecoder(r.Body).Decode(&upd); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	n, _ := h.Repo.Update(id, upd)
	if n == nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, n)
}

// DeleteNote godoc
// @Summary      Удалить заметку
// @Tags         notes
// @Param        id  path  int  true  "ID"
// @Success      204  "No Content"
// @Failure      404  {object}  map[string]string
// @Router       /notes/{id} [delete]
func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	h.Repo.Delete(id)
	w.WriteHeader(http.StatusNoContent)
}
