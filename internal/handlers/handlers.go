package handlers

import (
	"awesomeProject/internal/services"
	"awesomeProject/internal/utils"
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	s     *services.Service
	valid *validator.Validate
}

func NewHandler(s *services.Service) *Handler {
	h := &Handler{
		s:     s,
		valid: validator.New(),
	}

	if err := h.valid.RegisterValidation("full_url", utils.ValidFullUrl); err != nil {
		log.Printf("Failed to register full url validation: %v", err)
	}

	if err := h.valid.RegisterValidation("short_url", utils.ValidShortenUrl); err != nil {
		log.Printf("Failed to register short url validation: %v", err)
	}

	return h
}

// ShortenUrlReq ..
type ShortenUrlReq struct {
	Long string `json:"long" validate:"required,full_url"`
}

// ShortenUrl ..
func (h *Handler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var req ShortenUrlReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResp(w, http.StatusBadRequest, "Bad request")
		return
	}

	err = h.valid.Struct(&req)
	if err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			utils.ErrorResp(w, http.StatusForbidden, "Field validation error")
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := h.s.ShortenURL(ctx, req.Long)
	if err != nil {
		utils.ErrorResp(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	utils.SuccessResp(w, http.StatusCreated, "Link successfully shortened", result)
}

// RedirectUrlReq ..
type RedirectUrlReq struct {
	Short string `json:"short" validate:"required,short_url"`
}

// RedirectUrl ..
func (h *Handler) RedirectUrl(w http.ResponseWriter, r *http.Request) {
	shortenUrl := chi.URLParam(r, "shorten_url")
	if shortenUrl == "" {
		utils.ErrorResp(w, http.StatusBadRequest, "Bad request")
		return
	}

	req := RedirectUrlReq{
		Short: shortenUrl,
	}

	err := h.valid.Struct(&req)
	if err != nil {
		for _, err = range err.(validator.ValidationErrors) {
			utils.ErrorResp(w, http.StatusForbidden, "Field validation error")
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	short, err := h.s.RedirectURL(ctx, shortenUrl)
	if err != nil {
		utils.ErrorResp(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	http.Redirect(w, r, short.Long, http.StatusFound)
}
